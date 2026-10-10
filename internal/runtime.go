// SPDX-License-Identifier: GPL-3.0-only
package game

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const Width, Height, SampleRate = 320, 200, 44100

// Control contains key transitions in chronological order, followed by text input.
// DOS scan codes are used so the engine is independent of a windowing toolkit.
type KeyEvent struct {
	Code int
	Down bool
}
type Control struct {
	Keys []KeyEvent
	Text string
}
type Config struct {
	FrameMS    int
	Seed       uint32
	ScoresPath string
	Players    int
	StartLevel int
	Silent     bool
}
type Frame struct {
	Pixels                                         [Width * Height]byte
	Intensity                                      int
	DurationMS                                     int
	PCM                                            []byte // signed 16-bit stereo, little-endian, 44100 Hz
	Playing, Paused                                bool
	Player, Players, Level, Lives, Score, Emeralds int
	Error                                          error
}

// Engine owns all mutable game state in one goroutine. Step is synchronous;
// neither the renderer nor the audio device reads the game's live state.
type Engine struct {
	d                       *Digger
	cfg                     Config
	input                   chan Control
	frames                  chan Frame
	stop                    chan struct{}
	once                    sync.Once
	phase0, phase2          float64
	soundClock, sampleClock float64
	paused                  bool
}
type stopped struct{}

func New(cfg Config) (*Engine, Frame, error) {
	if cfg.FrameMS == 0 {
		cfg.FrameMS = 80
	}
	cfg.FrameMS = max(40, min(200, cfg.FrameMS))
	if cfg.Players != 2 {
		cfg.Players = 1
	}
	cfg.StartLevel = max(1, min(1000, cfg.StartLevel))
	e := &Engine{cfg: cfg, input: make(chan Control), frames: make(chan Frame), stop: make(chan struct{})}
	d := newDiggerFields()
	e.d = d
	d.host = e
	d.running = true
	d.frametime = cfg.FrameMS
	d.bags = newBags(d)
	d.main = newMain(d)
	d.sound = newSound(d)
	d.monster = newMonster(d)
	d.scores = newScores(d)
	d.sprite = newSpriteEngine(d)
	d.drawing = newDrawing(d)
	d.input = newInput(d)
	d.display = newCgaDisplay(d)
	d.display.pixels = make([]int, 65536)
	d.main.randSeed = int(cfg.Seed)
	d.sound.soundFlag = !cfg.Silent
	go func() {
		defer close(e.frames)
		defer func() {
			if p := recover(); p != nil {
				if _, ok := p.(stopped); !ok {
					select {
					case e.frames <- Frame{Error: fmt.Errorf("game engine: %v", p)}:
					case <-e.stop:
					}
				}
			}
		}()
		d.main.main()
	}()
	f, ok := <-e.frames
	if !ok {
		return nil, Frame{}, errors.New("game stopped during initialization")
	}
	if f.Error != nil {
		e.Close()
		return nil, f, f.Error
	}
	return e, f, nil
}
func (e *Engine) Step(c Control) (Frame, error) {
	select {
	case e.input <- c:
	case <-e.stop:
		return Frame{}, errors.New("game closed")
	}
	f, ok := <-e.frames
	if !ok {
		return Frame{}, errors.New("game stopped")
	}
	return f, f.Error
}
func (e *Engine) Close()    { e.once.Do(func() { close(e.stop) }) }
func (d *Digger) newFrame() { d.waitFrame(d.frametime) }
func (d *Digger) waitFrame(ms int) {
	e := d.host
	d.time += ms
	f := Frame{Intensity: d.display.intensity, DurationMS: ms, Paused: e.paused, Playing: d.diggerVisible,
		Player: d.main.currentPlayer + 1, Players: d.main.numPlayers, Level: d.main.getLevelNumber(), Lives: d.main.getLives(d.main.currentPlayer + 1), Emeralds: d.countEmeralds()}
	if d.main.currentPlayer == 0 {
		f.Score = d.scores.player1Score
	} else {
		f.Score = d.scores.player2Score
	}
	for i := range f.Pixels {
		f.Pixels[i] = byte(d.display.pixels[i] & 3)
	}
	f.PCM = e.generateAudio(ms)
	select {
	case e.frames <- f:
	case <-e.stop:
		panic(stopped{})
	}
	var c Control
	select {
	case c = <-e.input:
	case <-e.stop:
		panic(stopped{})
	}
	e.applyControl(c)
	d.input.checkKeyboard()
}
func (e *Engine) applyControl(c Control) {
	i, s := e.d.input, e.d.sound
	for _, k := range c.Keys {
		switch k.Code {
		case 72:
			if k.Down {
				i.keyUpPressed()
			} else {
				i.keyUpReleased()
			}
		case 80:
			if k.Down {
				i.keyDownPressed()
			} else {
				i.keyDownReleased()
			}
		case 75:
			if k.Down {
				i.keyLeftPressed()
			} else {
				i.keyLeftReleased()
			}
		case 77:
			if k.Down {
				i.keyRightPressed()
			} else {
				i.keyRightReleased()
			}
		case 59:
			if k.Down {
				i.keyF1Pressed()
			} else {
				i.keyF1Released()
			}
		case 65:
			if k.Down {
				s.musicFlag = !s.musicFlag
				if s.musicFlag && !s.musicPlaying && s.soundFlag {
					s.music(s.tuneno)
				}
			}
		case 67:
			if k.Down {
				s.soundFlag = !s.soundFlag
			}
		case 68:
			if k.Down {
				i.escape = true
			}
		case 43:
			if k.Down {
				i.plusPressed = true
			}
		case 45:
			if k.Down {
				i.minusPressed = true
			}
		default:
			if k.Down {
				i.processKey(k.Code)
			}
		}
	}
	for _, r := range c.Text {
		if r >= 'a' && r <= 'z' {
			r -= 32
		}
		if r >= 32 && r < 127 {
			i.processKey(int(r))
		}
	}
}

func (m *Main) randomNumber(n int) int {
	m.randSeed = int(int32(uint32(m.randSeed)*0x15a4e35 + 1))
	return (m.randSeed & 0x7fffffff) % n
}
func (m *Main) getLevelPlan() int {
	l := m.getLevelNumber()
	if l > 8 {
		l = (l & 3) + 5
	}
	return l
}
func (m *Main) testPause() {
	d := m.dig
	if d.input.lastAsciiKeyCode == 32 {
		d.input.lastAsciiKeyCode = 0
		d.sound.soundPause()
		m.clearTopLine()
		d.drawing.drawText4("PAUSED", 124, 0, 1)
		d.host.paused = true
		d.input.lastKeyCode = 0
		for d.input.lastKeyCode == 0 && !d.input.escape {
			d.newFrame()
		}
		d.host.paused = false
		m.clearTopLine()
		d.scores.drawScores()
		d.scores.addScore(0)
		d.drawing.drawLives()
		d.input.lastKeyCode = 0
		d.input.lastAsciiKeyCode = 0
		d.sound.soundPauseOff()
	} else {
		d.sound.soundPauseOff()
	}
}
func (d *CgaDisplay) clearScreen()                  { clear(d.pixels) }
func (d *CgaDisplay) getCurrentTimeMillis() int     { return d.dig.time }
func (d *CgaDisplay) setIntensity(i int)            { d.intensity = i & 1 }
func (d *CgaDisplay) setPalette(p int)              {} // original Digger uses CGA palette zero
func (d *CgaDisplay) drawChar4(x, y, ch, color int) { d.drawChar5(x, y, ch, color, false) }
func (d *CgaDisplay) drawChar5(x, y, ch, color int, refresh bool) {
	ch -= 32
	if ch < 0 || ch >= len(Alphabet_CGA_ASCII_TABLE) {
		return
	}
	p := Alphabet_CGA_ASCII_TABLE[ch]
	if p == nil {
		return
	}
	for row := 0; row < 12; row++ {
		for col := 0; col < 12; col++ {
			xx, yy := x+col, y+row
			if xx < 0 || xx >= Width || yy < 0 || yy >= Height {
				continue
			}
			d.pixels[yy*Width+xx] = (p[row*3+col/4] >> uint(6-2*(col%4))) & (color & 3)
		}
	}
}

func (s *Sound) initSound() {
	s.wavetype = 2
	s.t0val = 12000
	s.musvol = 8
	s.t2val = 40
	s.soundT0Flag = true
	s.sndFlag = s.soundFlag
	s.spkrmode = 0
	s.int8Flag = true
	s.musicPlaying = false
	s.setSoundT2()
	s.soundStop()
}
func (s *Sound) killSound()    { s.stopInt8() }
func (s *Sound) setupSound()   { s.sndFlag = s.soundFlag; s.startInt8() }
func (s *Sound) pushToEngine() {}
func (s *Sound) soundPause() {
	s.soundPausedFlag = true
	s.t0val = Sound_SILENCE_T0VAL
	s.t2val = Sound_DEFAULT_T2VAL
	s.spkrmode = 0
}
func (s *Sound) soundLevDone() {
	s.soundStop()
	s.nljpointer = 0
	s.nljnoteduration = 20
	s.soundLevDoneFlag = true
	for s.soundLevDoneFlag && !s.dig.input.escape {
		s.dig.waitFrame(14)
	}
	s.soundLevDoneFlag = false
}

// Drive the original timer interrupt at 1193180/16384 Hz, independent of render FPS.
// Two timer phases preserve the PC speaker's fast alternation of music and effects.
func (e *Engine) generateAudio(ms int) []byte {
	e.sampleClock += float64(ms) * SampleRate / 1000
	n := int(e.sampleClock)
	e.sampleClock -= float64(n)
	b := make([]byte, n*4)
	s := e.d.sound
	for j := 0; j < n; j++ {
		e.soundClock += 1193180.0 / 16384.0 / SampleRate
		if e.soundClock >= 1 {
			e.soundClock -= 1
			if s.int8Flag {
				if s.soundLevDoneFlag {
					s.timerClock++
					s.soundLevDoneUpdate()
				} else {
					s.soundInt()
				}
			}
		}
		sample := 0.0
		if s.soundFlag && s.int8Flag && !s.soundPausedFlag {
			div := s.t2val
			amp := 0.22
			if s.spkrmode != 0 {
				div = s.t0val
				amp *= min(1, float64(max(1, s.pulsewidth))/50)
			}
			if div > 0 {
				hz := 1193180.0 / float64(div)
				if hz >= 20 && hz < 20000 && div != 32000 {
					phase := &e.phase2
					if s.spkrmode != 0 {
						phase = &e.phase0
					}
					if *phase < 0.5 {
						sample = amp
					} else {
						sample = -amp
					}
					*phase += hz / SampleRate
					*phase -= float64(int(*phase))
				}
			}
		}
		v := uint16(int16(sample * 32767))
		binary.LittleEndian.PutUint16(b[j*4:], v)
		binary.LittleEndian.PutUint16(b[j*4+2:], v)
	}
	return b
}

// The score table retains the original one-based ranking and player-two life quirk.
type scoreRecord struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s *Scores) readScores() {
	for i := range s.scoreinit {
		s.scoreinit[i] = "---"
	}
	clear(s.scoreHigh)
	records, _ := decodeHSC(initialHSC)
	if s.dig.host.cfg.ScoresPath != "" {
		if b, err := os.ReadFile(s.dig.host.cfg.ScoresPath); err == nil {
			var loaded []scoreRecord
			if json.Unmarshal(b, &loaded) == nil {
				records = loaded
			}
		}
	}

	for i, r := range records {
		if i >= 10 {
			break
		}
		s.scoreinit[i+1] = cleanInitials(r.Name)
		s.scoreHigh[i+2] = max(0, min(999999, r.Score))
	}
}
func (s *Scores) saveScores() {
	p := s.dig.host.cfg.ScoresPath
	if p == "" {
		return
	}
	a := make([]scoreRecord, 10)
	for i := range a {
		a[i] = scoreRecord{s.scoreinit[i+1], s.scoreHigh[i+2]}
	}
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return
	}
	if err = os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return
	}
	f, err := os.CreateTemp(filepath.Dir(p), "scores-*.json")
	if err != nil {
		return
	}
	tmp := f.Name()
	defer os.Remove(tmp)
	if _, err = f.Write(b); err != nil {
		f.Close()
		return
	}
	if err = f.Close(); err != nil {
		return
	}
	// Replace works on POSIX; on Windows remove the previous table before renaming.
	if err = os.Rename(tmp, p); err != nil {
		if os.Remove(p) == nil {
			_ = os.Rename(tmp, p)
		}
	}
}
func cleanInitials(name string) string {
	p := []byte("---")
	for i, c := range []byte(name) {
		if i == 3 {
			break
		}
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		if c >= 32 && c <= 90 {
			p[i] = c
		}
	}
	return string(p)
}
func (s *Scores) numberToString(n int) string { return fmt.Sprintf("%6d", n%1000000) }
func (s *Scores) flashyWait(n int)            { s.dig.waitFrame(max(1, n*2)) }
func (s *Scores) getInitials() {
	d := s.dig
	d.drawing.drawText4("ENTER YOUR", 100, 70, 3)
	d.drawing.drawText4(" INITIALS", 100, 90, 3)
	d.drawing.drawText4("_ _ _", 128, 130, 3)
	s.scoreinit[0] = "..."
	d.sound.killSound()
	chars := []byte("...")
	pos := 0
	elapsed := 0
	d.input.lastKeyCode = 0
	for pos < 3 && elapsed < 30000 && !d.input.escape {
		d.waitFrame(50)
		elapsed += 50
		k := d.input.lastKeyCode
		d.input.lastKeyCode = 0
		if k == 8 && pos > 0 {
			pos--
			chars[pos] = '.'
			d.display.drawChar4(128+pos*24, 130, '_', 3)
			continue
		}
		if s.isRegularKey(k) && d.input.getAsciiKey(k) != 0 {
			chars[pos] = byte(k)
			d.display.drawChar4(128+pos*24, 130, k, 3)
			pos++
			elapsed = 0
		}
	}
	s.scoreinit[0] = string(chars)
	d.input.lastAsciiKeyCode = 0
	d.input.lastKeyCode = 0
	d.sound.setupSound()
	d.display.setIntensity(0)
}

func choose[T any](c bool, a, b T) T {
	if c {
		return a
	}
	return b
}
func numString(n int) string  { return strconv.Itoa(n) }
func charString(n int) string { return string(rune(n)) }
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func postInc(n *int) int { o := *n; *n++; return o }
func postDec(n *int) int { o := *n; *n--; return o }
func preInc(n *int) int  { *n++; return *n }
func preDec(n *int) int  { *n--; return *n }
func matrix(n, m int) [][]int {
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}
