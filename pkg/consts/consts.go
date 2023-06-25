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
	BlockSpeed = 0.65

	// BlockColumnNum is the number of columns of the block.
	BlockColumnNum = 7

	// Life is the number of life of the player.
	Life = 3

	// HighScoreNum is the number of high score.
	HighScoreNum = 3

	// Margin is the gap of the edge.
	Margin = (Width - BlockWidth*BlockColumnNum) / 2

	// MaxX is the maximum x position of the player.
	MaxX = Width - BlockWidth/2

	// MinX is the minimum x position of the player.
	MinX = BlockWidth / 2

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

	// LargeNumberSprite is the sprite of the large number.
	LargeNumberSprite = "img/numbers.png"

	// SmallNumberSprite is the sprite of the small number.
	SmallNumberSprite = "img/numbers_small.png"
)

var (
	// ColliderColor is the color of the collider.
	ColliderColor = color.RGBA{0xff, 0x00, 0x00, 0x60}
)
