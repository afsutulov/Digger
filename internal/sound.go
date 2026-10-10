// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

const Sound_SILENCE_T0VAL int = 0x7d00
const Sound_DEFAULT_T2VAL int = 40

type Sound struct {
	dig                  *Digger
	wavetype             int
	t2val                int
	t0val                int
	musvol               int
	spkrmode             int
	pulsewidth           int
	volume               int
	timerClock           int
	soundFlag            bool
	musicFlag            bool
	sndFlag              bool
	soundPausedFlag      bool
	soundLevDoneFlag     bool
	nljpointer           int
	nljnoteduration      int
	newlevjingle         []int
	soundFallFlag        bool
	soundFallF           bool
	soundfallvalue       int
	soundfalln           int
	soundBreakFlag       bool
	soundbreakduration   int
	soundbreakvalue      int
	soundWobbleFlag      bool
	soundwobblen         int
	soundFireFlag        bool
	soundfirevalue       int
	soundfiren           int
	soundExplodeFlag     bool
	soundexplodevalue    int
	soundexplodeduration int
	soundBonusFlag       bool
	soundbonusn          int
	soundEmFlag          bool
	soundEmeraldFlag     bool
	soundemeraldduration int
	emerfreq             int
	soundemeraldn        int
	soundGoldFlag        bool
	soundGoldF           bool
	soundgoldvalue1      int
	soundgoldvalue2      int
	soundgoldduration    int
	soundEatmFlag        bool
	soundeatmvalue       int
	soundeatmduration    int
	soundeatmn           int
	soundDdieFlag        bool
	soundddien           int
	soundddievalue       int
	sound1UpFlag         bool
	sound1upduration     int
	musicPlaying         bool
	musicp               int
	tuneno               int
	noteduration         int
	notevalue            int
	musicmaxvol          int
	musicattackrate      int
	musicsustainlevel    int
	musicdecayrate       int
	musicnotewidth       int
	musicreleaserate     int
	musicstage           int
	musicn               int
	bonusjingle          []int
	backgjingle          []int
	dirge                []int
	soundT0Flag          bool
	int8Flag             bool
}

func newSound(parent *Digger) *Sound {
	self := &Sound{}
	self.wavetype = 0
	self.t2val = 0
	self.t0val = 0
	self.musvol = 0
	self.spkrmode = 0
	self.pulsewidth = 1
	self.volume = 0
	self.timerClock = 0
	self.soundFlag = true
	self.musicFlag = true
	self.sndFlag = false
	self.soundPausedFlag = false
	self.soundLevDoneFlag = false
	self.nljpointer = 0
	self.nljnoteduration = 0
	self.newlevjingle = []int{0x8e8, 0x712, 0x5f2, 0x7f0, 0x6ac, 0x54c, 0x712, 0x5f2, 0x4b8, 0x474, 0x474}
	self.soundFallFlag = false
	self.soundFallF = false
	self.soundfalln = 0
	self.soundBreakFlag = false
	self.soundbreakduration = 0
	self.soundbreakvalue = 0
	self.soundWobbleFlag = false
	self.soundwobblen = 0
	self.soundFireFlag = false
	self.soundfiren = 0
	self.soundExplodeFlag = false
	self.soundBonusFlag = false
	self.soundbonusn = 0
	self.soundEmFlag = false
	self.soundEmeraldFlag = false
	self.soundGoldFlag = false
	self.soundGoldF = false
	self.soundEatmFlag = false
	self.soundDdieFlag = false
	self.sound1UpFlag = false
	self.sound1upduration = 0
	self.musicPlaying = false
	self.musicp = 0
	self.tuneno = 0
	self.noteduration = 0
	self.notevalue = 0
	self.musicmaxvol = 0
	self.musicattackrate = 0
	self.musicsustainlevel = 0
	self.musicdecayrate = 0
	self.musicnotewidth = 0
	self.musicreleaserate = 0
	self.musicstage = 0
	self.musicn = 0
	self.bonusjingle = []int{0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0xd59, 4, 0xbe4, 4, 0xa98, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0xd59, 2, 0xa98, 2, 0xbe4, 4, 0xe24, 4, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0xd59, 4, 0xbe4, 4, 0xa98, 4, 0xd59, 2, 0xa98, 2, 0x8e8, 10, 0xa00, 2, 0xa98, 2, 0xbe4, 2, 0xd59, 4, 0xa98, 4, 0xd59, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0xd59, 4, 0xbe4, 4, 0xa98, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0xd59, 2, 0xa98, 2, 0xbe4, 4, 0xe24, 4, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0x11d1, 4, 0x11d1, 2, 0x11d1, 2, 0xd59, 4, 0xbe4, 4, 0xa98, 4, 0xd59, 2, 0xa98, 2, 0x8e8, 10, 0xa00, 2, 0xa98, 2, 0xbe4, 2, 0xd59, 4, 0xa98, 4, 0xd59, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0xbe4, 4, 0xd59, 4, 0xe24, 4, 0xfdf, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0x8e8, 4, 0x970, 4, 0x8e8, 4, 0x970, 4, 0x8e8, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0xbe4, 4, 0xd59, 4, 0xe24, 4, 0xfdf, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0xa98, 2, 0xa98, 2, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0xa98, 4, 0x7f0, 4, 0x8e8, 4, 0x970, 4, 0x8e8, 4, 0x970, 4, 0x8e8, 4, 0x7d64}
	self.backgjingle = []int{0xfdf, 2, 0x11d1, 2, 0xfdf, 2, 0x1530, 2, 0x1ab2, 2, 0x1530, 2, 0x1fbf, 4, 0xfdf, 2, 0x11d1, 2, 0xfdf, 2, 0x1530, 2, 0x1ab2, 2, 0x1530, 2, 0x1fbf, 4, 0xfdf, 2, 0xe24, 2, 0xd59, 2, 0xe24, 2, 0xd59, 2, 0xfdf, 2, 0xe24, 2, 0xfdf, 2, 0xe24, 2, 0x11d1, 2, 0xfdf, 2, 0x11d1, 2, 0xfdf, 2, 0x1400, 2, 0xfdf, 4, 0xfdf, 2, 0x11d1, 2, 0xfdf, 2, 0x1530, 2, 0x1ab2, 2, 0x1530, 2, 0x1fbf, 4, 0xfdf, 2, 0x11d1, 2, 0xfdf, 2, 0x1530, 2, 0x1ab2, 2, 0x1530, 2, 0x1fbf, 4, 0xfdf, 2, 0xe24, 2, 0xd59, 2, 0xe24, 2, 0xd59, 2, 0xfdf, 2, 0xe24, 2, 0xfdf, 2, 0xe24, 2, 0x11d1, 2, 0xfdf, 2, 0x11d1, 2, 0xfdf, 2, 0xe24, 2, 0xd59, 4, 0xa98, 2, 0xbe4, 2, 0xa98, 2, 0xd59, 2, 0x11d1, 2, 0xd59, 2, 0x1530, 4, 0xa98, 2, 0xbe4, 2, 0xa98, 2, 0xd59, 2, 0x11d1, 2, 0xd59, 2, 0x1530, 4, 0xa98, 2, 0x970, 2, 0x8e8, 2, 0x970, 2, 0x8e8, 2, 0xa98, 2, 0x970, 2, 0xa98, 2, 0x970, 2, 0xbe4, 2, 0xa98, 2, 0xbe4, 2, 0xa98, 2, 0xd59, 2, 0xa98, 4, 0xa98, 2, 0xbe4, 2, 0xa98, 2, 0xd59, 2, 0x11d1, 2, 0xd59, 2, 0x1530, 4, 0xa98, 2, 0xbe4, 2, 0xa98, 2, 0xd59, 2, 0x11d1, 2, 0xd59, 2, 0x1530, 4, 0xa98, 2, 0x970, 2, 0x8e8, 2, 0x970, 2, 0x8e8, 2, 0xa98, 2, 0x970, 2, 0xa98, 2, 0x970, 2, 0xbe4, 2, 0xa98, 2, 0xbe4, 2, 0xa98, 2, 0xd59, 2, 0xa98, 4, 0x7f0, 2, 0x8e8, 2, 0xa98, 2, 0xd59, 2, 0x11d1, 2, 0xd59, 2, 0x1530, 4, 0xa98, 2, 0xbe4, 2, 0xa98, 2, 0xd59, 2, 0x11d1, 2, 0xd59, 2, 0x1530, 4, 0xa98, 2, 0x970, 2, 0x8e8, 2, 0x970, 2, 0x8e8, 2, 0xa98, 2, 0x970, 2, 0xa98, 2, 0x970, 2, 0xbe4, 2, 0xa98, 2, 0xbe4, 2, 0xd59, 2, 0xbe4, 2, 0xa98, 4, 0x7d64}
	self.dirge = []int{0x7d00, 0x2, 0x11d1, 0x6, 0x11d1, 0x4, 0x11d1, 0x2, 0x11d1, 0x6, 0xefb, 0x4, 0xfdf, 0x2, 0xfdf, 0x4, 0x11d1, 0x2, 0x11d1, 0x4, 0x12e0, 0x2, 0x11d1, 0xc, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d00, 0x10, 0x7d64}
	self.soundT0Flag = false
	self.int8Flag = false
	self.dig = parent
	return self
}
func (self *Sound) music(tune int) {
	self.tuneno = tune
	self.musicp = 0
	self.noteduration = 0
	switch tune {
	case 0:
		self.musicmaxvol = 50
		self.musicattackrate = 20
		self.musicsustainlevel = 20
		self.musicdecayrate = 10
		self.musicreleaserate = 4
	case 1:
		self.musicmaxvol = 50
		self.musicattackrate = 50
		self.musicsustainlevel = 8
		self.musicdecayrate = 15
		self.musicreleaserate = 1
	case 2:
		self.musicmaxvol = 50
		self.musicattackrate = 50
		self.musicsustainlevel = 25
		self.musicdecayrate = 5
		self.musicreleaserate = 1
	}
	self.musicPlaying = true
	if tune == 2 {
		self.soundDdieOff()
	}
}
func (self *Sound) musicOff() {
	self.musicPlaying = false
	self.musicp = 0
}
func (self *Sound) musicUpdate() {
	if !(self.musicPlaying) {
		return
	}
	if self.noteduration != 0 {
		self.noteduration--
	} else {
		self.musicn = 0
		self.musicstage = self.musicn
		switch self.tuneno {
		case 0:
			self.noteduration = (self.bonusjingle[(self.musicp+1)] * 3)
			self.musicnotewidth = (self.noteduration - 3)
			self.notevalue = self.bonusjingle[self.musicp]
			self.musicp += 2
			if self.bonusjingle[self.musicp] == 0x7d64 {
				self.musicp = 0
			}
		case 1:
			self.noteduration = (self.backgjingle[(self.musicp+1)] * 6)
			self.musicnotewidth = 12
			self.notevalue = self.backgjingle[self.musicp]
			self.musicp += 2
			if self.backgjingle[self.musicp] == 0x7d64 {
				self.musicp = 0
			}
		case 2:
			self.noteduration = (self.dirge[(self.musicp+1)] * 10)
			self.musicnotewidth = (self.noteduration - 10)
			self.notevalue = self.dirge[self.musicp]
			self.musicp += 2
			if self.dirge[self.musicp] == 0x7d64 {
				self.musicp = 0
			}
		}
	}
	self.musicn++
	self.wavetype = 1
	self.t0val = self.notevalue
	if self.musicn >= self.musicnotewidth {
		self.musicstage = 2
	}
	switch self.musicstage {
	case 0:
		if (self.musvol + self.musicattackrate) >= self.musicmaxvol {
			self.musicstage = 1
			self.musvol = self.musicmaxvol
			break
		}
		self.musvol += self.musicattackrate
	case 1:
		if (self.musvol - self.musicdecayrate) <= self.musicsustainlevel {
			self.musvol = self.musicsustainlevel
			break
		}
		self.musvol -= self.musicdecayrate
	case 2:
		if (self.musvol - self.musicreleaserate) <= 1 {
			self.musvol = 1
			break
		}
		self.musvol -= self.musicreleaserate
	}
	if self.musvol == 1 {
		self.t0val = Sound_SILENCE_T0VAL
	}
}
func (self *Sound) setSoundMode() {
	self.spkrmode = self.wavetype
	if !(self.soundT0Flag) && self.sndFlag {
		self.soundT0Flag = true
	}
}
func (self *Sound) setSoundT2() {
	if self.soundT0Flag {
		self.spkrmode = 0
		self.soundT0Flag = false
	}
}
func (self *Sound) setT0() {
	if self.sndFlag {
		if (self.t0val < 1000) && ((self.wavetype == 1) || (self.wavetype == 2)) {
			self.t0val = 1000
		}
		if self.musvol < 1 {
			self.musvol = 1
		}
		if self.musvol > 50 {
			self.musvol = 50
		}
		self.pulsewidth = (self.musvol * self.volume)
		self.setSoundMode()
	}
}
func (self *Sound) setT2Val(t2v int) {
	self.t2val = t2v
}
func (self *Sound) sound1Up() {
	self.sound1upduration = 96
	self.sound1UpFlag = true
}
func (self *Sound) sound1UpOff() {
	self.sound1UpFlag = false
}
func (self *Sound) sound1UpUpdate() {
	if self.sound1UpFlag {
		if ((self.sound1upduration / 3) % 2) != 0 {
			self.t2val = ((self.sound1upduration << 2) + 600)
		}
		self.sound1upduration--
		if self.sound1upduration < 1 {
			self.sound1UpFlag = false
		}
	}
}
func (self *Sound) soundBonus() {
	self.soundBonusFlag = true
}
func (self *Sound) soundBonusOff() {
	self.soundBonusFlag = false
	self.soundbonusn = 0
}
func (self *Sound) soundBonusUpdate() {
	if self.soundBonusFlag {
		self.soundbonusn++
		if self.soundbonusn > 15 {
			self.soundbonusn = 0
		}
		if (self.soundbonusn >= 0) && (self.soundbonusn < 6) {
			self.t2val = 0x4ce
		}
		if (self.soundbonusn >= 8) && (self.soundbonusn < 14) {
			self.t2val = 0x5e9
		}
	}
}
func (self *Sound) soundBreak() {
	self.soundbreakduration = 3
	if self.soundbreakvalue < 15000 {
		self.soundbreakvalue = 15000
	}
	self.soundBreakFlag = true
}
func (self *Sound) soundBreakOff() {
	self.soundBreakFlag = false
}
func (self *Sound) soundBreakUpdate() {
	if self.soundBreakFlag {
		if self.soundbreakduration != 0 {
			self.soundbreakduration--
			self.t2val = self.soundbreakvalue
		} else {
			self.soundBreakFlag = false
		}
	}
}
func (self *Sound) soundDdie() {
	self.soundddien = 0
	self.soundddievalue = 20000
	self.soundDdieFlag = true
}
func (self *Sound) soundDdieOff() {
	self.soundDdieFlag = false
}
func (self *Sound) soundDdieUpdate() {
	if self.soundDdieFlag {
		self.soundddien++
		if self.soundddien == 1 {
			self.musicOff()
		}
		if (self.soundddien >= 1) && (self.soundddien <= 10) {
			self.soundddievalue = (20000 - (self.soundddien * 1000))
		}
		if self.soundddien > 10 {
			self.soundddievalue += 500
		}
		if self.soundddievalue > 30000 {
			self.soundDdieOff()
		}
		self.t2val = self.soundddievalue
	}
}
func (self *Sound) soundEatm() {
	self.soundeatmduration = 20
	self.soundeatmn = 3
	self.soundeatmvalue = 2000
	self.soundEatmFlag = true
}
func (self *Sound) soundEatmOff() {
	self.soundEatmFlag = false
}
func (self *Sound) soundEatmUpdate() {
	if self.soundEatmFlag {
		if self.soundeatmn != 0 {
			if self.soundeatmduration != 0 {
				if (self.soundeatmduration % 4) == 1 {
					self.t2val = self.soundeatmvalue
				}
				if (self.soundeatmduration % 4) == 3 {
					self.t2val = (self.soundeatmvalue - (self.soundeatmvalue >> 4))
				}
				self.soundeatmduration--
				self.soundeatmvalue -= (self.soundeatmvalue >> 4)
			} else {
				self.soundeatmduration = 20
				self.soundeatmn--
				self.soundeatmvalue = 2000
			}
		} else {
			self.soundEatmFlag = false
		}
	}
}
func (self *Sound) soundEm() {
	self.soundEmFlag = true
}
func (self *Sound) soundEmerald(emocttime int) {
	if emocttime != 0 {
		switch self.emerfreq {
		case 0x8e8:
			self.emerfreq = 0x7f0
		case 0x7f0:
			self.emerfreq = 0x712
		case 0x712:
			self.emerfreq = 0x6ac
		case 0x6ac:
			self.emerfreq = 0x5f2
		case 0x5f2:
			self.emerfreq = 0x54c
		case 0x54c:
			self.emerfreq = 0x4b8
		case 0x4b8:
			self.emerfreq = 0x474
			self.dig.scores.scoreOctave()
		case 0x474:
			self.emerfreq = 0x8e8
		}
	} else {
		self.emerfreq = 0x8e8
	}
	self.soundemeraldduration = 7
	self.soundemeraldn = 0
	self.soundEmeraldFlag = true
}
func (self *Sound) soundEmeraldOff() {
	self.soundEmeraldFlag = false
}
func (self *Sound) soundEmeraldUpdate() {
	if self.soundEmeraldFlag {
		if self.soundemeraldduration != 0 {
			if (self.soundemeraldn == 0) || (self.soundemeraldn == 1) {
				self.t2val = self.emerfreq
			}
			self.soundemeraldn++
			if self.soundemeraldn > 7 {
				self.soundemeraldn = 0
				self.soundemeraldduration--
			}
		} else {
			self.soundEmeraldOff()
		}
	}
}
func (self *Sound) soundEmOff() {
	self.soundEmFlag = false
}
func (self *Sound) soundEmUpdate() {
	if self.soundEmFlag {
		self.t2val = 1000
		self.soundEmOff()
	}
}
func (self *Sound) soundExplode() {
	self.soundexplodevalue = 1500
	self.soundexplodeduration = 10
	self.soundExplodeFlag = true
	self.soundFireOff()
}
func (self *Sound) soundExplodeOff() {
	self.soundExplodeFlag = false
}
func (self *Sound) soundExplodeUpdate() {
	if self.soundExplodeFlag {
		if self.soundexplodeduration != 0 {
			self.t2val = (self.soundexplodevalue - (self.soundexplodevalue >> 3))
			self.soundexplodevalue = self.t2val
			self.soundexplodeduration--
		} else {
			self.soundExplodeFlag = false
		}
	}
}
func (self *Sound) soundFall() {
	self.soundfallvalue = 1000
	self.soundFallFlag = true
}
func (self *Sound) soundFallOff() {
	self.soundFallFlag = false
	self.soundfalln = 0
}
func (self *Sound) soundFallUpdate() {
	if self.soundFallFlag {
		if self.soundfalln < 1 {
			self.soundfalln++
			if self.soundFallF {
				self.t2val = self.soundfallvalue
			}
		} else {
			self.soundfalln = 0
			if self.soundFallF {
				self.soundfallvalue += 50
				self.soundFallF = false
			} else {
				self.soundFallF = true
			}
		}
	}
}
func (self *Sound) soundFire() {
	self.soundfirevalue = 500
	self.soundFireFlag = true
}
func (self *Sound) soundFireOff() {
	self.soundFireFlag = false
	self.soundfiren = 0
}
func (self *Sound) soundFireUpdate() {
	if self.soundFireFlag {
		if self.soundfiren == 1 {
			self.soundfiren = 0
			self.soundfirevalue += (self.soundfirevalue / 55)
			self.t2val = (self.soundfirevalue + self.dig.main.randomNumber((self.soundfirevalue >> 3)))
			if self.soundfirevalue > 30000 {
				self.soundFireOff()
			}
		} else {
			self.soundfiren++
		}
	}
}
func (self *Sound) soundGold() {
	self.soundgoldvalue1 = 500
	self.soundgoldvalue2 = 4000
	self.soundgoldduration = 30
	self.soundGoldF = false
	self.soundGoldFlag = true
}
func (self *Sound) soundGoldOff() {
	self.soundGoldFlag = false
}
func (self *Sound) soundGoldUpdate() {
	if self.soundGoldFlag {
		if self.soundgoldduration != 0 {
			self.soundgoldduration--
		} else {
			self.soundGoldFlag = false
		}
		if self.soundGoldF {
			self.soundGoldF = false
			self.t2val = self.soundgoldvalue1
		} else {
			self.soundGoldF = true
			self.t2val = self.soundgoldvalue2
		}
		self.soundgoldvalue1 += (self.soundgoldvalue1 >> 4)
		self.soundgoldvalue2 -= (self.soundgoldvalue2 >> 4)
	}
}
func (self *Sound) soundInt() {
	if self.soundLevDoneFlag {
		return
	}
	self.timerClock++
	if self.soundFlag && !(self.sndFlag) {
		self.sndFlag = true
	}
	if !(self.soundFlag) && self.sndFlag {
		self.sndFlag = false
		self.setSoundT2()

	}
	if !(self.musicFlag) && self.musicPlaying {
		self.musicOff()
	}
	if self.sndFlag && !(self.soundPausedFlag) {
		self.t0val = Sound_SILENCE_T0VAL
		self.t2val = Sound_DEFAULT_T2VAL
		if self.musicFlag {
			self.musicUpdate()
		}
		self.soundEmeraldUpdate()
		self.soundWobbleUpdate()
		self.soundDdieUpdate()
		self.soundBreakUpdate()
		self.soundGoldUpdate()
		self.soundEmUpdate()
		self.soundExplodeUpdate()
		self.soundFireUpdate()
		self.soundEatmUpdate()
		self.soundFallUpdate()
		self.sound1UpUpdate()
		self.soundBonusUpdate()
		if (self.t0val == Sound_SILENCE_T0VAL) || (self.t2val != Sound_DEFAULT_T2VAL) {
			self.setSoundT2()
		} else {
			self.setSoundMode()
			self.setT0()
		}
		self.setT2Val(self.t2val)
		self.pushToEngine()
	}
}
func (self *Sound) soundLevDoneOff() {
	self.soundLevDoneFlag = false
}
func (self *Sound) soundLevDoneUpdate() {
	if self.sndFlag {
		if self.nljpointer < 11 {
			self.t2val = self.newlevjingle[self.nljpointer]
		}
		self.t0val = (self.t2val + 35)
		self.musvol = 50
		self.setSoundMode()
		self.setT0()
		self.setT2Val(self.t2val)
		self.pushToEngine()
		if self.nljnoteduration > 0 {
			self.nljnoteduration--
		} else {
			self.nljnoteduration = 20
			self.nljpointer++
			if self.nljpointer > 10 {
				self.soundLevDoneOff()
			}
		}
	} else {
		self.soundLevDoneFlag = false
	}
}
func (self *Sound) soundPauseOff() {
	self.soundPausedFlag = false
}
func (self *Sound) soundStop() {
	self.soundFallOff()
	self.soundWobbleOff()
	self.soundFireOff()
	self.musicOff()
	self.soundBonusOff()
	self.soundExplodeOff()
	self.soundBreakOff()
	self.soundEmOff()
	self.soundEmeraldOff()
	self.soundGoldOff()
	self.soundEatmOff()
	self.soundDdieOff()
	self.sound1UpOff()
}
func (self *Sound) soundWobble() {
	self.soundWobbleFlag = true
}
func (self *Sound) soundWobbleOff() {
	self.soundWobbleFlag = false
	self.soundwobblen = 0
}
func (self *Sound) soundWobbleUpdate() {
	if self.soundWobbleFlag {
		self.soundwobblen++
		if self.soundwobblen > 63 {
			self.soundwobblen = 0
		}
		switch self.soundwobblen {
		case 0:
			self.t2val = 0x7d0
		case 16, 48:
			self.t2val = 0x9c4
		case 32:
			self.t2val = 0xbb8
		}
	}
}
func (self *Sound) startInt8() {
	if !(self.int8Flag) {
		self.int8Flag = true
	}
}
func (self *Sound) stopInt8() {
	if self.int8Flag {
		self.int8Flag = false
	}
	self.setT2Val(Sound_DEFAULT_T2VAL)
}
