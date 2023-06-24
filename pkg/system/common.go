package system

import (
	"fmt"
	"math/rand"
	"revdriller/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

func randomRotation() float64 {
	return float64(rand.Intn(4) * 90)
}

func drawNumber(screen *ebiten.Image, num int, x, y float64) {
	s := fmt.Sprintf("%d", num)
	w := float64(assets.GetSprite("img/numbers.png").Width())
	for i, r := range s {
		drawChar(screen, r, x+float64(i)*w, y)
	}
}

func drawChar(screen *ebiten.Image, char rune, x, y float64) {
	spr := assets.GetSprite("img/numbers.png")
	num := atoi(char)
	ganim8.DrawSprite(screen, spr, num,
		x, y,
		0, 1, 1, .5, .0)
}

func atoi(r rune) int {
	return int(r - '0')
}
