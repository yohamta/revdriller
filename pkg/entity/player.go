package entity

import (
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewPlayer(ecs *ecs.ECS) {
	entry := ecs.World.Entry(ecs.Create(
		layers.Player,
		transform.Transform,
		components.Player,
		components.Animation,
		components.Velocity,
		components.Size,
	))

	player := components.Player.Get(entry)
	player.AnimationName = "player1"

	animation := components.Animation.Get(entry)
	animation.Animation = assets.GetAnimation("player1" + "_idle")

	components.Size.SetValue(entry, math.NewVec2(32, 32))

	transform.SetWorldPosition(
		entry, math.NewVec2(consts.Width/2, consts.Height/2),
	)
}
