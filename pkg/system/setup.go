package system

import (
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

func Setup(ecs *ecs.ECS) {
	// Add systems.
	ecs.AddSystem(updateInput)
	ecs.AddSystem(updateAnimation)
	ecs.AddSystem(updateFragments)
	ecs.AddSystem(updatePlayer)
	ecs.AddSystem(updateGame)
	ecs.AddSystem(checkCollisions)
	ecs.AddSystem(updateBlocks)
	ecs.AddSystem(processEvents)

	// Add renderers.
	ecs.AddRenderer(layers.Default, drawGame)
	ecs.AddRenderer(layers.Blocks, drawAnimation)
	ecs.AddRenderer(layers.Player, drawAnimation)
	ecs.AddRenderer(layers.Fx, drawAnimation)
	ecs.AddRenderer(layers.System, drawCollider)

	// Add entities.
	newGame(ecs)
	newInput(ecs)
	newPlayer(ecs)
}

func processEvents(ecs *ecs.ECS) {
	events.ProcessAllEvents(ecs.World)
}
