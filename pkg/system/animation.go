package system

import (
	"math"
	"revdriller/pkg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
)

// updateAnimation updates animation
func updateAnimation(ecs *ecs.ECS) {
	components.Animation.Each(ecs.World, func(entry *donburi.Entry) {
		a := components.Animation.Get(entry)
		a.Animation.Update()
	})
}

// drawAnimation draws animation on specified layer
func drawAnimation(layer ecs.LayerID) func(ecs *ecs.ECS, screen *ebiten.Image) {
	animations := ecs.NewQuery(layer, filter.Contains(
		components.Animation,
		transform.Transform,
	))
	return func(ecs *ecs.ECS, screen *ebiten.Image) {
		animations.Each(ecs.World, func(entry *donburi.Entry) {
			a := components.Animation.Get(entry)
			pos := transform.WorldPosition(entry)
			rot := transform.WorldRotation(entry)
			rad := math.Pi * rot / 180

			ganim8.DrawAnime(screen, a.Animation, pos.X, pos.Y, rad, 1, 1, .5, .5)
		})
	}
}
