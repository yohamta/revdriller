package components

import (
	"revdriller/pkg/collision"

	"github.com/yohamta/donburi"
)

type ColliderData struct {
	Hitboxes []collision.Hitbox
}

var Collider = donburi.NewComponentType[ColliderData]()
