package system

import (
	"revdriller/pkg/components"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"

	pkgevents "revdriller/pkg/events"
)

func Setup(ecs *ecs.ECS, game components.GameData) {
	// Add systems.
	ecs.AddSystem(updateInput)
	ecs.AddSystem(updateAnimation)
	ecs.AddSystem(updateFragments)
	ecs.AddSystem(updatePlayer)
	ecs.AddSystem(updateBlocks)
	ecs.AddSystem(updateEffects)
	ecs.AddSystem(updateGame)
	ecs.AddSystem(checkDrillCollision)
	ecs.AddSystem(processEvents)

	// Add renderers.
	ecs.AddRenderer(layers.Background, drawAnimation(layers.Background))
	ecs.AddRenderer(layers.Blocks, drawAnimation(layers.Blocks))
	ecs.AddRenderer(layers.Player, drawAnimation(layers.Player))
	ecs.AddRenderer(layers.Fx, drawAnimation(layers.Fx))
	ecs.AddRenderer(layers.System, drawAnimation(layers.System))
	ecs.AddRenderer(layers.System, drawCollider)
	ecs.AddRenderer(layers.System, drawGame)

	// Add entities.
	newGame(ecs, game)
	newBackground(ecs)
	newStage(ecs, game.Stage)
	newReverse(ecs)
	newInput(ecs)
	newPlayer(ecs)

	// Subscribe events.
	pkgevents.CollideWithDrillEvent.Subscribe(ecs.World, onCollideWithBlock)
	pkgevents.ReverseBlockBrokenEvent.Subscribe(ecs.World, reverseBlocks)
	pkgevents.BombBlockBrokenEvent.Subscribe(ecs.World, bombBlock)
}

func processEvents(ecs *ecs.ECS) {
	events.ProcessAllEvents(ecs.World)
}
