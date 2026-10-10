// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

type Main struct {
	dig                 *Digger
	digSprOrder         []int
	gameData            []*GameState
	playerDisplayBuffer string
	currentPlayer       int
	numPlayers          int
	penalty             int
	levelNotDrawn       bool
	flashPlayer         bool
	speedMul            int
	randSeed            int
	levelData           [][]string
}

func newMain(parent *Digger) *Main {
	self := &Main{}
	self.digSprOrder = []int{14, 13, 7, 6, 5, 4, 3, 2, 1, 12, 11, 10, 9, 8, 15, 0}
	self.gameData = []*GameState{newGameState(), newGameState()}
	self.playerDisplayBuffer = ""
	self.currentPlayer = 0
	self.numPlayers = 0
	self.penalty = 0
	self.levelNotDrawn = false
	self.flashPlayer = false
	self.speedMul = 40
	self.levelData = [][]string{[]string{"S   B     HHHHS", "V  CC  C  V B  ", "VB CC  C  V    ", "V  CCB CB V CCC", "V  CC  C  V CCC", "HH CC  C  V CCC", " V    B B V    ", " HHHH     V    ", "C   V     V   C", "CC  HHHHHHH  CC"}, []string{"SHHHHH  B B  HS", " CC  V       V ", " CC  V CCCCC V ", "BCCB V CCCCC V ", "CCCC V       V ", "CCCC V B  HHHH ", " CC  V CC V    ", " BB  VCCCCV CC ", "C    V CC V CC ", "CC   HHHHHH    "}, []string{"SHHHHB B BHHHHS", "CC  V C C V BB ", "C   V C C V CC ", " BB V C C VCCCC", "CCCCV C C VCCCC", "CCCCHHHHHHH CC ", " CC  C V C  CC ", " CC  C V C     ", "C    C V C    C", "CC   C H C   CC"}, []string{"SHBCCCCBCCCCBHS", "CV  CCCCCCC  VC", "CHHH CCCCC HHHC", "C  V  CCC  V  C", "   HHH C HHH   ", "  B  V B V  B  ", "  C  VCCCV  C  ", " CCC HHHHH CCC ", "CCCCC CVC CCCCC", "CCCCC CHC CCCCC"}, []string{"SHHHHHHHHHHHHHS", "VBCCCCBVCCCCCCV", "VCCCCCCV CCBC V", "V CCCC VCCBCCCV", "VCCCCCCV CCCC V", "V CCCC VBCCCCCV", "VCCBCCCV CCCC V", "V CCBC VCCCCCCV", "VCCCCCCVCCCCCCV", "HHHHHHHHHHHHHHH"}, []string{"SHHHHHHHHHHHHHS", "VCBCCV V VCCBCV", "VCCC VBVBV CCCV", "VCCCHH V HHCCCV", "VCC V CVC V CCV", "VCCHH CVC HHCCV", "VC V CCVCC V CV", "VCHHBCCVCCBHHCV", "VCVCCCCVCCCCVCV", "HHHHHHHHHHHHHHH"}, []string{"SHCCCCCVCCCCCHS", " VCBCBCVCBCBCV ", "BVCCCCCVCCCCCVB", "CHHCCCCVCCCCHHC", "CCV CCCVCCC VCC", "CCHHHCCVCCHHHCC", "CCCCV CVC VCCCC", "CCCCHH V HHCCCC", "CCCCCV V VCCCCC", "CCCCCHHHHHCCCCC"}, []string{"HHHHHHHHHHHHHHS", "V CCBCCCCCBCC V", "HHHCCCCBCCCCHHH", "VBV CCCCCCC VBV", "VCHHHCCCCCHHHCV", "VCCBV CCC VBCCV", "VCCCHHHCHHHCCCV", "VCCCC V V CCCCV", "VCCCCCV VCCCCCV", "HHHHHHHHHHHHHHH"}}
	self.dig = parent
	return self
}
func (self *Main) addLife(pl int) {
	postInc(&self.gameData[(pl - 1)].lives)
	self.dig.sound.sound1Up()
}
func (self *Main) calibrate() {
	self.dig.sound.volume = 1
}
func (self *Main) checkLevelDone() {
	if ((self.dig.countEmeralds() == 0) || (self.dig.monster.getMonstersLeft() == 0)) && self.dig.digOnScreen {
		self.gameData[self.currentPlayer].levelDone = true
	} else {
		self.gameData[self.currentPlayer].levelDone = false
	}
}
func (self *Main) clearTopLine() {
	self.dig.drawing.drawText4("                          ", 0, 0, 3)
	self.dig.drawing.drawText4(" ", 308, 0, 3)
}
func (self *Main) drawScreen() {
	self.dig.drawing.createAllSprites()
	self.dig.drawing.drawFieldAndBackground()
	self.dig.bags.drawBags()
	self.dig.drawEmeralds()
	self.dig.initDigger()
	self.dig.monster.initMonsters()
}
func (self *Main) getCurrentPlayer() int {
	return self.currentPlayer
}
func (self *Main) getLevelChar(x int, y int, l int) int {
	if l == 0 {
		l++
	}
	return int(self.levelData[(l - 1)][y][x])
}
func (self *Main) getLives(pl int) int {
	return self.gameData[(pl - 1)].lives
}
func (self *Main) incrementPenalty() {
	self.penalty++
}
func (self *Main) initChars() {
	self.dig.drawing.initAllSprites()
	self.dig.initDigger()
	self.dig.monster.initMonsters()
}
func (self *Main) initLevel() {
	self.gameData[self.currentPlayer].levelDone = false
	self.dig.drawing.buildField()
	self.dig.makeEmeraldField()
	self.dig.bags.initBags()
	self.levelNotDrawn = true
}
func (self *Main) getLevelNumber() int {
	return self.gameData[self.currentPlayer].level
}
func (self *Main) getLevelNumberClampedToTen() int {
	if self.gameData[self.currentPlayer].level > 10 {
		return 10
	}
	return self.gameData[self.currentPlayer].level
}
func (self *Main) main() {
	var frame int
	_ = frame
	var t int
	_ = t
	var x int = 0
	_ = x
	var start bool
	_ = start
	self.dig.time = self.dig.display.getCurrentTimeMillis()
	self.calibrate()
	self.dig.ftime = (self.speedMul * 2000)
	self.dig.display.init()
	self.dig.display.setPalette(0)
	self.dig.input.initKeyboard()
	self.dig.input.detectJoystick()
	self.dig.scores.loadScores()
	self.dig.sound.initSound()
	self.numPlayers = self.dig.host.cfg.Players
	for {
		self.dig.sound.soundStop()
		self.dig.sprite.setSpriteOrder(self.digSprOrder)
		self.dig.drawing.createAllSprites()
		self.dig.input.detectJoystick()
		self.dig.display.clearScreen()
		self.dig.display.drawTitleScreen()
		self.dig.drawing.drawText4("D I G G E R", 100, 0, 3)
		self.showPlayers()
		self.dig.scores.showTable()
		start = false
		frame = 0
		self.dig.time = self.dig.display.getCurrentTimeMillis()
		for !(start) {
			start = self.dig.input.testStart()
			if self.dig.input.lastKeyCode != 0 {
				if self.dig.input.lastKeyCode == 27 {
					self.switchPlayerCount()
					self.showPlayers()
				}
				self.dig.input.lastAsciiKeyCode = 0
				self.dig.input.lastKeyCode = 0
			}
			if frame == 0 {
				{
					t = 54
					for ; t < 174; t += 12 {
						self.dig.drawing.drawText4("            ", 164, t, 0)
					}
				}
			}
			if frame == 50 {
				self.dig.sprite.moveDrawSprite(8, 292, 63)
				x = 292
			}
			if (frame > 50) && (frame <= 77) {
				x -= 4
				self.dig.drawing.drawMonster(0, true, 4, x, 63)
			}
			if frame > 77 {
				self.dig.drawing.drawMonster(0, true, 0, 184, 63)
			}
			if frame == 83 {
				self.dig.drawing.drawText4("NOBBIN", 216, 64, 2)
			}
			if frame == 90 {
				self.dig.sprite.moveDrawSprite(9, 292, 82)
				self.dig.drawing.drawMonster(1, false, 4, 292, 82)
				x = 292
			}
			if (frame > 90) && (frame <= 117) {
				x -= 4
				self.dig.drawing.drawMonster(1, false, 4, x, 82)
			}
			if frame > 117 {
				self.dig.drawing.drawMonster(1, false, 0, 184, 82)
			}
			if frame == 123 {
				self.dig.drawing.drawText4("HOBBIN", 216, 83, 2)
			}
			if frame == 130 {
				self.dig.sprite.moveDrawSprite(0, 292, 101)
				self.dig.drawing.drawDigger(4, 292, 101, true)
				x = 292
			}
			if (frame > 130) && (frame <= 157) {
				x -= 4
				self.dig.drawing.drawDigger(4, x, 101, true)
			}
			if frame > 157 {
				self.dig.drawing.drawDigger(0, 184, 101, true)
			}
			if frame == 163 {
				self.dig.drawing.drawText4("DIGGER", 216, 102, 2)
			}
			if frame == 178 {
				self.dig.sprite.moveDrawSprite(1, 184, 120)
				self.dig.drawing.drawGold(1, 0, 184, 120)
			}
			if frame == 183 {
				self.dig.drawing.drawText4("GOLD", 216, 121, 2)
			}
			if frame == 198 {
				self.dig.drawing.drawEmerald(184, 141)
			}
			if frame == 203 {
				self.dig.drawing.drawText4("EMERALD", 216, 140, 2)
			}
			if frame == 218 {
				self.dig.drawing.drawBonus(184, 158)
			}
			if frame == 223 {
				self.dig.drawing.drawText4("BONUS", 216, 159, 2)
			}
			self.dig.newFrame()
			frame++
			if frame > 250 {
				frame = 0
			}
			if !(self.dig.running) && (frame > 1) {
				return
			}
		}
		self.dig.input.lastAsciiKeyCode = 0
		self.dig.input.lastKeyCode = 0
		self.gameData[0].level = self.dig.host.cfg.StartLevel
		self.gameData[0].lives = 3
		if self.numPlayers == 2 {
			self.gameData[1].level = self.dig.host.cfg.StartLevel
			self.gameData[1].lives = 3
		} else {
			self.gameData[1].lives = 0
		}
		self.dig.display.clearScreen()
		self.currentPlayer = 0
		self.initLevel()
		self.currentPlayer = 1
		self.initLevel()
		self.dig.scores.zeroScores()
		self.dig.bonusVisible = true
		if self.numPlayers == 2 {
			self.flashPlayer = true
		}
		self.currentPlayer = 0
		for ((self.gameData[0].lives != 0) || (self.gameData[1].lives != 0)) && !(self.dig.input.escape) {
			self.gameData[self.currentPlayer].dead = false
			for (!(self.gameData[self.currentPlayer].dead) && (self.gameData[self.currentPlayer].lives != 0)) && !(self.dig.input.escape) {
				self.dig.drawing.initAllSprites()
				self.play()
			}
			if self.gameData[(1-self.currentPlayer)].lives != 0 {
				self.currentPlayer = (1 - self.currentPlayer)
				self.levelNotDrawn = true
				self.flashPlayer = self.levelNotDrawn
			}
		}
		self.dig.input.escape = false
		if !(true) {
			break
		}
	}
}
func (self *Main) play() {
	var t int
	_ = t
	var c int
	_ = c
	if self.levelNotDrawn {
		self.levelNotDrawn = false
		self.drawScreen()
		self.dig.time = self.dig.display.getCurrentTimeMillis()
		if self.flashPlayer {
			self.flashPlayer = false
			self.playerDisplayBuffer = "PLAYER "
			self.playerDisplayBuffer += choose((self.currentPlayer == 0), "1", "2")
			self.clearTopLine()
			{
				t = 0
				for ; t < 15; t++ {
					{
						c = 1
						for ; c <= 3; c++ {
							self.dig.drawing.drawText4(self.playerDisplayBuffer, 108, 0, c)
							self.dig.scores.writeCurrentScore(c)
							self.dig.newFrame()
							if self.dig.input.escape {
								return
							}
						}
					}
				}
			}
			self.dig.scores.drawScores()
			self.dig.scores.addScore(0)
		}
	} else {
		self.initChars()
	}
	self.dig.input.lastKeyCode = 0
	self.dig.drawing.drawText4("        ", 108, 0, 3)
	self.dig.scores.initScores()
	self.dig.drawing.drawLives()
	self.dig.sound.music(1)
	self.dig.input.readDirection()
	self.dig.time = self.dig.display.getCurrentTimeMillis()
	for (!(self.gameData[self.currentPlayer].dead) && !(self.gameData[self.currentPlayer].levelDone)) && !(self.dig.input.escape) {
		self.penalty = 0
		self.dig.doDigger()
		self.dig.monster.doMonsters()
		self.dig.bags.doBags()
		if self.penalty > 8 {
			self.dig.monster.increaseMonsterDelay((self.penalty - 8))
		}
		self.testPause()
		self.checkLevelDone()
	}
	self.dig.eraseDigger()
	self.dig.sound.musicOff()
	t = 20
	for ((self.dig.bags.getMovingBagsCount() != 0) || (t != 0)) && !(self.dig.input.escape) {
		if t != 0 {
			t--
		}
		self.penalty = 0
		self.dig.bags.doBags()
		self.dig.doDigger()
		self.dig.monster.doMonsters()
		if self.penalty < 8 {
			t = 0
		}
	}
	self.dig.sound.soundStop()
	self.dig.killFire()
	self.dig.eraseBonus()
	self.dig.bags.cleanupBags()
	self.dig.drawing.saveFieldSnapshot()
	self.dig.monster.eraseMonsters()
	self.dig.newFrame()
	if self.gameData[self.currentPlayer].levelDone {
		self.dig.sound.soundLevDone()
	}
	if self.dig.countEmeralds() == 0 {
		self.gameData[self.currentPlayer].level++
		if self.gameData[self.currentPlayer].level > 1000 {
			self.gameData[self.currentPlayer].level = 1000
		}
		self.initLevel()
	}
	if self.gameData[self.currentPlayer].dead {
		self.gameData[self.currentPlayer].lives--
		self.dig.drawing.drawLives()
		if (self.gameData[self.currentPlayer].lives == 0) && !(self.dig.input.escape) {
			self.dig.scores.endOfGame()
		}
	}
	if self.gameData[self.currentPlayer].levelDone {
		self.gameData[self.currentPlayer].level++
		if self.gameData[self.currentPlayer].level > 1000 {
			self.gameData[self.currentPlayer].level = 1000
		}
		self.initLevel()
	}
}
func (self *Main) setDead(bp6 bool) {
	self.gameData[self.currentPlayer].dead = bp6
}
func (self *Main) showPlayers() {
	if self.numPlayers == 1 {
		self.dig.drawing.drawText4("ONE", 220, 25, 3)
		self.dig.drawing.drawText4(" PLAYER ", 192, 39, 3)
	} else {
		self.dig.drawing.drawText4("TWO", 220, 25, 3)
		self.dig.drawing.drawText4(" PLAYERS", 184, 39, 3)
	}
}
func (self *Main) switchPlayerCount() {
	self.numPlayers = (3 - self.numPlayers)
}
