// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

const Drawing_FIELD_WIDTH int = 15
const Drawing_FIELD_HEIGHT int = 10
const Drawing_FIELD_SIZE int = (Drawing_FIELD_WIDTH * Drawing_FIELD_HEIGHT)
const Drawing_BIT_EMERALD int = 0x2000
const Drawing_BIT_TUNNEL_DRAWN int = 0x1000
const Drawing_CLEAR_HORIZONTAL int = 0xd03f
const Drawing_CLEAR_VERTICAL int = 0xdfe0
const Drawing_ALL_HORIZONTAL int = 0x1f
const Drawing_ALL_VERTICAL int = 0xfc0

type Drawing struct {
	bitMasks        []int
	dig             *Digger
	field1          []int
	field2          []int
	field           []int
	diggerBuf       []int
	bagBuf1         []int
	bagBuf2         []int
	bagBuf3         []int
	bagBuf4         []int
	bagBuf5         []int
	bagBuf6         []int
	bagBuf7         []int
	monBuf1         []int
	monBuf2         []int
	monBuf3         []int
	monBuf4         []int
	monBuf5         []int
	monBuf6         []int
	bonusBuf        []int
	fireBuf         []int
	monSpriteFrame  []int
	monSpriteDir    []int
	digSpriteFrame  int
	digSpriteDir    int
	fireSpriteFrame int
	fireHeight      int
}

func newDrawing(parent *Digger) *Drawing {
	self := &Drawing{}
	self.bitMasks = []int{0xfffe, 0xfffd, 0xfffb, 0xfff7, 0xffef, 0xffdf, 0xffbf, 0xff7f, 0xfeff, 0xfdff, 0xfbff, 0xf7ff}
	self.field1 = make([]int, Drawing_FIELD_SIZE)
	self.field2 = make([]int, Drawing_FIELD_SIZE)
	self.field = make([]int, Drawing_FIELD_SIZE)
	self.diggerBuf = make([]int, 480)
	self.bagBuf1 = make([]int, 480)
	self.bagBuf2 = make([]int, 480)
	self.bagBuf3 = make([]int, 480)
	self.bagBuf4 = make([]int, 480)
	self.bagBuf5 = make([]int, 480)
	self.bagBuf6 = make([]int, 480)
	self.bagBuf7 = make([]int, 480)
	self.monBuf1 = make([]int, 480)
	self.monBuf2 = make([]int, 480)
	self.monBuf3 = make([]int, 480)
	self.monBuf4 = make([]int, 480)
	self.monBuf5 = make([]int, 480)
	self.monBuf6 = make([]int, 480)
	self.bonusBuf = make([]int, 480)
	self.fireBuf = make([]int, 128)
	self.monSpriteFrame = []int{0, 0, 0, 0, 0, 0}
	self.monSpriteDir = []int{0, 0, 0, 0, 0, 0}
	self.digSpriteFrame = 0
	self.digSpriteDir = 0
	self.fireSpriteFrame = 0
	self.fireHeight = 8
	self.dig = parent
	return self
}
func (self *Drawing) createDiggerBonusFireSprites() {
	self.digSpriteDir = 1
	self.digSpriteFrame = 0
	self.fireSpriteFrame = 0
	self.dig.sprite.createSprite(0, 0, self.diggerBuf, 4, 15, 0, 0)
	self.dig.sprite.createSprite(14, 81, self.bonusBuf, 4, 15, 0, 0)
	self.dig.sprite.createSprite(15, 82, self.fireBuf, 2, self.fireHeight, 0, 0)
}
func (self *Drawing) createAllSprites() {
	self.dig.sprite.createSprite(1, 62, self.bagBuf1, 4, 15, 0, 0)
	self.dig.sprite.createSprite(2, 62, self.bagBuf2, 4, 15, 0, 0)
	self.dig.sprite.createSprite(3, 62, self.bagBuf3, 4, 15, 0, 0)
	self.dig.sprite.createSprite(4, 62, self.bagBuf4, 4, 15, 0, 0)
	self.dig.sprite.createSprite(5, 62, self.bagBuf5, 4, 15, 0, 0)
	self.dig.sprite.createSprite(6, 62, self.bagBuf6, 4, 15, 0, 0)
	self.dig.sprite.createSprite(7, 62, self.bagBuf7, 4, 15, 0, 0)
	self.dig.sprite.createSprite(8, 71, self.monBuf1, 4, 15, 0, 0)
	self.dig.sprite.createSprite(9, 71, self.monBuf2, 4, 15, 0, 0)
	self.dig.sprite.createSprite(10, 71, self.monBuf3, 4, 15, 0, 0)
	self.dig.sprite.createSprite(11, 71, self.monBuf4, 4, 15, 0, 0)
	self.dig.sprite.createSprite(12, 71, self.monBuf5, 4, 15, 0, 0)
	self.dig.sprite.createSprite(13, 71, self.monBuf6, 4, 15, 0, 0)
	self.createDiggerBonusFireSprites()
	{
		var i int = 0
		_ = i
		for ; i < 6; i++ {
			self.monSpriteFrame[i] = 0
			self.monSpriteDir[i] = 1
		}
	}
}
func (self *Drawing) drawBackground(levelPlan int) {
	{
		var y int = 14
		_ = y
		for ; y < 200; y += 4 {
			{
				var x int = 0
				_ = x
				for ; x < 320; x += 20 {
					self.dig.sprite.drawMiscSprite(x, y, (93 + levelPlan), 5, 4)
				}
			}
		}
	}
}
func (self *Drawing) drawBonus(x int, y int) {
	self.dig.sprite.initSprite(14, 81, 4, 15, 0, 0)
	self.dig.sprite.moveDrawSprite(14, x, y)
}
func (self *Drawing) drawTunnelEdgeBottom(x int, y int) {
	self.dig.sprite.initMiscSprite((x - 4), (y + 15), 6, 6)
	self.dig.sprite.drawMiscSprite((x - 4), (y + 15), 105, 6, 6)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) drawDigger(frame int, x int, y int, right bool) int {
	self.digSpriteFrame += self.digSpriteDir
	if (self.digSpriteFrame == 2) || (self.digSpriteFrame == 0) {
		self.digSpriteDir = -(self.digSpriteDir)
	}
	self.digSpriteFrame = max(0, min(2, self.digSpriteFrame))
	if ((frame >= 0) && (frame <= 6)) && !((frame & 1) != 0) {
		self.dig.sprite.initSprite(0, ((((frame + choose(right, 0, 1)) * 3) + self.digSpriteFrame) + 1), 4, 15, 0, 0)
		return self.dig.sprite.redrawSprite(0, x, y)
	}
	if (frame >= 10) && (frame <= 15) {
		self.dig.sprite.initSprite(0, (40 - frame), 4, 15, 0, 0)
		return self.dig.sprite.redrawSprite(0, x, y)
	}
	return 0
}
func (self *Drawing) drawEmerald(x int, y int) {
	self.dig.sprite.initMiscSprite(x, y, 4, 10)
	self.dig.sprite.drawMiscSprite(x, y, 108, 4, 10)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) drawField() {
	var x int
	_ = x
	var y int
	_ = y
	var xp int
	_ = xp
	var yp int
	_ = yp
	{
		x = 0
		for ; x < Drawing_FIELD_WIDTH; x++ {
			{
				y = 0
				for ; y < Drawing_FIELD_HEIGHT; y++ {
					if (self.field[((y*Drawing_FIELD_WIDTH)+x)] & Drawing_BIT_EMERALD) == 0 {
						xp = ((x * 20) + 12)
						yp = ((y * 18) + 18)
						if (self.field[((y*Drawing_FIELD_WIDTH)+x)] & Drawing_ALL_VERTICAL) != Drawing_ALL_VERTICAL {
							self.field[((y * Drawing_FIELD_WIDTH) + x)] &= Drawing_CLEAR_HORIZONTAL
							self.drawTunnelEdgeBottom(xp, (yp - 15))
							self.drawTunnelEdgeBottom(xp, (yp - 12))
							self.drawTunnelEdgeBottom(xp, (yp - 9))
							self.drawTunnelEdgeBottom(xp, (yp - 6))
							self.drawTunnelEdgeBottom(xp, (yp - 3))
							self.drawTunnelEdgeTop(xp, (yp + 3))
						}
						if (self.field[((y*Drawing_FIELD_WIDTH)+x)] & Drawing_ALL_HORIZONTAL) != Drawing_ALL_HORIZONTAL {
							self.field[((y * Drawing_FIELD_WIDTH) + x)] &= Drawing_CLEAR_VERTICAL
							self.drawTunnelEdgeRight((xp - 16), yp)
							self.drawTunnelEdgeRight((xp - 12), yp)
							self.drawTunnelEdgeRight((xp - 8), yp)
							self.drawTunnelEdgeRight((xp - 4), yp)
							self.drawTunnelEdgeLeft((xp + 4), yp)
						}
						if x < 14 {
							if (self.field[(((y*Drawing_FIELD_WIDTH)+x)+1)] & 0xfdf) != 0xfdf {
								self.drawTunnelEdgeRight(xp, yp)
							}
						}
						if y < 9 {
							if (self.field[(((y+1)*Drawing_FIELD_WIDTH)+x)] & 0xfdf) != 0xfdf {
								self.drawTunnelEdgeBottom(xp, yp)
							}
						}
					}
				}
			}
		}
	}
}
func (self *Drawing) drawFire(x int, y int, frame int) int {
	if frame == 0 {
		self.fireSpriteFrame++
		if self.fireSpriteFrame > 2 {
			self.fireSpriteFrame = 0
		}
		self.dig.sprite.initSprite(15, (82 + self.fireSpriteFrame), 2, self.fireHeight, 0, 0)
	} else {
		self.dig.sprite.initSprite(15, (84 + frame), 2, self.fireHeight, 0, 0)
	}
	return self.dig.sprite.redrawSprite(15, x, y)
}
func (self *Drawing) drawBagFallDebris(x int, y int) {
	self.dig.sprite.initMiscSprite((x - 4), (y + 15), 6, 8)
	self.dig.sprite.drawMiscSprite((x - 4), (y + 15), 107, 6, 8)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) drawGold(spriteIndex int, frame int, x int, y int) int {
	self.dig.sprite.initSprite(spriteIndex, (frame + 62), 4, 15, 0, 0)
	return self.dig.sprite.redrawSprite(spriteIndex, x, y)
}
func (self *Drawing) drawTunnelEdgeLeft(x int, y int) {
	self.dig.sprite.initMiscSprite((x - 8), (y - 1), 2, 18)
	self.dig.sprite.drawMiscSprite((x - 8), (y - 1), 104, 2, 18)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) drawLife(kind int, x int, y int) {
	self.dig.sprite.drawMiscSprite(x, y, (kind + 110), 4, 12)
}
func (self *Drawing) drawLives() {
	var l int
	_ = l
	var n int
	_ = n
	n = (self.dig.main.getLives(1) - 1)
	{
		l = 1
		for ; l < 5; l++ {
			self.drawLife(choose((n > 0), 0, 2), ((l * 20) + 60), 0)
			n--
		}
	}
	if self.dig.main.numPlayers == 2 {
		n = (self.dig.main.getLives(2) - 1)
		{
			l = 1
			for ; l < 5; l++ {
				self.drawLife(choose((n > 0), 1, 2), (244 - (l * 20)), 0)
				n--
			}
		}
	}
}
func (self *Drawing) drawMonster(index int, isNobbin bool, direction int, x int, y int) int {
	self.monSpriteFrame[index] += self.monSpriteDir[index]
	if (self.monSpriteFrame[index] == 2) || (self.monSpriteFrame[index] == 0) {
		self.monSpriteDir[index] = -(self.monSpriteDir[index])
	}
	self.monSpriteFrame[index] = max(0, min(2, self.monSpriteFrame[index]))
	if isNobbin {
		self.dig.sprite.initSprite((index + 8), (self.monSpriteFrame[index] + 69), 4, 15, 0, 0)
	} else {
		switch direction {
		case 0:
			self.dig.sprite.initSprite((index + 8), (self.monSpriteFrame[index] + 73), 4, 15, 0, 0)
		case 4:
			self.dig.sprite.initSprite((index + 8), (self.monSpriteFrame[index] + 77), 4, 15, 0, 0)
		}
	}
	return self.dig.sprite.redrawSprite((index + 8), x, y)
}
func (self *Drawing) drawMonsterDeath(index int, isNobbin bool, direction int, x int, y int) int {
	if isNobbin {
		self.dig.sprite.initSprite((index + 8), 72, 4, 15, 0, 0)
	} else {
		switch direction {
		case 0:
			self.dig.sprite.initSprite((index + 8), 76, 4, 15, 0, 0)
		case 4:
			self.dig.sprite.initSprite((index + 8), 80, 4, 14, 0, 0)
		}
	}
	return self.dig.sprite.redrawSprite((index + 8), x, y)
}
func (self *Drawing) drawTunnelEdgeRight(x int, y int) {
	self.dig.sprite.initMiscSprite((x + 16), (y - 1), 2, 18)
	self.dig.sprite.drawMiscSprite((x + 16), (y - 1), 102, 2, 18)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) drawBagGroundCrack(x int, y int) {
	self.dig.sprite.initMiscSprite((x - 4), (y + 17), 6, 6)
	self.dig.sprite.drawMiscSprite((x - 4), (y + 17), 106, 6, 6)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) drawFieldAndBackground() {
	var x int
	_ = x
	var y int
	_ = y
	{
		x = 0
		for ; x < Drawing_FIELD_WIDTH; x++ {
			{
				y = 0
				for ; y < Drawing_FIELD_HEIGHT; y++ {
					if self.dig.main.getCurrentPlayer() == 0 {
						self.field[((y * Drawing_FIELD_WIDTH) + x)] = self.field1[((y * Drawing_FIELD_WIDTH) + x)]
					} else {
						self.field[((y * Drawing_FIELD_WIDTH) + x)] = self.field2[((y * Drawing_FIELD_WIDTH) + x)]
					}
				}
			}
		}
	}
	self.dig.display.setPalette(0)
	self.dig.display.setIntensity(0)
	self.drawBackground(self.dig.main.getLevelPlan())
	self.drawField()

}
func (self *Drawing) drawTunnelEdgeTop(x int, y int) {
	self.dig.sprite.initMiscSprite((x - 4), (y - 6), 6, 6)
	self.dig.sprite.drawMiscSprite((x - 4), (y - 6), 103, 6, 6)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) digTunnel(x int, y int, dir int) {
	var h int = ((x - 12) / 20)
	_ = h
	var xr int = (((x - 12) % 20) / 4)
	_ = xr
	var v int = ((y - 18) / 18)
	_ = v
	var yr int = (((y - 18) % 18) / 3)
	_ = yr
	self.dig.main.incrementPenalty()
	switch dir {
	case 0:
		h++
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= self.bitMasks[xr]
		if (self.field[((v*Drawing_FIELD_WIDTH)+h)] & Drawing_ALL_HORIZONTAL) != 0 {
			break
		}
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= ^(Drawing_BIT_TUNNEL_DRAWN)
	case 4:
		xr--
		if xr < 0 {
			xr += 5
			h--
		}
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= self.bitMasks[xr]
		if (self.field[((v*Drawing_FIELD_WIDTH)+h)] & Drawing_ALL_HORIZONTAL) != 0 {
			break
		}
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= ^(Drawing_BIT_TUNNEL_DRAWN)
	case 2:
		yr--
		if yr < 0 {
			yr += 6
			v--
		}
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= self.bitMasks[(6 + yr)]
		if (self.field[((v*Drawing_FIELD_WIDTH)+h)] & Drawing_ALL_VERTICAL) != 0 {
			break
		}
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= ^(Drawing_BIT_TUNNEL_DRAWN)
	case 6:
		v++
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= self.bitMasks[(6 + yr)]
		if (self.field[((v*Drawing_FIELD_WIDTH)+h)] & Drawing_ALL_VERTICAL) != 0 {
			break
		}
		self.field[((v * Drawing_FIELD_WIDTH) + h)] &= ^(Drawing_BIT_TUNNEL_DRAWN)
	}
}
func (self *Drawing) eraseEmerald(x int, y int) {
	self.dig.sprite.initMiscSprite(x, y, 4, 10)
	self.dig.sprite.drawMiscSprite(x, y, 109, 4, 10)
	self.dig.sprite.captureAndDrawSprites()
}
func (self *Drawing) initDiggerBonusFireSprites() {
	self.digSpriteDir = 1
	self.digSpriteFrame = 0
	self.fireSpriteFrame = 0
	self.dig.sprite.initSprite(0, 0, 4, 15, 0, 0)
	self.dig.sprite.initSprite(14, 81, 4, 15, 0, 0)
	self.dig.sprite.initSprite(15, 82, 2, self.fireHeight, 0, 0)
}
func (self *Drawing) initAllSprites() {
	self.dig.sprite.initSprite(1, 62, 4, 15, 0, 0)
	self.dig.sprite.initSprite(2, 62, 4, 15, 0, 0)
	self.dig.sprite.initSprite(3, 62, 4, 15, 0, 0)
	self.dig.sprite.initSprite(4, 62, 4, 15, 0, 0)
	self.dig.sprite.initSprite(5, 62, 4, 15, 0, 0)
	self.dig.sprite.initSprite(6, 62, 4, 15, 0, 0)
	self.dig.sprite.initSprite(7, 62, 4, 15, 0, 0)
	self.dig.sprite.initSprite(8, 71, 4, 15, 0, 0)
	self.dig.sprite.initSprite(9, 71, 4, 15, 0, 0)
	self.dig.sprite.initSprite(10, 71, 4, 15, 0, 0)
	self.dig.sprite.initSprite(11, 71, 4, 15, 0, 0)
	self.dig.sprite.initSprite(12, 71, 4, 15, 0, 0)
	self.dig.sprite.initSprite(13, 71, 4, 15, 0, 0)
	self.initDiggerBonusFireSprites()
}
func (self *Drawing) buildField() {
	var c int
	_ = c
	var x int
	_ = x
	var y int
	_ = y
	{
		x = 0
		for ; x < Drawing_FIELD_WIDTH; x++ {
			{
				y = 0
				for ; y < Drawing_FIELD_HEIGHT; y++ {
					self.field[((y * Drawing_FIELD_WIDTH) + x)] = -(1)
					c = self.dig.main.getLevelChar(x, y, self.dig.main.getLevelPlan())
					if (c == 83) || (c == 86) {
						self.field[((y * Drawing_FIELD_WIDTH) + x)] &= Drawing_CLEAR_HORIZONTAL
					}
					if (c == 83) || (c == 72) {
						self.field[((y * Drawing_FIELD_WIDTH) + x)] &= Drawing_CLEAR_VERTICAL
					}
					if self.dig.main.getCurrentPlayer() == 0 {
						self.field1[((y * Drawing_FIELD_WIDTH) + x)] = self.field[((y * Drawing_FIELD_WIDTH) + x)]
					} else {
						self.field2[((y * Drawing_FIELD_WIDTH) + x)] = self.field[((y * Drawing_FIELD_WIDTH) + x)]
					}
				}
			}
		}
	}
}
func (self *Drawing) drawText4(text string, x int, y int, color int) {
	self.drawText5(text, x, y, color, false)
}
func (self *Drawing) drawText5(text string, x int, y int, color int, refresh bool) {
	var rx int = x
	_ = rx
	{
		var i int = 0
		_ = i
		for ; i < len(text); i++ {
			self.dig.display.drawChar4(x, y, int(text[i]), color)
			x += 12
		}
	}
	if refresh {

	}
}
func (self *Drawing) saveFieldSnapshot() {
	var x int
	_ = x
	var y int
	_ = y
	{
		x = 0
		for ; x < Drawing_FIELD_WIDTH; x++ {
			{
				y = 0
				for ; y < Drawing_FIELD_HEIGHT; y++ {
					if self.dig.main.getCurrentPlayer() == 0 {
						self.field1[((y * Drawing_FIELD_WIDTH) + x)] = self.field[((y * Drawing_FIELD_WIDTH) + x)]
					} else {
						self.field2[((y * Drawing_FIELD_WIDTH) + x)] = self.field[((y * Drawing_FIELD_WIDTH) + x)]
					}
				}
			}
		}
	}
}
