package components

import (
	"revdriller/assets"
	"revdriller/pkg/collision"

	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/ganim8/v2"
)

type PlayerData struct {
	IsDead bool
	State  PlayerState
	Power  int
}

type PlayerState string

const (
	PlayerStateIdle  PlayerState = "idle"
	PlayerStateDrill PlayerState = "drill"
)

var Player = donburi.NewComponentType[PlayerData]()

func (p *PlayerData) Animation() *ganim8.Animation {
	return assets.GetAnimation("player1_" + string(p.State))
}

func (p *PlayerData) Hitboxes() []collision.Hitbox {
	return assets.GetHitboxes("player_" + string(p.State))
}

func (p *PlayerData) Size() dmath.Vec2 {
	if p.State == PlayerStateDrill {
		return dmath.NewVec2(32, 60)
	}
	return dmath.NewVec2(32, 24)
}
