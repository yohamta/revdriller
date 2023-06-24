package system

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

func randomRotation() float64 {
	return float64(rand.Intn(4) * 90)
}

// drawNumberL draws number from left
func drawNumberL(screen *ebiten.Image, num int, x, y float64, spr *ganim8.Sprite) {
	w := float64(spr.Width())
	for i, r := range fmt.Sprintf("%d", num) {
		drawChar(screen, r, x+float64(i)*w, y, spr)
	}
}

// drawNumberR draws number from right
func drawNumberR(screen *ebiten.Image, num int, x, y float64, spr *ganim8.Sprite) {
	s := fmt.Sprintf("%d", num)
	w := float64(spr.Width())
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		x := x - float64(len(s)-i-1)*w
		drawChar(screen, r, x, y, spr)
	}
}

func drawChar(screen *ebiten.Image, char rune, x, y float64, spr *ganim8.Sprite) {
	num := atoi(char)
	ganim8.DrawSprite(screen, spr, num,
		x, y,
		0, 1, 1, .5, .0)
}

func atoi(r rune) int {
	return int(r - '0')
}
