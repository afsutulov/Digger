// SPDX-License-Identifier: GPL-3.0-only
// Ported to Go on 2026-10-09 from Digger Remastered / the Java port.
package game

type GameState struct {
	lives     int
	level     int
	dead      bool
	levelDone bool
}

func newGameState() *GameState {
	self := &GameState{}
	return self
}
