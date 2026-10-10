// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

const Scores_MAX_SCORES int = 11
const Scores_PLAYER_ONE int = 0
const Scores_PLAYER_TWO int = 1

type Scores struct {
	dig                  *Digger
	substr               string
	highbuf              []int
	scoreHigh            []int
	scoreinit            []string
	finalScore           int
	player1Score         int
	player2Score         int
	player1NextLifeScore int
	player2NextLifeScore int
	hsbuf                string
	bonusScore           int
	gotInitFlag          bool
}

func newScores(parent *Digger) *Scores {
	self := &Scores{}
	self.highbuf = make([]int, 10)
	self.scoreHigh = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	self.scoreinit = make([]string, Scores_MAX_SCORES)
	self.finalScore = 0
	self.player1Score = 0
	self.player2Score = 0
	self.player1NextLifeScore = 0
	self.player2NextLifeScore = 0
	self.bonusScore = 20000
	self.gotInitFlag = false
	self.dig = parent
	return self
}
func (self *Scores) addScore(score int) {
	if self.dig.main.getCurrentPlayer() == Scores_PLAYER_ONE {
		self.player1Score += score
		if self.player1Score > 999999 {
			self.player1Score = 0
		}
		self.writenum(self.player1Score, 0, 0, 6, 1)
		if self.player1Score >= self.player1NextLifeScore {
			if self.dig.main.getLives(1) < 5 {
				self.dig.main.addLife(1)
				self.dig.drawing.drawLives()
			}
			self.player1NextLifeScore += self.bonusScore
		}
	} else {
		self.player2Score += score
		if self.player2Score > 999999 {
			self.player2Score = 0
		}
		if self.player2Score < 100000 {
			self.writenum(self.player2Score, 236, 0, 6, 1)
		} else {
			self.writenum(self.player2Score, 248, 0, 6, 1)
		}
		if self.player2Score > self.player2NextLifeScore {
			if self.dig.main.getLives(2) < 5 {
				self.dig.main.addLife(2)
				self.dig.drawing.drawLives()
			}
			self.player2NextLifeScore += self.bonusScore
		}
	}
	self.dig.main.incrementPenalty()
	self.dig.main.incrementPenalty()
	self.dig.main.incrementPenalty()
}
func (self *Scores) drawScores() {
	self.writenum(self.player1Score, 0, 0, 6, 3)
	if self.dig.main.numPlayers == 2 {
		if self.player2Score < 100000 {
			self.writenum(self.player2Score, 236, 0, 6, 3)
		} else {
			self.writenum(self.player2Score, 248, 0, 6, 3)
		}
	}
}
func (self *Scores) endOfGame() {
	var i int
	_ = i
	var j int
	_ = j
	self.addScore(0)
	if self.dig.main.getCurrentPlayer() == Scores_PLAYER_ONE {
		self.finalScore = self.player1Score
	} else {
		self.finalScore = self.player2Score
	}
	if self.finalScore > self.scoreHigh[11] {
		self.dig.display.clearScreen()
		self.drawScores()
		self.dig.main.playerDisplayBuffer = "PLAYER "
		if self.dig.main.getCurrentPlayer() == Scores_PLAYER_ONE {
			self.dig.main.playerDisplayBuffer += "1"
		} else {
			self.dig.main.playerDisplayBuffer += "2"
		}
		self.dig.drawing.drawText5(self.dig.main.playerDisplayBuffer, 108, 0, 2, true)
		self.getInitials()
		self.shuffleHigh()
		self.saveScores()
	} else {
		self.dig.main.clearTopLine()
		self.dig.drawing.drawText5("GAME OVER", 104, 0, 3, true)
		self.dig.sound.killSound()
		{
			j = 0
			for ; j < 20; j++ {
				{
					i = 0
					for ; i < 2; i++ {
						self.dig.display.setPalette((1 - (j & 1)))
						self.flashyWait(1)
						self.dig.display.setPalette(0)
						self.dig.display.setIntensity(((1 - i) & 1))
						self.dig.newFrame()
					}
				}
			}
		}
		self.dig.sound.setupSound()
		self.dig.drawing.drawText5("         ", 104, 0, 3, true)
	}
}
func (self *Scores) isRegularKey(kp int) bool {
	return ((kp != 0) && ((kp & 0x80) == 0))
}
func (self *Scores) initScores() {
	self.addScore(0)
}
func (self *Scores) loadScores() {
	self.readScores()
}
func (self *Scores) scoreBonus() {
	self.addScore(1000)
}
func (self *Scores) scoreEatMonster() {
	self.addScore((self.dig.monsterEatMultiplier * 200))
	self.dig.monsterEatMultiplier <<= 1
}
func (self *Scores) scoreEmerald() {
	self.addScore(25)
}
func (self *Scores) scoreGold() {
	self.addScore(500)
}
func (self *Scores) scoreKillMonster() {
	self.addScore(250)
}
func (self *Scores) scoreOctave() {
	self.addScore(250)
}
func (self *Scores) showTable() {
	var i int
	_ = i
	var col int
	_ = col
	self.dig.drawing.drawText4("HIGH SCORES", 16, 25, 3)
	col = 2
	{
		i = 1
		for ; i < 11; i++ {
			self.hsbuf = ((self.scoreinit[i] + "  ") + self.numberToString(self.scoreHigh[(i+1)]))
			self.dig.drawing.drawText4(self.hsbuf, 16, (31 + (13 * i)), col)
			col = 1
		}
	}
}
func (self *Scores) shuffleHigh() {
	var i int
	_ = i
	var j int
	_ = j
	{
		j = 10
		for ; j > 1; j-- {
			if self.finalScore < self.scoreHigh[j] {
				break
			}
		}
	}
	{
		i = 10
		for ; i > j; i-- {
			self.scoreHigh[(i + 1)] = self.scoreHigh[i]
			self.scoreinit[i] = self.scoreinit[(i - 1)]
		}
	}
	self.scoreHigh[(j + 1)] = self.finalScore
	self.scoreinit[j] = self.scoreinit[0]
}
func (self *Scores) writeCurrentScore(bp6 int) {
	if self.dig.main.getCurrentPlayer() == Scores_PLAYER_ONE {
		self.writenum(self.player1Score, 0, 0, 6, bp6)
	} else {
		if self.player2Score < 100000 {
			self.writenum(self.player2Score, 236, 0, 6, bp6)
		} else {
			self.writenum(self.player2Score, 248, 0, 6, bp6)
		}
	}
}
func (self *Scores) writenum(n int, x int, y int, w int, c int) {
	var d int
	_ = d
	var xp int = (((w - 1) * 12) + x)
	_ = xp
	for w > 0 {
		d = (n % 10)
		if (w > 1) || (d > 0) {
			self.dig.display.drawChar5(xp, y, (d + 48), c, false)
		}
		n /= 10
		w--
		xp -= 12
	}
}
func (self *Scores) zeroScores() {
	self.player2Score = 0
	self.player1Score = 0
	self.finalScore = 0
	self.player1NextLifeScore = self.bonusScore
	self.player2NextLifeScore = self.bonusScore
}
