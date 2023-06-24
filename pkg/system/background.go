package system

import (
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func newBackground(ecs *ecs.ECS) {
	entry := ecs.World.Entry(ecs.Create(
		layers.Background,
		transform.Transform,
		components.Animation,
	))

	// set background's position
	transform.SetWorldPosition(
		entry, dmath.NewVec2(consts.Width/2, consts.Height/2),
	)

	// set background's animation
	animation := components.Animation.Get(entry)
	animation.Animation = assets.GetAnimation("background")
}
