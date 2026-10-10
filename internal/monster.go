// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

const Monster_MAX_MONSTERS int = 6

type Monster struct {
	dig            *Digger
	mondat         []*MonsterState
	nextMonster    int
	totalMonsters  int
	maxMonOnScreen int
	nextMonTime    int
	monGapTime     int
	unbonusFlag    bool
	monGotGold     bool
}

func newMonster(parent *Digger) *Monster {
	self := &Monster{}
	self.mondat = make([]*MonsterState, Monster_MAX_MONSTERS)
	self.nextMonster = 0
	self.totalMonsters = 0
	self.maxMonOnScreen = 0
	self.nextMonTime = 0
	self.monGapTime = 0
	self.unbonusFlag = false
	self.monGotGold = false
	self.dig = parent
	{
		var i int = 0
		_ = i
		for ; i < Monster_MAX_MONSTERS; i++ {
			self.mondat[i] = newMonsterState()
		}
	}
	return self
}
func (self *Monster) checkcoincide(mon int, bits int) {
	{
		var m int = 0
		_ = m
		var b int = 256
		_ = b
		for ; m < Monster_MAX_MONSTERS; func() { postInc(&m); b <<= 1 }() {
			if ((((bits & b) != 0) && (self.mondat[mon].direction == self.mondat[m].direction)) && (self.mondat[m].spawnTime == 0)) && (self.mondat[mon].spawnTime == 0) {
				self.mondat[m].direction = self.dig.reversedir(self.mondat[m].direction)
			}
		}
	}
}
func (self *Monster) checkMonScared(h int) {
	{
		var m int = 0
		_ = m
		for ; m < Monster_MAX_MONSTERS; m++ {
			if (h == self.mondat[m].h) && (self.mondat[m].direction == 2) {
				self.mondat[m].direction = 6
			}
		}
	}
}
func (self *Monster) createMonster() {
	{
		var i int = 0
		_ = i
		for ; i < Monster_MAX_MONSTERS; i++ {
			if !(self.mondat[i].flag) {
				self.mondat[i].flag = true
				self.mondat[i].alive = true
				self.mondat[i].kind = 0
				self.mondat[i].nob = true
				self.mondat[i].huntTime = 0
				self.mondat[i].h = 14
				self.mondat[i].v = 0
				self.mondat[i].x = 292
				self.mondat[i].y = 18
				self.mondat[i].xr = 0
				self.mondat[i].yr = 0
				self.mondat[i].direction = 4
				self.mondat[i].horizontalDir = 4
				self.nextMonster++
				self.nextMonTime = self.monGapTime
				self.mondat[i].spawnTime = 5
				self.dig.sprite.moveDrawSprite((i + 8), self.mondat[i].x, self.mondat[i].y)
				break
			}
		}
	}
}
func (self *Monster) doMonsters() {
	if self.nextMonTime > 0 {
		self.nextMonTime--
	} else {
		if (((self.nextMonster < self.totalMonsters) && (self.getMonstersOnScreenCount() < self.maxMonOnScreen)) && self.dig.digOnScreen) && !(self.dig.bonusMode) {
			self.createMonster()
		}
		if (self.unbonusFlag && (self.nextMonster == self.totalMonsters)) && (self.nextMonTime == 0) {
			if self.dig.digOnScreen {
				self.unbonusFlag = false
				self.dig.createbonus()
			}
		}
	}
	{
		var i int = 0
		_ = i
		for ; i < Monster_MAX_MONSTERS; i++ {
			if self.mondat[i].flag {
				if self.mondat[i].huntTime > (10 - self.dig.main.getLevelNumberClampedToTen()) {
					if self.mondat[i].nob {
						self.mondat[i].nob = false
						self.mondat[i].huntTime = 0
					}
				}
				if self.mondat[i].alive {
					if self.mondat[i].kind == 0 {
						self.handleMonsterAi(i)
						if (self.dig.main.randomNumber((15 - self.dig.main.getLevelNumberClampedToTen())) == 0) && self.mondat[i].nob {
							self.handleMonsterAi(i)
						}
					} else {
						self.mondat[i].kind--
					}
				} else {
					self.handleMonsterDeath(i)
				}
			}
		}
	}
}
func (self *Monster) eraseMonsters() {
	{
		var i int = 0
		_ = i
		for ; i < Monster_MAX_MONSTERS; i++ {
			if self.mondat[i].flag {
				self.dig.sprite.eraseSprite((i + 8))
			}
		}
	}
}
func (self *Monster) fieldClear(dir int, x int, y int) bool {
	switch dir {
	case 0:
		if x < 14 {
			if (self.getfield((x+1), y) & 0x2000) == 0 {
				if ((self.getfield((x+1), y) & 1) == 0) || ((self.getfield(x, y) & 0x10) == 0) {
					return true
				}
			}
		}
	case 4:
		if x > 0 {
			if (self.getfield((x-1), y) & 0x2000) == 0 {
				if ((self.getfield((x-1), y) & 0x10) == 0) || ((self.getfield(x, y) & 1) == 0) {
					return true
				}
			}
		}
	case 2:
		if y > 0 {
			if (self.getfield(x, (y-1)) & 0x2000) == 0 {
				if ((self.getfield(x, (y-1)) & 0x800) == 0) || ((self.getfield(x, y) & 0x40) == 0) {
					return true
				}
			}
		}
	case 6:
		if y < 9 {
			if (self.getfield(x, (y+1)) & 0x2000) == 0 {
				if ((self.getfield(x, (y+1)) & 0x40) == 0) || ((self.getfield(x, y) & 0x800) == 0) {
					return true
				}
			}
		}
	}
	return false
}
func (self *Monster) getfield(x int, y int) int {
	return self.dig.drawing.field[((y * 15) + x)]
}
func (self *Monster) increaseMonsterDelay(n int) {
	if n > 6 {
		n = 6
	}
	{
		var m int = 1
		_ = m
		for ; m < n; m++ {
			self.mondat[m].kind++
		}
	}
}
func (self *Monster) incpenalties(bits int) {
	{
		var m int = 0
		_ = m
		var b int = 256
		_ = b
		for ; m < Monster_MAX_MONSTERS; func() { postInc(&m); b <<= 1 }() {
			if (bits & b) != 0 {
				self.dig.main.incrementPenalty()
			}
			b <<= 1
		}
	}
}
func (self *Monster) initMonsters() {
	{
		var i int = 0
		_ = i
		for ; i < Monster_MAX_MONSTERS; i++ {
			self.mondat[i].flag = false
		}
	}
	self.nextMonster = 0
	self.monGapTime = (45 - (self.dig.main.getLevelNumberClampedToTen() << 1))
	self.totalMonsters = (self.dig.main.getLevelNumberClampedToTen() + 5)
	switch self.dig.main.getLevelNumberClampedToTen() {
	case 1:
		self.maxMonOnScreen = 3
	case 2, 3, 4, 5, 6, 7:
		self.maxMonOnScreen = 4
	case 8, 9, 10:
		self.maxMonOnScreen = 5
	}
	self.nextMonTime = 10
	self.unbonusFlag = true
}
func (self *Monster) killMonster(mon int) {
	if self.mondat[mon].flag {
		self.mondat[mon].alive = false
		self.mondat[mon].flag = self.mondat[mon].alive
		self.dig.sprite.eraseSprite((mon + 8))
		if self.dig.bonusMode {
			self.totalMonsters++
		}
	}
}
func (self *Monster) killmonsters(bits int) int {
	var n int = 0
	_ = n
	{
		var m int = 0
		_ = m
		var b int = 256
		_ = b
		for ; m < Monster_MAX_MONSTERS; func() { postInc(&m); b <<= 1 }() {
			if (bits & b) != 0 {
				self.killMonster(m)
				n++
			}
		}
	}
	return n
}
func (self *Monster) handleMonsterAi(mon int) {
	var clbits int
	_ = clbits
	var monox int
	_ = monox
	var monoy int
	_ = monoy
	var dir int
	_ = dir
	var mdirp1 int
	_ = mdirp1
	var mdirp2 int
	_ = mdirp2
	var mdirp3 int
	_ = mdirp3
	var mdirp4 int
	_ = mdirp4
	var t int
	_ = t
	var push bool
	_ = push
	monox = self.mondat[mon].x
	monoy = self.mondat[mon].y
	if (self.mondat[mon].xr == 0) && (self.mondat[mon].yr == 0) {
		if self.mondat[mon].huntTime > (30 + (self.dig.main.getLevelNumberClampedToTen() << 1)) {
			if !(self.mondat[mon].nob) {
				self.mondat[mon].huntTime = 0
				self.mondat[mon].nob = true
			}
		}
		if abs((self.dig.diggery - self.mondat[mon].y)) > abs((self.dig.diggerx - self.mondat[mon].x)) {
			mdirp1 = choose((self.dig.diggery < self.mondat[mon].y), 2, 6)
			mdirp4 = choose((self.dig.diggery < self.mondat[mon].y), 6, 2)
			mdirp2 = choose((self.dig.diggerx < self.mondat[mon].x), 4, 0)
			mdirp3 = choose((self.dig.diggerx < self.mondat[mon].x), 0, 4)
		} else {
			mdirp1 = choose((self.dig.diggerx < self.mondat[mon].x), 4, 0)
			mdirp4 = choose((self.dig.diggerx < self.mondat[mon].x), 0, 4)
			mdirp2 = choose((self.dig.diggery < self.mondat[mon].y), 2, 6)
			mdirp3 = choose((self.dig.diggery < self.mondat[mon].y), 6, 2)
		}
		if self.dig.bonusMode {
			t = mdirp1
			mdirp1 = mdirp4
			mdirp4 = t
			t = mdirp2
			mdirp2 = mdirp3
			mdirp3 = t
		}
		dir = self.dig.reversedir(self.mondat[mon].direction)
		if dir == mdirp1 {
			mdirp1 = mdirp2
			mdirp2 = mdirp3
			mdirp3 = mdirp4
			mdirp4 = dir
		}
		if dir == mdirp2 {
			mdirp2 = mdirp3
			mdirp3 = mdirp4
			mdirp4 = dir
		}
		if dir == mdirp3 {
			mdirp3 = mdirp4
			mdirp4 = dir
		}
		if (self.dig.main.randomNumber((self.dig.main.getLevelNumberClampedToTen() + 5)) == 1) && (self.dig.main.getLevelNumberClampedToTen() < 6) {
			t = mdirp1
			mdirp1 = mdirp3
			mdirp3 = t
		}
		if self.fieldClear(mdirp1, self.mondat[mon].h, self.mondat[mon].v) {
			dir = mdirp1
		} else {
			if self.fieldClear(mdirp2, self.mondat[mon].h, self.mondat[mon].v) {
				dir = mdirp2
			} else {
				if self.fieldClear(mdirp3, self.mondat[mon].h, self.mondat[mon].v) {
					dir = mdirp3
				} else {
					if self.fieldClear(mdirp4, self.mondat[mon].h, self.mondat[mon].v) {
						dir = mdirp4
					}
				}
			}
		}
		if !(self.mondat[mon].nob) {
			dir = mdirp1
		}
		if self.mondat[mon].direction != dir {
			self.mondat[mon].kind++
		}
		self.mondat[mon].direction = dir
	}
	if ((((self.mondat[mon].x == 292) && (self.mondat[mon].direction == 0)) || ((self.mondat[mon].x == 12) && (self.mondat[mon].direction == 4))) || ((self.mondat[mon].y == 180) && (self.mondat[mon].direction == 6))) || ((self.mondat[mon].y == 18) && (self.mondat[mon].direction == 2)) {
		self.mondat[mon].direction = -(1)
	}
	if (self.mondat[mon].direction == 4) || (self.mondat[mon].direction == 0) {
		self.mondat[mon].horizontalDir = self.mondat[mon].direction
	}
	if !(self.mondat[mon].nob) {
		self.dig.drawing.digTunnel(self.mondat[mon].x, self.mondat[mon].y, self.mondat[mon].direction)
	}
	switch self.mondat[mon].direction {
	case 0:
		if !(self.mondat[mon].nob) {
			self.dig.drawing.drawTunnelEdgeRight(self.mondat[mon].x, self.mondat[mon].y)
		}
		self.mondat[mon].x += 4
	case 4:
		if !(self.mondat[mon].nob) {
			self.dig.drawing.drawTunnelEdgeLeft(self.mondat[mon].x, self.mondat[mon].y)
		}
		self.mondat[mon].x -= 4
	case 2:
		if !(self.mondat[mon].nob) {
			self.dig.drawing.drawTunnelEdgeTop(self.mondat[mon].x, self.mondat[mon].y)
		}
		self.mondat[mon].y -= 3
	case 6:
		if !(self.mondat[mon].nob) {
			self.dig.drawing.drawTunnelEdgeBottom(self.mondat[mon].x, self.mondat[mon].y)
		}
		self.mondat[mon].y += 3
	}
	if !(self.mondat[mon].nob) {
		self.dig.hitemerald(((self.mondat[mon].x - 12) / 20), ((self.mondat[mon].y - 18) / 18), ((self.mondat[mon].x - 12) % 20), ((self.mondat[mon].y - 18) % 18), self.mondat[mon].direction)
	}
	if !(self.dig.digOnScreen) {
		self.mondat[mon].x = monox
		self.mondat[mon].y = monoy
	}
	if self.mondat[mon].spawnTime != 0 {
		self.mondat[mon].spawnTime--
		self.mondat[mon].x = monox
		self.mondat[mon].y = monoy
	}
	if !(self.mondat[mon].nob) && (self.mondat[mon].huntTime < 100) {
		self.mondat[mon].huntTime++
	}
	push = true
	clbits = self.dig.drawing.drawMonster(mon, self.mondat[mon].nob, self.mondat[mon].horizontalDir, self.mondat[mon].x, self.mondat[mon].y)
	self.dig.main.incrementPenalty()
	if (clbits & 0x3f00) != 0 {
		self.mondat[mon].kind++
		self.checkcoincide(mon, clbits)
		self.incpenalties(clbits)
	}
	if (clbits & self.dig.bags.bagbits()) != 0 {
		self.mondat[mon].kind++
		self.monGotGold = false
		if (self.mondat[mon].direction == 4) || (self.mondat[mon].direction == 0) {
			push = self.dig.bags.pushbags(self.mondat[mon].direction, clbits)
			self.mondat[mon].kind++
		} else {
			if !(self.dig.bags.pushudbags(clbits)) {
				push = false
			}
		}
		if self.monGotGold {
			self.mondat[mon].kind = 0
		}
		if !(self.mondat[mon].nob) && (self.mondat[mon].huntTime > 1) {
			self.dig.bags.removebags(clbits)
		}
	}
	if (self.mondat[mon].nob && ((clbits & 0x3f00) != 0)) && self.dig.digOnScreen {
		self.mondat[mon].huntTime++
	}
	if !(push) {
		self.mondat[mon].x = monox
		self.mondat[mon].y = monoy
		self.dig.drawing.drawMonster(mon, self.mondat[mon].nob, self.mondat[mon].horizontalDir, self.mondat[mon].x, self.mondat[mon].y)
		self.dig.main.incrementPenalty()
		if self.mondat[mon].nob {
			self.mondat[mon].huntTime++
		}
		if ((self.mondat[mon].direction == 2) || (self.mondat[mon].direction == 6)) && self.mondat[mon].nob {
			self.mondat[mon].direction = self.dig.reversedir(self.mondat[mon].direction)
		}
	}
	if ((clbits & 1) != 0) && self.dig.digOnScreen {
		if self.dig.bonusMode {
			self.killMonster(mon)
			self.dig.scores.scoreEatMonster()
			self.dig.sound.soundEatm()
		} else {
			self.dig.killdigger(3, 0)
		}
	}
	self.mondat[mon].h = ((self.mondat[mon].x - 12) / 20)
	self.mondat[mon].v = ((self.mondat[mon].y - 18) / 18)
	self.mondat[mon].xr = ((self.mondat[mon].x - 12) % 20)
	self.mondat[mon].yr = ((self.mondat[mon].y - 18) % 18)
}
func (self *Monster) handleMonsterDeath(mon int) {
	switch self.mondat[mon].death {
	case 1:
		if (self.dig.bags.bagy(self.mondat[mon].bag) + 6) > self.mondat[mon].y {
			self.mondat[mon].y = self.dig.bags.bagy(self.mondat[mon].bag)
		}
		self.dig.drawing.drawMonsterDeath(mon, self.mondat[mon].nob, self.mondat[mon].horizontalDir, self.mondat[mon].x, self.mondat[mon].y)
		self.dig.main.incrementPenalty()
		if self.dig.bags.getbagdir(self.mondat[mon].bag) == -(1) {
			self.mondat[mon].deathTime = 1
			self.mondat[mon].death = 4
		}
	case 4:
		if self.mondat[mon].deathTime != 0 {
			self.mondat[mon].deathTime--
		} else {
			self.killMonster(mon)
			self.dig.scores.scoreKillMonster()
		}
	}
}
func (self *Monster) monsterGotGold() {
	self.monGotGold = true
}
func (self *Monster) getMonstersLeft() int {
	return ((self.getMonstersOnScreenCount() + self.totalMonsters) - self.nextMonster)
}
func (self *Monster) getMonstersOnScreenCount() int {
	var n int = 0
	_ = n
	{
		var i int = 0
		_ = i
		for ; i < Monster_MAX_MONSTERS; i++ {
			if self.mondat[i].flag {
				n++
			}
		}
	}
	return n
}
func (self *Monster) squashmonster(mon int, death int, bag int) {
	self.mondat[mon].alive = false
	self.mondat[mon].death = death
	self.mondat[mon].bag = bag
}
func (self *Monster) squashmonsters(bag int, bits int) {
	{
		var m int = 0
		_ = m
		var b int = 256
		_ = b
		for ; m < Monster_MAX_MONSTERS; func() { postInc(&m); b <<= 1 }() {
			if (bits & b) != 0 {
				if self.mondat[m].y >= self.dig.bags.bagy(bag) {
					self.squashmonster(m, 1, bag)
				}
			}
		}
	}
}
