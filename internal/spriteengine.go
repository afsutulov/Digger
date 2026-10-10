// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

type SpriteEngine struct {
	dig             *Digger
	sprDrawFlag     []bool
	sprRecFlag      []bool
	sprEnabled      []bool
	sprChar         []int
	sprBackground   [][]int
	sprX            []int
	sprY            []int
	sprWidth        []int
	sprHeight       []int
	sprBWidth       []int
	sprBHeight      []int
	sprNewChar      []int
	sprNewWidth     []int
	sprNewHeight    []int
	sprNewBWidth    []int
	sprNewBHeight   []int
	defaultSprOrder []int
	sprOrder        []int
}

func newSpriteEngine(parent *Digger) *SpriteEngine {
	self := &SpriteEngine{}
	self.sprDrawFlag = make([]bool, 17)
	self.sprRecFlag = make([]bool, 17)
	self.sprEnabled = make([]bool, 16)
	self.sprChar = make([]int, 17)
	self.sprBackground = make([][]int, 16)
	self.sprX = make([]int, 17)
	self.sprY = make([]int, 17)
	self.sprWidth = make([]int, 17)
	self.sprHeight = make([]int, 17)
	self.sprBWidth = make([]int, 16)
	self.sprBHeight = make([]int, 16)
	self.sprNewChar = make([]int, 16)
	self.sprNewWidth = make([]int, 16)
	self.sprNewHeight = make([]int, 16)
	self.sprNewBWidth = make([]int, 16)
	self.sprNewBHeight = make([]int, 16)
	self.defaultSprOrder = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	self.sprOrder = self.defaultSprOrder
	self.dig = parent
	return self
}
func (self *SpriteEngine) bcollide(bx int, si int) bool {
	if self.sprX[bx] >= self.sprX[si] {
		if (self.sprX[bx] + self.sprBWidth[bx]) > ((((self.sprWidth[si] * 4) + self.sprX[si]) - self.sprBWidth[si]) - 1) {
			return false
		}
	} else {
		if (self.sprX[si] + self.sprBWidth[si]) > ((((self.sprWidth[bx] * 4) + self.sprX[bx]) - self.sprBWidth[bx]) - 1) {
			return false
		}
	}
	if self.sprY[bx] >= self.sprY[si] {
		if (self.sprY[bx] + self.sprBHeight[bx]) <= (((self.sprHeight[si] + self.sprY[si]) - self.sprBHeight[si]) - 1) {
			return true
		}
		return false
	}
	if (self.sprY[si] + self.sprBHeight[si]) <= (((self.sprHeight[bx] + self.sprY[bx]) - self.sprBHeight[bx]) - 1) {
		return true
	}
	return false
}
func (self *SpriteEngine) bcollides(bx int) int {
	var si int = bx
	_ = si
	var ax int = 0
	_ = ax
	var dx int = 0
	_ = dx
	bx = 0
	for {
		if self.sprEnabled[bx] && (bx != si) {
			if self.bcollide(bx, si) {
				ax |= (1 << dx)
			}
			self.sprX[bx] += 320
			self.sprY[bx] -= 2
			if self.bcollide(bx, si) {
				ax |= (1 << dx)
			}
			self.sprX[bx] -= 640
			self.sprY[bx] += 4
			if self.bcollide(bx, si) {
				ax |= (1 << dx)
			}
			self.sprX[bx] += 320
			self.sprY[bx] -= 2
		}
		bx++
		dx++
		if !(dx != 16) {
			break
		}
	}
	return ax
}
func (self *SpriteEngine) clearDrawFlags() {
	self.clearRecFlags()
	{
		var i int = 0
		_ = i
		for ; i < 17; i++ {
			self.sprDrawFlag[i] = false
		}
	}
}
func (self *SpriteEngine) clearRecFlags() {
	{
		var i int = 0
		_ = i
		for ; i < 17; i++ {
			self.sprRecFlag[i] = false
		}
	}
}
func (self *SpriteEngine) collide(bx int, si int) bool {
	if self.sprX[bx] >= self.sprX[si] {
		if self.sprX[bx] > (((self.sprWidth[si] * 4) + self.sprX[si]) - 1) {
			return false
		}
	} else {
		if self.sprX[si] > (((self.sprWidth[bx] * 4) + self.sprX[bx]) - 1) {
			return false
		}
	}
	if self.sprY[bx] >= self.sprY[si] {
		if self.sprY[bx] <= ((self.sprHeight[si] + self.sprY[si]) - 1) {
			return true
		}
		return false
	}
	if self.sprY[si] <= ((self.sprHeight[bx] + self.sprY[bx]) - 1) {
		return true
	}
	return false
}
func (self *SpriteEngine) createSprite(n int, ch int, mov []int, wid int, hei int, bwid int, bhei int) {
	self.sprChar[(n & 15)] = ch
	self.sprNewChar[(n & 15)] = self.sprChar[(n & 15)]
	self.sprBackground[(n & 15)] = mov
	self.sprWidth[(n & 15)] = wid
	self.sprNewWidth[(n & 15)] = self.sprWidth[(n & 15)]
	self.sprHeight[(n & 15)] = hei
	self.sprNewHeight[(n & 15)] = self.sprHeight[(n & 15)]
	self.sprBWidth[(n & 15)] = bwid
	self.sprNewBWidth[(n & 15)] = self.sprBWidth[(n & 15)]
	self.sprBHeight[(n & 15)] = bhei
	self.sprNewBHeight[(n & 15)] = self.sprBHeight[(n & 15)]
	self.sprEnabled[(n & 15)] = false
}
func (self *SpriteEngine) drawMiscSprite(x int, y int, ch int, wid int, hei int) {
	self.sprX[16] = (x & -(4))
	self.sprY[16] = y
	self.sprChar[16] = ch
	self.sprWidth[16] = wid
	self.sprHeight[16] = hei
	self.dig.display.drawSpriteMasked(self.sprX[16], self.sprY[16], self.sprChar[16], self.sprWidth[16], self.sprHeight[16])
}
func (self *SpriteEngine) redrawSprite(n int, x int, y int) int {
	var bx int
	_ = bx
	var t1 int
	_ = t1
	var t2 int
	_ = t2
	var t3 int
	_ = t3
	var t4 int
	_ = t4
	bx = (n & 15)
	x &= -(4)
	self.clearDrawFlags()
	self.setRedrawFlags(bx)
	t1 = self.sprX[bx]
	t2 = self.sprY[bx]
	t3 = self.sprWidth[bx]
	t4 = self.sprHeight[bx]
	self.sprX[bx] = x
	self.sprY[bx] = y
	self.sprWidth[bx] = self.sprNewWidth[bx]
	self.sprHeight[bx] = self.sprNewHeight[bx]
	self.clearRecFlags()
	self.setRedrawFlags(bx)
	self.sprHeight[bx] = t4
	self.sprWidth[bx] = t3
	self.sprY[bx] = t2
	self.sprX[bx] = t1
	self.sprDrawFlag[bx] = true
	self.restoreBackgrounds()
	self.sprX[bx] = x
	self.sprY[bx] = y
	self.sprChar[bx] = self.sprNewChar[bx]
	self.sprWidth[bx] = self.sprNewWidth[bx]
	self.sprHeight[bx] = self.sprNewHeight[bx]
	self.sprBWidth[bx] = self.sprNewBWidth[bx]
	self.sprBHeight[bx] = self.sprNewBHeight[bx]
	self.dig.display.readSpritePixels(self.sprX[bx], self.sprY[bx], self.sprBackground[bx], self.sprWidth[bx], self.sprHeight[bx])
	self.drawMaskedSprites()
	return self.bcollides(bx)
}
func (self *SpriteEngine) eraseSprite(n int) {
	var bx int = (n & 15)
	_ = bx
	self.dig.display.drawSprite(self.sprX[bx], self.sprY[bx], self.sprBackground[bx], self.sprWidth[bx], self.sprHeight[bx])
	self.sprEnabled[bx] = false
	self.clearDrawFlags()
	self.setRedrawFlags(bx)
	self.drawMaskedSprites()
}
func (self *SpriteEngine) captureAndDrawSprites() {
	{
		var i int = 0
		_ = i
		for ; i < 16; i++ {
			if self.sprDrawFlag[i] {
				self.dig.display.readSpritePixels(self.sprX[i], self.sprY[i], self.sprBackground[i], self.sprWidth[i], self.sprHeight[i])
			}
		}
	}
	self.drawMaskedSprites()
}
func (self *SpriteEngine) initMiscSprite(x int, y int, wid int, hei int) {
	self.sprX[16] = x
	self.sprY[16] = y
	self.sprWidth[16] = wid
	self.sprHeight[16] = hei
	self.clearDrawFlags()
	self.setRedrawFlags(16)
	self.restoreBackgrounds()
}
func (self *SpriteEngine) initSprite(n int, ch int, wid int, hei int, bwid int, bhei int) {
	self.sprNewChar[(n & 15)] = ch
	self.sprNewWidth[(n & 15)] = wid
	self.sprNewHeight[(n & 15)] = hei
	self.sprNewBWidth[(n & 15)] = bwid
	self.sprNewBHeight[(n & 15)] = bhei
}
func (self *SpriteEngine) moveDrawSprite(n int, x int, y int) int {
	var bx int = (n & 15)
	_ = bx
	self.sprX[bx] = (x & -(4))
	self.sprY[bx] = y
	self.sprChar[bx] = self.sprNewChar[bx]
	self.sprWidth[bx] = self.sprNewWidth[bx]
	self.sprHeight[bx] = self.sprNewHeight[bx]
	self.sprBWidth[bx] = self.sprNewBWidth[bx]
	self.sprBHeight[bx] = self.sprNewBHeight[bx]
	self.clearDrawFlags()
	self.setRedrawFlags(bx)
	self.restoreBackgrounds()
	self.dig.display.readSpritePixels(self.sprX[bx], self.sprY[bx], self.sprBackground[bx], self.sprWidth[bx], self.sprHeight[bx])
	self.sprEnabled[bx] = true
	self.sprDrawFlag[bx] = true
	self.drawMaskedSprites()
	return self.bcollides(bx)
}
func (self *SpriteEngine) drawMaskedSprites() {
	{
		var i int = 0
		_ = i
		for ; i < 16; i++ {
			var j int = self.sprOrder[i]
			_ = j
			if self.sprDrawFlag[j] {
				self.dig.display.drawSpriteMasked(self.sprX[j], self.sprY[j], self.sprChar[j], self.sprWidth[j], self.sprHeight[j])
			}
		}
	}
}
func (self *SpriteEngine) restoreBackgrounds() {
	{
		var i int = 0
		_ = i
		for ; i < 16; i++ {
			if self.sprDrawFlag[i] {
				self.dig.display.drawSprite(self.sprX[i], self.sprY[i], self.sprBackground[i], self.sprWidth[i], self.sprHeight[i])
			}
		}
	}
}
func (self *SpriteEngine) setRedrawFlags(n int) {
	if !(self.sprRecFlag[n]) {
		self.sprRecFlag[n] = true
		{
			var i int = 0
			_ = i
			for ; i < 16; i++ {
				if self.sprEnabled[i] && (i != n) {
					if self.collide(i, n) {
						self.sprDrawFlag[i] = true
						self.setRedrawFlags(i)
					}
					self.sprX[i] += 320
					self.sprY[i] -= 2
					if self.collide(i, n) {
						self.sprDrawFlag[i] = true
						self.setRedrawFlags(i)
					}
					self.sprX[i] -= 640
					self.sprY[i] += 4
					if self.collide(i, n) {
						self.sprDrawFlag[i] = true
						self.setRedrawFlags(i)
					}
					self.sprX[i] += 320
					self.sprY[i] -= 2
				}
			}
		}
	}
}
func (self *SpriteEngine) setSpriteOrder(newsprorder []int) {
	if newsprorder == nil {
		self.sprOrder = self.defaultSprOrder
	} else {
		self.sprOrder = newsprorder
	}
}
