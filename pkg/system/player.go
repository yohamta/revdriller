package system

import (
	"revdriller/pkg/components"
	"revdriller/pkg/consts"

	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
)

func UpdatePlayer(ecs *ecs.ECS) {
	player, ok := components.Player.First(ecs.World)
	if !ok {
		return
	}

	// apply gravity
	vel := components.Velocity.Get(player)
	vel.Y += consts.Gravity

	// update position
	pos := transform.WorldPosition(player)
	pos.Y += vel.Y
	transform.SetWorldPosition(player, pos)

	// check die
	if pos.Y > consts.Height {
		components.Player.Get(player).IsDead = true
	}
}
