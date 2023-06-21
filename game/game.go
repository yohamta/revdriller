package game

import "github.com/hajimehoshi/ebiten/v2"

type game struct {
}

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
	return nil
}
