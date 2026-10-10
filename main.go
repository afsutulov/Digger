// SPDX-License-Identifier: GPL-3.0-only
package main

import (
	"Digger/internal"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type settings struct {
	FrameMS      int  `json:"frame_ms"`
	SquarePixels bool `json:"square_pixels"`
	Fullscreen   bool `json:"fullscreen"`
}
type pcmStream struct {
	mu   sync.Mutex
	data []byte
}

func (p *pcmStream) Read(b []byte) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	n := copy(b, p.data)
	p.data = p.data[n:]
	clear(b[n:])
	return len(b), nil
}
func (p *pcmStream) push(b []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data = append(p.data, b...)
	if len(p.data) > game.SampleRate*4/2 {
		p.data = p.data[len(p.data)-game.SampleRate*4/2:]
	}
}

var _ io.Reader = (*pcmStream)(nil)

type app struct {
	engine        *game.Engine
	frame         game.Frame
	image         *ebiten.Image
	rgba          []byte
	pending       game.Control
	held          map[int]bool
	elapsed       float64
	square        bool
	player        *audio.Player
	stream        *pcmStream
	capture       string
	frames, smoke int
	settingsPath  string
}

var keyMap = []struct {
	code int
	keys []ebiten.Key
}{
	{72, []ebiten.Key{ebiten.KeyArrowUp, ebiten.KeyNumpad8}},
	{80, []ebiten.Key{ebiten.KeyArrowDown, ebiten.KeyNumpad2}},
	{75, []ebiten.Key{ebiten.KeyArrowLeft, ebiten.KeyNumpad4}},
	{77, []ebiten.Key{ebiten.KeyArrowRight, ebiten.KeyNumpad6}},
	{59, []ebiten.Key{ebiten.KeyF1}}, {65, []ebiten.Key{ebiten.KeyF7}},
	{67, []ebiten.Key{ebiten.KeyF9}}, {68, []ebiten.Key{ebiten.KeyF10}},
	{27, []ebiten.Key{ebiten.KeyEscape}}, {32, []ebiten.Key{ebiten.KeySpace}},
	{13, []ebiten.Key{ebiten.KeyEnter, ebiten.KeyNumpadEnter}}, {8, []ebiten.Key{ebiten.KeyBackspace}},
	{43, []ebiten.Key{ebiten.KeyEqual, ebiten.KeyNumpadAdd}},
	{45, []ebiten.Key{ebiten.KeyMinus, ebiten.KeyNumpadSubtract}},
}

func (a *app) Update() error {
	if (ebiten.IsKeyPressed(ebiten.KeyAlt) || ebiten.IsKeyPressed(ebiten.KeyAltLeft) || ebiten.IsKeyPressed(ebiten.KeyAltRight)) && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	pads := ebiten.AppendGamepadIDs(nil)
	for _, bind := range keyMap {
		down := false
		if ebiten.IsFocused() {
			for _, k := range bind.keys {
				down = down || ebiten.IsKeyPressed(k)
			}
			for _, id := range pads {
				if !ebiten.IsStandardGamepadLayoutAvailable(id) {
					continue
				}
				switch bind.code {
				case 72:
					down = down || ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical) < -.35 || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftTop)
				case 80:
					down = down || ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical) > .35 || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftBottom)
				case 75:
					down = down || ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal) < -.35 || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftLeft)
				case 77:
					down = down || ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal) > .35 || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftRight)
				case 59:
					down = down || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonRightBottom)
				case 13:
					down = down || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonCenterRight)
				}
			}
		}
		if bind.code == 13 && ebiten.IsKeyPressed(ebiten.KeyAlt) {
			down = false
		}
		if down != a.held[bind.code] {
			a.pending.Keys = append(a.pending.Keys, game.KeyEvent{Code: bind.code, Down: down})
			a.held[bind.code] = down
		}
	}
	for _, r := range ebiten.AppendInputChars(nil) {
		if r > 32 && r < 127 && r != '+' && r != '-' && r != '=' {
			a.pending.Text += string(r)
		}
	}
	a.elapsed += 1000.0 / 60
	if a.elapsed >= float64(a.frame.DurationMS) {
		a.elapsed -= float64(a.frame.DurationMS)
		f, err := a.engine.Step(a.pending)
		a.pending = game.Control{}
		if err != nil {
			return err
		}
		a.frame = f
		a.frames++
		a.updatePixels()
		if a.stream != nil {
			a.stream.push(f.PCM)
		}
		if a.smoke > 0 && a.frames >= a.smoke {
			return ebiten.Termination
		}
	}
	return nil
}

var palettes = [2][4][3]byte{{{0, 0, 0}, {0, 170, 0}, {170, 0, 0}, {170, 85, 0}}, {{0, 0, 0}, {85, 255, 85}, {255, 85, 85}, {255, 255, 85}}}

func (a *app) updatePixels() {
	p := palettes[a.frame.Intensity&1]
	for i, v := range a.frame.Pixels {
		c := p[v&3]
		a.rgba[i*4] = c[0]
		a.rgba[i*4+1] = c[1]
		a.rgba[i*4+2] = c[2]
		a.rgba[i*4+3] = 255
	}
	a.image.WritePixels(a.rgba)
}
func (a *app) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if !a.square {
		op.GeoM.Scale(1, 1.2)
	}
	op.Filter = ebiten.FilterNearest
	screen.DrawImage(a.image, op)
}
func (a *app) Layout(w, h int) (int, int) {
	if a.square {
		return game.Width, game.Height
	}
	return game.Width, 240
}
func defaultDataDir() string {
	p, err := os.UserConfigDir()
	if err != nil {
		return "digger-data"
	}
	return filepath.Join(p, "Digger-Go")
}
func main() {
	dataDir := flag.String("data-dir", defaultDataDir(), "directory for scores and settings")
	speed := flag.Int("speed", 0, "milliseconds per game frame (40..200, default 80)")
	fullscreen := flag.Bool("fullscreen", false, "start in fullscreen")
	square := flag.Bool("square-pixels", false, "320:200 display instead of original CRT 4:3")
	mute := flag.Bool("mute", false, "disable sound")
	players := flag.Int("players", 1, "one player or two alternating players")
	level := flag.Int("level", 1, "starting level (1..1000)")
	seed := flag.Uint("seed", uint(time.Now().UnixNano()), "random seed")
	smoke := flag.Int("smoke-frames", 0, "exit after this many engine frames (testing)")
	importHSC := flag.String("import-hsc", "", "import an original 512-byte DIGGER.HSC table")
	flag.Parse()
	settingsPath := filepath.Join(*dataDir, "settings.json")
	set := settings{FrameMS: 80}
	if b, err := os.ReadFile(settingsPath); err == nil {
		_ = json.Unmarshal(b, &set)
	}
	if *speed != 0 {
		set.FrameMS = *speed
	}
	if *fullscreen {
		set.Fullscreen = true
	}
	if *square {
		set.SquarePixels = true
	}
	if *importHSC != "" {
		if err := game.ImportHSC(*importHSC, filepath.Join(*dataDir, "scores.json")); err != nil {
			log.Fatal(err)
		}
	}
	e, f, err := game.New(game.Config{FrameMS: set.FrameMS, Seed: uint32(*seed), ScoresPath: filepath.Join(*dataDir, "scores.json"), Players: *players, StartLevel: *level, Silent: *mute})
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	a := &app{engine: e, frame: f, image: ebiten.NewImage(game.Width, game.Height), rgba: make([]byte, game.Width*game.Height*4), held: make(map[int]bool), square: set.SquarePixels, smoke: *smoke, settingsPath: settingsPath}
	a.updatePixels()
	if !*mute {
		ctx := audio.NewContext(game.SampleRate)
		a.stream = &pcmStream{}
		a.player, err = ctx.NewPlayer(a.stream)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Audio unavailable:", err)
			a.stream = nil
		} else {
			a.player.SetBufferSize(60 * time.Millisecond)
			a.stream.push(f.PCM)
			a.player.Play()
			defer a.player.Close()
		}
	}
	w, h := 960, 720
	if a.square {
		h = 600
	}
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Digger — 1983 / Go")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetFullscreen(set.Fullscreen)
	ebiten.SetTPS(60)
	ebiten.SetScreenFilterEnabled(false)
	if err = ebiten.RunGame(a); err != nil && !errors.Is(err, ebiten.Termination) {
		log.Fatal(err)
	}
	set.FrameMS = a.frame.DurationMS
	set.Fullscreen = ebiten.IsFullscreen()
	set.SquarePixels = a.square
	if os.MkdirAll(*dataDir, 0755) == nil {
		b, _ := json.MarshalIndent(set, "", "  ")
		_ = os.WriteFile(settingsPath, b, 0644)
	}
}
