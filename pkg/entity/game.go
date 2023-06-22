package entity

import (
	"revdriller/pkg/components"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
)

func NewGame(ecs *ecs.ECS) {
	_ = ecs.World.Entry(ecs.Create(
		layers.Default,
		components.Game,
	))
}
