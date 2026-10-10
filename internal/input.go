// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

type Input struct {
	dig              *Digger
	leftPressed      bool
	rightPressed     bool
	upPressed        bool
	downPressed      bool
	f1Pressed        bool
	firePressed      bool
	plusPressed      bool
	minusPressed     bool
	escape           bool
	lastKeyCode      int
	lastAsciiKeyCode int
	dynamicDir       int
	staticDir        int
	keyDir           int
	firePressedFlag  bool
	joyFlag          bool
}

func newInput(parent *Digger) *Input {
	self := &Input{}
	self.leftPressed = false
	self.rightPressed = false
	self.upPressed = false
	self.downPressed = false
	self.f1Pressed = false
	self.firePressed = false
	self.plusPressed = false
	self.minusPressed = false
	self.escape = false
	self.lastKeyCode = 0
	self.dynamicDir = -(1)
	self.staticDir = -(1)
	self.keyDir = 0
	self.firePressedFlag = false
	self.joyFlag = false
	self.dig = parent
	return self
}
func (self *Input) checkKeyboard() {
	if self.plusPressed {
		if self.dig.frametime > Digger_MIN_RATE {
			self.dig.frametime -= 5
		}
		self.plusPressed = false
	}
	if self.minusPressed {
		if self.dig.frametime < Digger_MAX_RATE {
			self.dig.frametime += 5
		}
		self.minusPressed = false
	}
}
func (self *Input) detectJoystick() {
	self.joyFlag = false
	self.dynamicDir = -(1)
	self.staticDir = self.dynamicDir
}
func (self *Input) getAsciiKey(make int) int {
	if (((make == 32) || ((make >= 97) && (make <= 122))) || ((make >= 65) && (make <= 90))) || ((make >= 48) && (make <= 57)) {
		return make
	} else {
		return 0
	}
}
func (self *Input) getDirection() int {
	return self.keyDir
}
func (self *Input) initKeyboard() {
}
func (self *Input) keyDownPressed() {
	self.downPressed = true
	self.staticDir = 6
	self.dynamicDir = self.staticDir
}
func (self *Input) keyDownReleased() {
	self.downPressed = false
	if self.dynamicDir == 6 {
		self.setDirection()
	}
}
func (self *Input) keyF1Pressed() {
	self.firePressed = true
	self.f1Pressed = true
}
func (self *Input) keyF1Released() {
	self.f1Pressed = false
}
func (self *Input) keyLeftPressed() {
	self.leftPressed = true
	self.staticDir = 4
	self.dynamicDir = self.staticDir
}
func (self *Input) keyLeftReleased() {
	self.leftPressed = false
	if self.dynamicDir == 4 {
		self.setDirection()
	}
}
func (self *Input) keyRightPressed() {
	self.rightPressed = true
	self.staticDir = 0
	self.dynamicDir = self.staticDir
}
func (self *Input) keyRightReleased() {
	self.rightPressed = false
	if self.dynamicDir == 0 {
		self.setDirection()
	}
}
func (self *Input) keyUpPressed() {
	self.upPressed = true
	self.staticDir = 2
	self.dynamicDir = self.staticDir
}
func (self *Input) keyUpReleased() {
	self.upPressed = false
	if self.dynamicDir == 2 {
		self.setDirection()
	}
}
func (self *Input) processKey(key int) {
	self.lastKeyCode = key
	self.lastAsciiKeyCode = key
}
func (self *Input) readDirection() {
	if self.dynamicDir != -(1) {
		self.keyDir = self.dynamicDir
	} else {
		if self.upPressed {
			self.keyDir = 2
		} else {
			if self.downPressed {
				self.keyDir = 6
			} else {
				if self.leftPressed {
					self.keyDir = 4
				} else {
					if self.rightPressed {
						self.keyDir = 0
					} else {
						self.keyDir = -(1)
					}
				}
			}
		}
	}
	self.firePressedFlag = (self.f1Pressed || self.firePressed)
	self.firePressed = false
}
func (self *Input) setDirection() {
	self.dynamicDir = -(1)
	if self.upPressed {
		self.staticDir = 2
		self.dynamicDir = self.staticDir
	}
	if self.downPressed {
		self.staticDir = 6
		self.dynamicDir = self.staticDir
	}
	if self.leftPressed {
		self.staticDir = 4
		self.dynamicDir = self.staticDir
	}
	if self.rightPressed {
		self.staticDir = 0
		self.dynamicDir = self.staticDir
	}
}
func (self *Input) testStart() bool {
	if ((self.lastKeyCode != 0) && ((self.lastKeyCode & 0x80) == 0)) && (self.lastKeyCode != 27) {
		self.joyFlag = false
		self.lastKeyCode = 0
		return true
	}
	return false
}
