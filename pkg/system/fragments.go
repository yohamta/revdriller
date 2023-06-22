package system

import (
	"fmt"
	"math/rand"
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/layers"
	"revdriller/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func newFragment(ecs *ecs.ECS, pos dmath.Vec2) {
	entry := ecs.World.Entry(ecs.Create(
		layers.Fx,
		transform.Transform,
		components.Animation,
		components.Velocity,
		tags.Fragment,
	))

	// set fragment's animation
	animation := components.Animation.Get(entry)
	animation.Animation = assets.GetAnimation(
		fmt.Sprintf("fragment_%d", rand.Intn(4)+1),
	)

	// set fragment's position
	transform.SetWorldPosition(entry, pos)

	// set random velocity
	vel := components.Velocity.Get(entry)
	vel.X = rand.Float64()*3 - 1
	vel.Y = rand.Float64()*2 - 1

	// set random rotation
	transform.SetWorldRotation(entry, randomRotation())
}

func updateFragments(ecs *ecs.ECS) {
	tags.Fragment.Each(ecs.World, func(entry *donburi.Entry) {
		vel := components.Velocity.Get(entry)
		pos := transform.WorldPosition(entry)
		pos.X += vel.X
		pos.Y += vel.Y
		transform.SetWorldPosition(entry, pos)

		// apply gravity
		vel.Y += consts.Gravity
	})
}
