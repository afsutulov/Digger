// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

type MonsterState struct {
	x             int
	y             int
	h             int
	v             int
	xr            int
	yr            int
	direction     int
	horizontalDir int
	kind          int
	huntTime      int
	death         int
	bag           int
	deathTime     int
	spawnTime     int
	flag          bool
	nob           bool
	alive         bool
}

func newMonsterState() *MonsterState {
	self := &MonsterState{}
	return self
}
