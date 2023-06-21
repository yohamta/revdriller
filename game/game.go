package game

import (
	"log"
	"revdriller/sprites"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
}

var (
	setupOnce sync.Once
)

var _ ebiten.Game = &game{}

func New() *game {
	return &game{}
}

// Draw implements ebiten.game.
func (*game) Draw(screen *ebiten.Image) {
}

// Layout implements ebiten.game.
func (*game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	factor := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * factor), int(float64(outsideHeight) * factor)
}

// Update implements ebiten.game.
func (*game) Update() error {
	setupOnce.Do(setup)

	return nil
}

// setup loads assets and initializes the game.
func setup() {
	loadAssets()
}

// loadAssets loads all assets.
func loadAssets() {
	for _, fn := range []func() error{
		sprites.Load,
	} {
		if err := fn(); err != nil {
			log.Fatal(err)
		}
	}
}
