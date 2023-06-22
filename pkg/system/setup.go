package system

import (
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
)

func Setup(ecs *ecs.ECS) {
	// Add systems.
	ecs.AddSystem(updateInput)
	ecs.AddSystem(updateAnimation)
	ecs.AddSystem(updatePlayer)
	ecs.AddSystem(updateGame)

	// Add renderers.
	ecs.AddRenderer(layers.Default, drawGame)
	ecs.AddRenderer(layers.Player, drawAnimation)
	ecs.AddRenderer(layers.System, drawCollider)

	// Add entities.
	newGame(ecs)
	newInput(ecs)
	newPlayer(ecs)
}
