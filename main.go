package main

import (
	"log"
	"math/rand"
	"revdriller/game"
	"revdriller/pkg/consts"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() { rand.Seed(time.Now().UnixNano()) }

func main() {
	ebiten.SetCursorShape(ebiten.CursorShapePointer)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("Revdriller")

	factor := ebiten.DeviceScaleFactor()

	ebiten.SetWindowSize(int(consts.Width*factor), int(consts.Height*factor))
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}
