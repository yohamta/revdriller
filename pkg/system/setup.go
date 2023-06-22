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
	ecs.AddSystem(updateBlocks)
	ecs.AddSystem(updateStage)
	ecs.AddSystem(updateGame)
	ecs.AddSystem(checkDrillCollision)
	ecs.AddSystem(processEvents)

	// Add renderers.
	ecs.AddRenderer(layers.Blocks, drawAnimation(layers.Blocks))
	ecs.AddRenderer(layers.Player, drawAnimation(layers.Player))
	ecs.AddRenderer(layers.Fx, drawAnimation(layers.Fx))
	ecs.AddRenderer(layers.System, drawAnimation(layers.System))
	ecs.AddRenderer(layers.System, drawCollider)
	ecs.AddRenderer(layers.System, drawGame)

	// Add entities.
	newGame(ecs)
	newInput(ecs)
	newPlayer(ecs)
}

func processEvents(ecs *ecs.ECS) {
	events.ProcessAllEvents(ecs.World)
}
