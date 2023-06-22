package assets

import (
	"revdriller/pkg/collision"

	dmath "github.com/yohamta/donburi/features/math"
)

type hitboxConfig struct {
	Hitboxes map[string][][]int `json:"hitboxes"`
}

var hitboxes = make(map[string][]collision.Hitbox)

func loadHitboxes(cfg *hitboxConfig) {
	for name, hs := range cfg.Hitboxes {
		hitboxes[name] = make([]collision.Hitbox, len(hs))
		for i, h := range hs {
			hitboxes[name][i] = collision.Hitbox{
				Min: dmath.Vec2{X: float64(h[0]), Y: float64(h[1])},
				Max: dmath.Vec2{X: float64(h[2]), Y: float64(h[3])},
			}
		}
	}
}

func GetHitboxes(name string) []collision.Hitbox {
	if _, ok := hitboxes[name]; !ok {
		panic("hitboxes not found: " + name)
	}
	return hitboxes[name]
}
