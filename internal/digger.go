// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

const Digger_MAX_RATE int = 200
const Digger_MIN_RATE int = 40

type Digger struct {
	host *Engine

	running              bool
	width                int
	height               int
	frametime            int
	bags                 *Bags
	main                 *Main
	sound                *Sound
	monster              *Monster
	scores               *Scores
	sprite               *SpriteEngine
	drawing              *Drawing
	input                *Input
	display              *CgaDisplay
	diggerx              int
	diggery              int
	diggerh              int
	diggerv              int
	diggerrx             int
	diggerry             int
	digmdir              int
	digdir               int
	digtime              int
	rechargetime         int
	firex                int
	firey                int
	firedir              int
	expsn                int
	deathstage           int
	deathbag             int
	deathani             int
	deathtime            int
	startbonustimeleft   int
	bonustimeleft        int
	monsterEatMultiplier int
	emocttime            int
	emmask               int
	emfield              []int
	digOnScreen          bool
	notFiring            bool
	bonusVisible         bool
	bonusMode            bool
	diggerVisible        bool
	time                 int
	ftime                int
	embox                []int
	deatharc             []int
}

func newDiggerFields() *Digger {
	self := &Digger{}
	self.width = 320
	self.height = 200
	self.frametime = 66
	self.diggerx = 0
	self.diggery = 0
	self.diggerh = 0
	self.diggerv = 0
	self.diggerrx = 0
	self.diggerry = 0
	self.digmdir = 0
	self.digdir = 0
	self.digtime = 0
	self.rechargetime = 0
	self.firex = 0
	self.firey = 0
	self.firedir = 0
	self.expsn = 0
	self.deathstage = 0
	self.deathbag = 0
	self.deathani = 0
	self.deathtime = 0
	self.startbonustimeleft = 0
	self.bonustimeleft = 0
	self.monsterEatMultiplier = 0
	self.emocttime = 0
	self.emmask = 0
	self.emfield = make([]int, 150)
	self.digOnScreen = false
	self.notFiring = false
	self.bonusVisible = false
	self.bonusMode = false
	self.diggerVisible = false
	self.ftime = 50
	self.embox = []int{8, 12, 12, 9, 16, 12, 6, 9}
	self.deatharc = []int{3, 5, 6, 6, 5, 3, 0}
	return self
}
func (self *Digger) isDiggerUnderBag(h int, v int) bool {
	if (self.digmdir == 2) || (self.digmdir == 6) {
		if ((self.diggerx - 12) / 20) == h {
			if (((self.diggery - 18) / 18) == v) || ((((self.diggery - 18) / 18) + 1) == v) {
				return true
			}
		}
	}
	return false
}
func (self *Digger) countEmeralds() int {
	var n int = 0
	_ = n
	{
		var x int = 0
		_ = x
		for ; x < 15; x++ {
			{
				var y int = 0
				_ = y
				for ; y < 10; y++ {
					if (self.emfield[((y*15)+x)] & self.emmask) != 0 {
						n++
					}
				}
			}
		}
	}
	return n
}
func (self *Digger) createbonus() {
	self.bonusVisible = true
	self.drawing.drawBonus(292, 18)
}
func (self *Digger) diggerdie() {
	var clbits int
	_ = clbits
	switch self.deathstage {
	case 1:
		if (self.bags.bagy(self.deathbag) + 6) > self.diggery {
			self.diggery = (self.bags.bagy(self.deathbag) + 6)
		}
		self.drawing.drawDigger(15, self.diggerx, self.diggery, false)
		self.main.incrementPenalty()
		if (self.bags.getbagdir(self.deathbag) + 1) == 0 {
			self.sound.soundDdie()
			self.deathtime = 5
			self.deathstage = 2
			self.deathani = 0
			self.diggery -= 6
		}
	case 2:
		if self.deathtime != 0 {
			self.deathtime--
			break
		}
		if self.deathani == 0 {
			self.sound.music(2)
		}
		clbits = self.drawing.drawDigger((14 - self.deathani), self.diggerx, self.diggery, false)
		self.main.incrementPenalty()
		if (self.deathani == 0) && ((clbits & 0x3f00) != 0) {
			self.monster.killmonsters(clbits)
		}
		if self.deathani < 4 {
			self.deathani++
			self.deathtime = 2
		} else {
			self.deathstage = 4
			if self.sound.musicFlag {
				self.deathtime = 60
			} else {
				self.deathtime = 10
			}
		}
	case 3:
		self.deathstage = 5
		self.deathani = 0
		self.deathtime = 0
	case 5:
		if (self.deathani >= 0) && (self.deathani <= 6) {
			self.drawing.drawDigger(15, self.diggerx, (self.diggery - self.deatharc[self.deathani]), false)
			if self.deathani == 6 {
				self.sound.musicOff()
			}
			self.main.incrementPenalty()
			self.deathani++
			if self.deathani == 1 {
				self.sound.soundDdie()
			}
			if self.deathani == 7 {
				self.deathtime = 5
				self.deathani = 0
				self.deathstage = 2
			}
		}
	case 4:
		if self.deathtime != 0 {
			self.deathtime--
		} else {
			self.main.setDead(true)
		}
	}
}
func (self *Digger) doDigger() {
	self.newFrame()
	if self.expsn != 0 {
		self.drawexplosion()
	} else {
		self.updatefire()
	}
	if self.diggerVisible {
		if self.digOnScreen {
			if self.digtime != 0 {
				self.drawing.drawDigger(self.digmdir, self.diggerx, self.diggery, (self.notFiring && (self.rechargetime == 0)))
				self.main.incrementPenalty()
				self.digtime--
			} else {
				self.updatedigger()
			}
		} else {
			self.diggerdie()
		}
	}
	if self.bonusMode && self.digOnScreen {
		if self.bonustimeleft != 0 {
			self.bonustimeleft--
			if (self.startbonustimeleft != 0) || (self.bonustimeleft < 20) {
				self.startbonustimeleft--
				if (self.bonustimeleft & 1) != 0 {
					self.display.setIntensity(1)
					self.sound.soundBonus()
				} else {
					self.display.setIntensity(0)
					self.sound.soundBonus()
				}
				if self.startbonustimeleft == 0 {
					self.sound.music(0)
					self.sound.soundBonusOff()
					self.display.setIntensity(0)
				}
			}
		} else {
			self.endbonusmode()
			self.sound.soundBonusOff()
			self.sound.music(1)
		}
	}
	if self.bonusMode && !(self.digOnScreen) {
		self.endbonusmode()
		self.sound.soundBonusOff()
		self.sound.music(1)
	}
	if self.emocttime > 0 {
		self.emocttime--
	}
}
func (self *Digger) drawEmeralds() {
	self.emmask = (1 << self.main.getCurrentPlayer())
	{
		var x int = 0
		_ = x
		for ; x < 15; x++ {
			{
				var y int = 0
				_ = y
				for ; y < 10; y++ {
					if (self.emfield[((y*15)+x)] & self.emmask) != 0 {
						self.drawing.drawEmerald(((x * 20) + 12), ((y * 18) + 21))
					}
				}
			}
		}
	}
}
func (self *Digger) drawexplosion() {
	switch self.expsn {
	case 1:
		self.sound.soundExplode()
		fallthrough
	case 2, 3:
		self.drawing.drawFire(self.firex, self.firey, self.expsn)
		self.main.incrementPenalty()
		self.expsn++
	default:
		self.killFire()
		self.expsn = 0
	}
}
func (self *Digger) endbonusmode() {
	self.bonusMode = false
	self.display.setIntensity(0)
}
func (self *Digger) eraseBonus() {
	if self.bonusVisible {
		self.bonusVisible = false
		self.sprite.eraseSprite(14)
	}
	self.display.setIntensity(0)
}
func (self *Digger) eraseDigger() {
	self.sprite.eraseSprite(0)
	self.diggerVisible = false
}
func (self *Digger) getfirepflag() bool {
	return self.input.firePressedFlag
}
func (self *Digger) hitemerald(x int, y int, rx int, ry int, dir int) bool {
	var hit bool = false
	_ = hit
	var r int
	_ = r
	if ((dir < 0) || (dir > 6)) || ((dir & 1) != 0) {
		return hit
	}
	if (dir == 0) && (rx != 0) {
		x++
	}
	if (dir == 6) && (ry != 0) {
		y++
	}
	if (dir == 0) || (dir == 4) {
		r = rx
	} else {
		r = ry
	}
	if (self.emfield[((y*15)+x)] & self.emmask) != 0 {
		if r == self.embox[dir] {
			self.drawing.drawEmerald(((x * 20) + 12), ((y * 18) + 21))
			self.main.incrementPenalty()
		}
		if r == self.embox[(dir+1)] {
			self.drawing.eraseEmerald(((x * 20) + 12), ((y * 18) + 21))
			self.main.incrementPenalty()
			hit = true
			self.emfield[((y * 15) + x)] &= ^(self.emmask)
		}
	}
	return hit
}
func (self *Digger) initbonusmode() {
	self.bonusMode = true
	self.eraseBonus()
	self.display.setIntensity(1)
	self.bonustimeleft = (250 - (self.main.getLevelNumberClampedToTen() * 20))
	self.startbonustimeleft = 20
	self.monsterEatMultiplier = 1
}
func (self *Digger) initDigger() {
	self.diggerv = 9
	self.digmdir = 4
	self.diggerh = 7
	self.diggerx = ((self.diggerh * 20) + 12)
	self.digdir = 0
	self.diggerrx = 0
	self.diggerry = 0
	self.digtime = 0
	self.digOnScreen = true
	self.deathstage = 1
	self.diggerVisible = true
	self.diggery = ((self.diggerv * 18) + 18)
	self.sprite.moveDrawSprite(0, self.diggerx, self.diggery)
	self.notFiring = true
	self.emocttime = 0
	self.bonusMode = false
	self.bonusVisible = self.bonusMode
	self.input.firePressed = false
	self.expsn = 0
	self.rechargetime = 0
}
func (self *Digger) killdigger(stage int, bag int) {
	if (self.deathstage < 2) || (self.deathstage > 4) {
		self.digOnScreen = false
		self.deathstage = stage
		self.deathbag = bag
	}
}
func (self *Digger) killemerald(x int, y int) {
	if (self.emfield[(((y*15)+x)+15)] & self.emmask) != 0 {
		self.emfield[(((y * 15) + x) + 15)] &= ^(self.emmask)
		self.drawing.eraseEmerald(((x * 20) + 12), (((y + 1) * 18) + 21))
	}
}
func (self *Digger) killFire() {
	if !(self.notFiring) {
		self.notFiring = true
		self.sprite.eraseSprite(15)
		self.sound.soundFireOff()
	}
}
func (self *Digger) makeEmeraldField() {
	self.emmask = (1 << self.main.getCurrentPlayer())
	{
		var x int = 0
		_ = x
		for ; x < 15; x++ {
			{
				var y int = 0
				_ = y
				for ; y < 10; y++ {
					if self.main.getLevelChar(x, y, self.main.getLevelPlan()) == 67 {
						self.emfield[((y * 15) + x)] |= self.emmask
					} else {
						self.emfield[((y * 15) + x)] &= ^(self.emmask)
					}
				}
			}
		}
	}
}
func (self *Digger) reversedir(dir int) int {
	switch dir {
	case 0:
		return 4
	case 4:
		return 0
	case 2:
		return 6
	case 6:
		return 2
	}
	return dir
}
func (self *Digger) updatedigger() {
	var dir int
	_ = dir
	var ddir int
	_ = ddir
	var clbits int
	_ = clbits
	var diggerox int
	_ = diggerox
	var diggeroy int
	_ = diggeroy
	var nmon int
	_ = nmon
	var push bool = true
	_ = push
	self.input.readDirection()
	dir = self.input.getDirection()
	if (((dir == 0) || (dir == 2)) || (dir == 4)) || (dir == 6) {
		ddir = dir
	} else {
		ddir = -(1)
	}
	if (self.diggerrx == 0) && ((ddir == 2) || (ddir == 6)) {
		self.digmdir = ddir
		self.digdir = self.digmdir
	}
	if (self.diggerry == 0) && ((ddir == 4) || (ddir == 0)) {
		self.digmdir = ddir
		self.digdir = self.digmdir
	}
	if dir == -1 {
		self.digmdir = -1
	} else {
		self.digmdir = self.digdir
	}
	if ((((self.diggerx == 292) && (self.digmdir == 0)) || ((self.diggerx == 12) && (self.digmdir == 4))) || ((self.diggery == 180) && (self.digmdir == 6))) || ((self.diggery == 18) && (self.digmdir == 2)) {
		self.digmdir = -(1)
	}
	diggerox = self.diggerx
	diggeroy = self.diggery
	if self.digmdir != -(1) {
		self.drawing.digTunnel(diggerox, diggeroy, self.digmdir)
	}
	switch self.digmdir {
	case 0:
		self.drawing.drawTunnelEdgeRight(self.diggerx, self.diggery)
		self.diggerx += 4
	case 4:
		self.drawing.drawTunnelEdgeLeft(self.diggerx, self.diggery)
		self.diggerx -= 4
	case 2:
		self.drawing.drawTunnelEdgeTop(self.diggerx, self.diggery)
		self.diggery -= 3
	case 6:
		self.drawing.drawTunnelEdgeBottom(self.diggerx, self.diggery)
		self.diggery += 3
	}
	if self.hitemerald(((self.diggerx - 12) / 20), ((self.diggery - 18) / 18), ((self.diggerx - 12) % 20), ((self.diggery - 18) % 18), self.digmdir) {
		self.scores.scoreEmerald()
		self.sound.soundEm()
		self.sound.soundEmerald(self.emocttime)
		self.emocttime = 9
	}
	clbits = self.drawing.drawDigger(self.digdir, self.diggerx, self.diggery, (self.notFiring && (self.rechargetime == 0)))
	self.main.incrementPenalty()
	if (self.bags.bagbits() & clbits) != 0 {
		if (self.digmdir == 0) || (self.digmdir == 4) {
			push = self.bags.pushbags(self.digmdir, clbits)
			self.digtime++
		} else {
			if !(self.bags.pushudbags(clbits)) {
				push = false
			}
		}
		if !(push) {
			self.diggerx = diggerox
			self.diggery = diggeroy
			self.drawing.drawDigger(self.digmdir, self.diggerx, self.diggery, (self.notFiring && (self.rechargetime == 0)))
			self.main.incrementPenalty()
			self.digdir = self.reversedir(self.digmdir)
		}
	}
	if ((clbits & 0x3f00) != 0) && self.bonusMode {
		{
			nmon = self.monster.killmonsters(clbits)
			for ; nmon != 0; nmon-- {
				self.sound.soundEatm()
				self.scores.scoreEatMonster()
			}
		}
	}
	if (clbits & 0x4000) != 0 {
		self.scores.scoreBonus()
		self.initbonusmode()
	}
	self.diggerh = ((self.diggerx - 12) / 20)
	self.diggerrx = ((self.diggerx - 12) % 20)
	self.diggerv = ((self.diggery - 18) / 18)
	self.diggerry = ((self.diggery - 18) % 18)
}
func (self *Digger) updatefire() {
	var clbits int
	_ = clbits
	var b int
	_ = b
	var mon int
	_ = mon
	var pix int = 0
	_ = pix
	if self.notFiring {
		if self.rechargetime != 0 {
			self.rechargetime--
		} else {
			if self.getfirepflag() {
				if self.digOnScreen {
					self.rechargetime = ((self.main.getLevelNumberClampedToTen() * 3) + 60)
					self.notFiring = false
					switch self.digdir {
					case 0:
						self.firex = (self.diggerx + 8)
						self.firey = (self.diggery + 4)
					case 4:
						self.firex = self.diggerx
						self.firey = (self.diggery + 4)
					case 2:
						self.firex = (self.diggerx + 4)
						self.firey = self.diggery
					case 6:
						self.firex = (self.diggerx + 4)
						self.firey = (self.diggery + 8)
					}
					self.firedir = self.digdir
					self.sprite.moveDrawSprite(15, self.firex, self.firey)
					self.sound.soundFire()
				}
			}
		}
	} else {
		switch self.firedir {
		case 0:
			self.firex += 8
			pix = (self.display.getPixel(self.firex, (self.firey+4)) | self.display.getPixel((self.firex+4), (self.firey+4)))
		case 4:
			self.firex -= 8
			pix = (self.display.getPixel(self.firex, (self.firey+4)) | self.display.getPixel((self.firex+4), (self.firey+4)))
		case 2:
			self.firey -= 7
			pix = (((((((self.display.getPixel((self.firex+4), self.firey) | self.display.getPixel((self.firex+4), (self.firey+1))) | self.display.getPixel((self.firex+4), (self.firey+2))) | self.display.getPixel((self.firex+4), (self.firey+3))) | self.display.getPixel((self.firex+4), (self.firey+4))) | self.display.getPixel((self.firex+4), (self.firey+5))) | self.display.getPixel((self.firex+4), (self.firey+6))) & 0xc0)
		case 6:
			self.firey += 7
			pix = (((((((self.display.getPixel(self.firex, self.firey) | self.display.getPixel(self.firex, (self.firey+1))) | self.display.getPixel(self.firex, (self.firey+2))) | self.display.getPixel(self.firex, (self.firey+3))) | self.display.getPixel(self.firex, (self.firey+4))) | self.display.getPixel(self.firex, (self.firey+5))) | self.display.getPixel(self.firex, (self.firey+6))) & 3)
		}
		clbits = self.drawing.drawFire(self.firex, self.firey, 0)
		self.main.incrementPenalty()
		if (clbits & 0x3f00) != 0 {
			{
				mon = 0
				b = 256
				for ; mon < 6; func() { postInc(&mon); b <<= 1 }() {
					if (clbits & b) != 0 {
						self.monster.killMonster(mon)
						self.scores.scoreKillMonster()
						self.expsn = 1
					}
				}
			}
		}
		if (clbits & 0x40fe) != 0 {
			self.expsn = 1
		}
		switch self.firedir {
		case 0:
			if self.firex > 296 {
				self.expsn = 1
			} else {
				if (pix != 0) && (clbits == 0) {
					self.expsn = 1
					self.firex -= 8
					self.drawing.drawFire(self.firex, self.firey, 0)
				}
			}
		case 4:
			if self.firex < 16 {
				self.expsn = 1
			} else {
				if (pix != 0) && (clbits == 0) {
					self.expsn = 1
					self.firex += 8
					self.drawing.drawFire(self.firex, self.firey, 0)
				}
			}
		case 2:
			if self.firey < 15 {
				self.expsn = 1
			} else {
				if (pix != 0) && (clbits == 0) {
					self.expsn = 1
					self.firey += 7
					self.drawing.drawFire(self.firex, self.firey, 0)
				}
			}
		case 6:
			if self.firey > 183 {
				self.expsn = 1
			} else {
				if (pix != 0) && (clbits == 0) {
					self.expsn = 1
					self.firey -= 7
					self.drawing.drawFire(self.firex, self.firey, 0)
				}
			}
		}
	}
}
