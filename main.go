package main

import (
	"log"
	"math/rand"
	"revdriller/game"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() { rand.Seed(time.Now().UnixNano()) }

const (
	widthInPixel  = 320
	heightInPixel = 480
)

func main() {
	ebiten.SetCursorShape(ebiten.CursorShapePointer)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("Revdriller")

	factor := ebiten.DeviceScaleFactor()

	ebiten.SetWindowSize(int(widthInPixel*factor), int(heightInPixel*factor))
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}

func loadAssets() error {

	return nil
}
