// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

type CgaDisplay struct {
	intensity int

	dig    *Digger
	width  int
	height int
	size   int
	pixels []int
}

func newCgaDisplay(parent *Digger) *CgaDisplay {
	self := &CgaDisplay{}
	self.width = 320
	self.height = 200
	self.size = (self.width * self.height)
	self.dig = parent
	return self
}
func (self *CgaDisplay) readSpritePixels(x int, y int, p []int, w int, h int) {
	var src int = 0
	_ = src
	var dest int = ((y * self.width) + (x & 0xfffc))
	_ = dest
	{
		var i int = 0
		_ = i
		for ; i < h; i++ {
			var d int = dest
			_ = d
			{
				var j int = 0
				_ = j
				for ; j < w; j++ {
					p[postInc(&src)] = ((((((self.pixels[d] << 2) | self.pixels[(d+1)]) << 2) | self.pixels[(d+2)]) << 2) | self.pixels[(d+3)])
					d += 4
					if src == len(p) {
						return
					}
				}
			}
			dest += self.width
		}
	}
}
func (self *CgaDisplay) getPixel(x int, y int) int {
	var ofs int = ((self.width * y) + (x & 0xfffc))
	_ = ofs
	return ((((((self.pixels[ofs] << 2) | self.pixels[(ofs+1)]) << 2) | self.pixels[(ofs+2)]) << 2) | self.pixels[(ofs+3)])
}
func (self *CgaDisplay) init() {
}
func (self *CgaDisplay) drawSprite(x int, y int, p []int, w int, h int) {
	var src int = 0
	_ = src
	var dest int = ((y * self.width) + (x & 0xfffc))
	_ = dest
	{
		var i int = 0
		_ = i
		for ; i < h; i++ {
			var d int = dest
			_ = d
			{
				var j int = 0
				_ = j
				for ; j < w; j++ {
					var px int = p[postInc(&src)]
					_ = px
					self.pixels[(d + 3)] = (px & 3)
					px >>= 2
					self.pixels[(d + 2)] = (px & 3)
					px >>= 2
					self.pixels[(d + 1)] = (px & 3)
					px >>= 2
					self.pixels[d] = (px & 3)
					d += 4
					if src == len(p) {
						return
					}
				}
			}
			dest += self.width
		}
	}
}
func (self *CgaDisplay) drawSpriteMasked(x int, y int, ch int, w int, h int) {
	var spr []int = CgaGrafx_CGA_TABLE[(ch * 2)]
	_ = spr
	var msk []int = CgaGrafx_CGA_TABLE[((ch * 2) + 1)]
	_ = msk
	var src int = 0
	_ = src
	var dest int = ((y * self.width) + (x & 0xfffc))
	_ = dest
	{
		var i int = 0
		_ = i
		for ; i < h; i++ {
			var d int = dest
			_ = d
			{
				var j int = 0
				_ = j
				for ; j < w; j++ {
					var px int = spr[src]
					_ = px
					var mx int = msk[src]
					_ = mx
					src++
					if (mx & 3) == 0 {
						self.pixels[(d + 3)] = (px & 3)
					}
					px >>= 2
					if (mx & (3 << 2)) == 0 {
						self.pixels[(d + 2)] = (px & 3)
					}
					px >>= 2
					if (mx & (3 << 4)) == 0 {
						self.pixels[(d + 1)] = (px & 3)
					}
					px >>= 2
					if (mx & (3 << 6)) == 0 {
						self.pixels[d] = (px & 3)
					}
					d += 4
					if (src == len(spr)) || (src == len(msk)) {
						return
					}
				}
			}
			dest += self.width
		}
	}
}
func (self *CgaDisplay) drawTitleScreen() {
	var src int = 0
	_ = src
	var dest int = 0
	_ = dest
	for true {
		if src >= len(CgaGrafx_CGA_TITLE_DATA) {
			break
		}
		var b int = CgaGrafx_CGA_TITLE_DATA[postInc(&src)]
		_ = b
		var l int
		_ = l
		var c int
		_ = c
		if b == 0xfe {
			l = CgaGrafx_CGA_TITLE_DATA[postInc(&src)]
			if l == 0 {
				l = 256
			}
			c = CgaGrafx_CGA_TITLE_DATA[postInc(&src)]
		} else {
			l = 1
			c = b
		}
		{
			var i int = 0
			_ = i
			for ; i < l; i++ {
				var px int = c
				_ = px
				var adst int
				_ = adst
				if dest < 32768 {
					adst = (((dest / 320) * 640) + (dest % 320))
				} else {
					adst = ((320 + (((dest - 32768) / 320) * 640)) + ((dest - 32768) % 320))
				}
				self.pixels[(adst + 3)] = (px & 3)
				px >>= 2
				self.pixels[(adst + 2)] = (px & 3)
				px >>= 2
				self.pixels[(adst + 1)] = (px & 3)
				px >>= 2
				self.pixels[(adst + 0)] = (px & 3)
				dest += 4
				if dest >= 65535 {
					break
				}
			}
		}
		if dest >= 65535 {
			break
		}
	}
}
