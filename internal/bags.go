// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

const Bags_MAX_BAGS int = 8
const Bags_GOLD_BASE_TIME int = 150

type Bags struct {
	dig        *Digger
	bagdat1    []*BagState
	bagdat2    []*BagState
	bagdat     []*BagState
	pushCount  int
	goldTime   int
	wobbleAnim []int
}

func newBags(parent *Digger) *Bags {
	self := &Bags{}
	self.bagdat1 = make([]*BagState, Bags_MAX_BAGS)
	self.bagdat2 = make([]*BagState, Bags_MAX_BAGS)
	self.bagdat = make([]*BagState, Bags_MAX_BAGS)
	self.pushCount = 0
	self.goldTime = 0
	self.wobbleAnim = []int{2, 0, 1, 0}
	self.dig = parent
	{
		var i int = 0
		_ = i
		for ; i < Bags_MAX_BAGS; i++ {
			self.bagdat[i] = newBagState()
			self.bagdat1[i] = newBagState()
			self.bagdat2[i] = newBagState()
		}
	}
	return self
}
func (self *Bags) bagbits() int {
	var bags int = 0
	_ = bags
	{
		var bag int = 1
		_ = bag
		var bit int = 2
		_ = bit
		for ; bag < Bags_MAX_BAGS; func() { postInc(&bag); bit <<= 1 }() {
			if self.bagdat[bag].exist {
				bags |= bit
			}
		}
	}
	return bags
}
func (self *Bags) baghitground(bag int) {
	var clbits int
	_ = clbits
	if (self.bagdat[bag].direction == 6) && (self.bagdat[bag].fallHeight > 1) {
		self.bagdat[bag].goldTime = 1
	} else {
		self.bagdat[bag].fallHeight = 0
	}
	self.bagdat[bag].direction = -(1)
	self.bagdat[bag].wobbleTime = 15
	self.bagdat[bag].wobbling = false
	clbits = self.dig.drawing.drawGold(bag, 0, self.bagdat[bag].x, self.bagdat[bag].y)
	self.dig.main.incrementPenalty()
	{
		var bn int = 1
		_ = bn
		var b int = 2
		_ = b
		for ; bn < Bags_MAX_BAGS; func() { postInc(&bn); b <<= 1 }() {
			if (b & clbits) != 0 {
				self.removebag(bn)
			}
		}
	}
}
func (self *Bags) bagy(bag int) int {
	return self.bagdat[bag].y
}
func (self *Bags) cleanupBags() {
	self.dig.sound.soundFallOff()
	{
		var bpa int = 1
		_ = bpa
		for ; bpa < Bags_MAX_BAGS; bpa++ {
			if self.bagdat[bpa].exist && (((((((self.bagdat[bpa].h == 7) && (self.bagdat[bpa].v == 9)) || (self.bagdat[bpa].xr != 0)) || (self.bagdat[bpa].yr != 0)) || (self.bagdat[bpa].goldTime != 0)) || (self.bagdat[bpa].fallHeight != 0)) || self.bagdat[bpa].wobbling) {
				self.bagdat[bpa].exist = false
				self.dig.sprite.eraseSprite(bpa)
			}
			if self.dig.main.getCurrentPlayer() == 0 {
				self.bagdat1[bpa].copyFrom(self.bagdat[bpa])
			} else {
				self.bagdat2[bpa].copyFrom(self.bagdat[bpa])
			}
		}
	}
}
func (self *Bags) doBags() {
	var bag int
	_ = bag
	var soundFallOff bool = true
	_ = soundFallOff
	var soundWobbleOff bool = true
	_ = soundWobbleOff
	{
		bag = 1
		for ; bag < Bags_MAX_BAGS; bag++ {
			if self.bagdat[bag].exist {
				if self.bagdat[bag].goldTime != 0 {
					if self.bagdat[bag].goldTime == 1 {
						self.dig.sound.soundBreak()
						self.dig.drawing.drawGold(bag, 4, self.bagdat[bag].x, self.bagdat[bag].y)
						self.dig.main.incrementPenalty()
					}
					if self.bagdat[bag].goldTime == 3 {
						self.dig.drawing.drawGold(bag, 5, self.bagdat[bag].x, self.bagdat[bag].y)
						self.dig.main.incrementPenalty()
					}
					if self.bagdat[bag].goldTime == 5 {
						self.dig.drawing.drawGold(bag, 6, self.bagdat[bag].x, self.bagdat[bag].y)
						self.dig.main.incrementPenalty()
					}
					self.bagdat[bag].goldTime++
					if self.bagdat[bag].goldTime == self.goldTime {
						self.removebag(bag)
					} else {
						if (self.bagdat[bag].v < 9) && (self.bagdat[bag].goldTime < (self.goldTime - 10)) {
							if (self.dig.monster.getfield(self.bagdat[bag].h, (self.bagdat[bag].v+1)) & 0x2000) == 0 {
								self.bagdat[bag].goldTime = (self.goldTime - 10)
							}
						}
					}
				} else {
					self.updatebag(bag)
				}
			}
		}
	}
	{
		bag = 1
		for ; bag < Bags_MAX_BAGS; bag++ {
			if (self.bagdat[bag].direction == 6) && self.bagdat[bag].exist {
				soundFallOff = false
			}
			if ((self.bagdat[bag].direction != 6) && self.bagdat[bag].wobbling) && self.bagdat[bag].exist {
				soundWobbleOff = false
			}
		}
	}
	if soundFallOff {
		self.dig.sound.soundFallOff()
	}
	if soundWobbleOff {
		self.dig.sound.soundWobbleOff()
	}
}
func (self *Bags) drawBags() {
	{
		var bag int = 1
		_ = bag
		for ; bag < Bags_MAX_BAGS; bag++ {
			if self.dig.main.getCurrentPlayer() == 0 {
				self.bagdat[bag].copyFrom(self.bagdat1[bag])
			} else {
				self.bagdat[bag].copyFrom(self.bagdat2[bag])
			}
			if self.bagdat[bag].exist {
				self.dig.sprite.moveDrawSprite(bag, self.bagdat[bag].x, self.bagdat[bag].y)
			}
		}
	}
}
func (self *Bags) getbagdir(bag int) int {
	if self.bagdat[bag].exist {
		return self.bagdat[bag].direction
	}
	return -(1)
}
func (self *Bags) getgold(bag int) {
	var clbits int = self.dig.drawing.drawGold(bag, 6, self.bagdat[bag].x, self.bagdat[bag].y)
	_ = clbits
	self.dig.main.incrementPenalty()
	if (clbits & 1) != 0 {
		self.dig.scores.scoreGold()
		self.dig.sound.soundGold()
		self.dig.digtime = 0
	} else {
		self.dig.monster.monsterGotGold()
	}
	self.removebag(bag)
}
func (self *Bags) getMovingBagsCount() int {
	var n int = 0
	_ = n
	{
		var bag int = 1
		_ = bag
		for ; bag < Bags_MAX_BAGS; bag++ {
			if (self.bagdat[bag].exist && (self.bagdat[bag].goldTime < 10)) && ((self.bagdat[bag].goldTime != 0) || self.bagdat[bag].wobbling) {
				n++
			}
		}
	}
	return n
}
func (self *Bags) initBags() {
	self.pushCount = 0
	self.goldTime = (Bags_GOLD_BASE_TIME - (self.dig.main.getLevelNumberClampedToTen() * 10))
	{
		var bag int = 1
		_ = bag
		for ; bag < Bags_MAX_BAGS; bag++ {
			self.bagdat[bag].exist = false
		}
	}
	var bag int = 1
	_ = bag
	{
		var x int = 0
		_ = x
		for ; x < 15; x++ {
			{
				var y int = 0
				_ = y
				for ; y < 10; y++ {
					if self.dig.main.getLevelChar(x, y, self.dig.main.getLevelPlan()) == 66 {
						if bag < Bags_MAX_BAGS {
							self.bagdat[bag].exist = true
							self.bagdat[bag].goldTime = 0
							self.bagdat[bag].fallHeight = 0
							self.bagdat[bag].direction = -(1)
							self.bagdat[bag].wobbling = false
							self.bagdat[bag].wobbleTime = 15
							self.bagdat[bag].unfallen = true
							self.bagdat[bag].x = ((x * 20) + 12)
							self.bagdat[bag].y = ((y * 18) + 18)
							self.bagdat[bag].h = x
							self.bagdat[bag].v = y
							self.bagdat[bag].xr = 0
							self.bagdat[bag].yr = 0
							bag++
						}
					}
				}
			}
		}
	}
	if self.dig.main.getCurrentPlayer() == 0 {
		{
			var i int = 1
			_ = i
			for ; i < Bags_MAX_BAGS; i++ {
				self.bagdat1[i].copyFrom(self.bagdat[i])
			}
		}
	} else {
		{
			var i int = 1
			_ = i
			for ; i < Bags_MAX_BAGS; i++ {
				self.bagdat2[i].copyFrom(self.bagdat[i])
			}
		}
	}
}
func (self *Bags) pushbag(bag int, dir int) bool {
	var x int
	_ = x
	var y int
	_ = y
	var h int
	_ = h
	var v int
	_ = v
	var ox int
	_ = ox
	var oy int
	_ = oy
	var clbits int
	_ = clbits
	var push bool = true
	_ = push
	x = self.bagdat[bag].x
	ox = x
	y = self.bagdat[bag].y
	oy = y
	h = self.bagdat[bag].h
	v = self.bagdat[bag].v
	if self.bagdat[bag].goldTime != 0 {
		self.getgold(bag)
		return true
	}
	if (self.bagdat[bag].direction == 6) && ((dir == 4) || (dir == 0)) {
		clbits = self.dig.drawing.drawGold(bag, 3, x, y)
		self.dig.main.incrementPenalty()
		if ((clbits & 1) != 0) && (self.dig.diggery >= y) {
			self.dig.killdigger(1, bag)
		}
		if (clbits & 0x3f00) != 0 {
			self.dig.monster.squashmonsters(bag, clbits)
		}
		return true
	}
	if ((((x == 292) && (dir == 0)) || ((x == 12) && (dir == 4))) || ((y == 180) && (dir == 6))) || ((y == 18) && (dir == 2)) {
		push = false
	}
	if push {
		switch dir {
		case 0:
			x += 4
		case 4:
			x -= 4
		case 6:
			if self.bagdat[bag].unfallen {
				self.bagdat[bag].unfallen = false
				self.dig.drawing.drawBagGroundCrack(x, y)
				self.dig.drawing.drawTunnelEdgeTop(x, (y + 21))
			} else {
				self.dig.drawing.drawBagFallDebris(x, y)
			}
			self.dig.drawing.digTunnel(x, y, dir)
			self.dig.killemerald(h, v)
			y += 6
		}
		switch dir {
		case 6:
			clbits = self.dig.drawing.drawGold(bag, 3, x, y)
			self.dig.main.incrementPenalty()
			if ((clbits & 1) != 0) && (self.dig.diggery >= y) {
				self.dig.killdigger(1, bag)
			}
			if (clbits & 0x3f00) != 0 {
				self.dig.monster.squashmonsters(bag, clbits)
			}
		case 0, 4:
			self.bagdat[bag].wobbleTime = 15
			self.bagdat[bag].wobbling = false
			clbits = self.dig.drawing.drawGold(bag, 0, x, y)
			self.dig.main.incrementPenalty()
			self.pushCount = 1
			if (clbits & 0xfe) != 0 {
				if !(self.pushbags(dir, clbits)) {
					x = ox
					y = oy
					self.dig.drawing.drawGold(bag, 0, ox, oy)
					self.dig.main.incrementPenalty()
					push = false
				}
			}
			if ((clbits & 1) != 0) || ((clbits & 0x3f00) != 0) {
				x = ox
				y = oy
				self.dig.drawing.drawGold(bag, 0, ox, oy)
				self.dig.main.incrementPenalty()
				push = false
			}
		}
		if push {
			self.bagdat[bag].direction = dir
		} else {
			self.bagdat[bag].direction = self.dig.reversedir(dir)
		}
		self.bagdat[bag].x = x
		self.bagdat[bag].y = y
		self.bagdat[bag].h = ((x - 12) / 20)
		self.bagdat[bag].v = ((y - 18) / 18)
		self.bagdat[bag].xr = ((x - 12) % 20)
		self.bagdat[bag].yr = ((y - 18) % 18)
	}
	return push
}
func (self *Bags) pushbags(dir int, bits int) bool {
	var push bool = true
	_ = push
	{
		var bag int = 1
		_ = bag
		var bit int = 2
		_ = bit
		for ; bag < Bags_MAX_BAGS; func() { postInc(&bag); bit <<= 1 }() {
			if (bits & bit) != 0 {
				if !(self.pushbag(bag, dir)) {
					push = false
				}
			}
		}
	}
	return push
}
func (self *Bags) pushudbags(bits int) bool {
	var push bool = true
	_ = push
	{
		var bag int = 1
		_ = bag
		var b int = 2
		_ = b
		for ; bag < Bags_MAX_BAGS; func() { postInc(&bag); b <<= 1 }() {
			if (bits & b) != 0 {
				if self.bagdat[bag].goldTime != 0 {
					self.getgold(bag)
				} else {
					push = false
				}
			}
		}
	}
	return push
}
func (self *Bags) removebag(bag int) {
	if self.bagdat[bag].exist {
		self.bagdat[bag].exist = false
		self.dig.sprite.eraseSprite(bag)
	}
}
func (self *Bags) removebags(bits int) {
	{
		var bag int = 1
		_ = bag
		var b int = 2
		_ = b
		for ; bag < Bags_MAX_BAGS; func() { postInc(&bag); b <<= 1 }() {
			if self.bagdat[bag].exist && ((bits & b) != 0) {
				self.removebag(bag)
			}
		}
	}
}
func (self *Bags) updatebag(bag int) {
	var x int
	_ = x
	var h int
	_ = h
	var xr int
	_ = xr
	var y int
	_ = y
	var v int
	_ = v
	var yr int
	_ = yr
	var wbl int
	_ = wbl
	x = self.bagdat[bag].x
	h = self.bagdat[bag].h
	xr = self.bagdat[bag].xr
	y = self.bagdat[bag].y
	v = self.bagdat[bag].v
	yr = self.bagdat[bag].yr
	switch self.bagdat[bag].direction {
	case -(1):
		if (y < 180) && (xr == 0) {
			if self.bagdat[bag].wobbling {
				if self.bagdat[bag].wobbleTime == 0 {
					self.bagdat[bag].direction = 6
					self.dig.sound.soundFall()
					break
				}
				self.bagdat[bag].wobbleTime--
				wbl = (self.bagdat[bag].wobbleTime % 8)
				if !((wbl & 1) != 0) {
					self.dig.drawing.drawGold(bag, self.wobbleAnim[(wbl>>1)], x, y)
					self.dig.main.incrementPenalty()
					self.dig.sound.soundWobble()
				}
			} else {
				if (self.dig.monster.getfield(h, (v+1)) & 0xfdf) != 0xfdf {
					if !(self.dig.isDiggerUnderBag(h, (v + 1))) {
						self.bagdat[bag].wobbling = true
					}
				}
			}
		} else {
			self.bagdat[bag].wobbleTime = 15
			self.bagdat[bag].wobbling = false
		}
	case 0, 4:
		if xr == 0 {
			if (y < 180) && ((self.dig.monster.getfield(h, (v+1)) & 0xfdf) != 0xfdf) {
				self.bagdat[bag].direction = 6
				self.bagdat[bag].wobbleTime = 0
				self.dig.sound.soundFall()
			} else {
				self.baghitground(bag)
			}
		}
	case 6:
		if yr == 0 {
			self.bagdat[bag].fallHeight++
		}
		if y >= 180 {
			self.baghitground(bag)
		} else {
			if (self.dig.monster.getfield(h, (v+1)) & 0xfdf) == 0xfdf {
				if yr == 0 {
					self.baghitground(bag)
				}
			}
		}
		self.dig.monster.checkMonScared(self.bagdat[bag].h)
	}
	if self.bagdat[bag].direction != -(1) {
		if (self.bagdat[bag].direction != 6) && (self.pushCount != 0) {
			self.pushCount--
		} else {
			self.pushbag(bag, self.bagdat[bag].direction)
		}
	}
}
