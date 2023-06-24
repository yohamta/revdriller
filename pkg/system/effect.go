package system

import (
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/layers"
	"revdriller/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/ganim8/v2"
)

func newEffect(ecs *ecs.ECS, parent *donburi.Entry, animation string, pos dmath.Vec2) {
	entry := ecs.World.Entry(ecs.Create(
		layers.Fx,
		tags.Effect,
		transform.Transform,
		components.Animation,
	))

	// set effect animation
	a := assets.GetAnimation(animation)
	a.SetOnLoop(ganim8.PauseAtEnd)
	components.Animation.Get(entry).Animation = a

	// add effect to parent
	if parent != nil {
		transform.AppendChild(parent, entry, false)
	}

	// set effect position
	t := transform.GetTransform(entry)
	t.LocalPosition = pos
}

func updateEffects(ecs *ecs.ECS) {
	tags.Effect.Each(ecs.World, func(entry *donburi.Entry) {
		// remove effect if animation is finished
		if components.Animation.Get(entry).Animation.IsEnd() {
			entry.Remove()
		}
	})
}
