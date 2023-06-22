package consts

import "time"

const (
	// Width and Height are the width and height of the screen.
	Width  = 240
	Height = 320

	// Gravity is the gravity of the game.
	Gravity = 0.05

	// DeadBuffer is the buffer for the player to be considered dead.
	DeadBuffer = 60

	// PlayerSpeed is the speed of the player.
	PlayerSpeed = 2

	// PlayerHorizontalSpeed is the horizontal speed of the player.
	PlayerHorizontalSpeed = .5

	// StateDuration is the minimum duration of the state.
	StateDuration = time.Second * 1
)
