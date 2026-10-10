// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

type BagState struct {
	x          int
	y          int
	h          int
	v          int
	xr         int
	yr         int
	direction  int
	wobbleTime int
	goldTime   int
	fallHeight int
	wobbling   bool
	unfallen   bool
	exist      bool
}

func newBagState() *BagState {
	self := &BagState{}
	return self
}
func (self *BagState) copyFrom(other *BagState) {
	self.x = other.x
	self.y = other.y
	self.h = other.h
	self.v = other.v
	self.xr = other.xr
	self.yr = other.yr
	self.direction = other.direction
	self.wobbleTime = other.wobbleTime
	self.goldTime = other.goldTime
	self.fallHeight = other.fallHeight
	self.wobbling = other.wobbling
	self.unfallen = other.unfallen
	self.exist = other.exist
}
