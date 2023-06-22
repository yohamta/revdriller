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

func updateAnimation(ecs *ecs.ECS) {
	components.Animation.Each(ecs.World, func(entry *donburi.Entry) {
		a := components.Animation.Get(entry)
		a.Animation.Update()
	})
}

var (
	animationQuery = donburi.NewQuery(filter.Contains(
		components.Animation,
		transform.Transform,
	))
)

func drawAnimation(ecs *ecs.ECS, screen *ebiten.Image) {
	animationQuery.Each(ecs.World, func(entry *donburi.Entry) {
		a := components.Animation.Get(entry)
		pos := transform.WorldPosition(entry)
		rot := transform.WorldRotation(entry)
		rad := math.Pi * rot / 180

		ganim8.DrawAnime(screen, a.Animation, pos.X, pos.Y, rad, 1, 1, .5, .5)
	})
}
