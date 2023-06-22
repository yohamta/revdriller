package system

import (
	"revdriller/pkg/components"

	"github.com/yohamta/donburi/ecs"
)

func UpdateGame(ecs *ecs.ECS) {
	game := components.Game.MustFirst(ecs.World)

	player, ok := components.Player.First(ecs.World)
	if ok {
		if components.Player.Get(player).IsDead {
			components.Game.Get(game).IsGameOver = true
		}
	}
}
