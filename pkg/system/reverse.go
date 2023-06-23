package system

import (
	"revdriller/pkg/components"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
)

func newReverse(ecs *ecs.ECS) {
	ecs.Create(
		layers.Default,
		components.Reverse,
	)
}

func getReverse(ecs *ecs.ECS) *components.ReverseData {
	return components.Reverse.Get(components.Reverse.MustFirst(ecs.World))
}

func isReversed(ecs *ecs.ECS) bool {
	return getReverse(ecs).IsReversed
}
