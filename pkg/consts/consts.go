package consts

import (
	"image/color"
	"time"
)

const (
	// Width and Height are the width and height of the screen.
	Width  = 240
	Height = 480

	// BlockWidth is the width of the block.
	BlockWidth = 32

	// BlockHeight is the height of the block.
	BlockHeight = 32

	// BlockSpeed is the speed of the block.
	BlockSpeed = 0.7

	// Gravity is the gravity of the game.
	Gravity = 0.05

	// DeadBuffer is the buffer for the player to be considered dead.
	DeadBuffer = 60

	// PlayerJumpSpeed is the jump speed of the player.
	PlayerJumpSpeed = 2

	// PlayerSpeed is the horizontal speed of the player.
	PlayerSpeed = .5

	// StateDuration is the minimum duration of the state.
	StateDuration = time.Millisecond * 300

	// DebugCollision is the flag to enable collision debug.
	DebugCollision = false
)

var (
	// ColliderColor is the color of the collider.
	ColliderColor = color.RGBA{0xff, 0x00, 0x00, 0x60}
)
