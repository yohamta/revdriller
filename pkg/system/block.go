package system

import (
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func newBlock(ecs *ecs.ECS, pos dmath.Vec2, speed float64) {
	entry := ecs.World.Entry(ecs.Create(
		layers.Blocks,
		transform.Transform,
		components.Block,
		components.Animation,
		components.Size,
		components.Collider,
		components.Velocity,
	))

	// set blocks data
	block := components.Block.Get(entry)
	block.MaxDurability = 10
	block.Durability = block.MaxDurability

	// set block's animation
	animation := components.Animation.Get(entry)
	animation.Animation = block.Animation()

	// set block's speed
	vel := components.Velocity.Get(entry)
	vel.Y = speed

	// set block's size
	components.Size.SetValue(entry, dmath.NewVec2(32, 32))

	// set block's position
	transform.SetWorldPosition(entry, pos)

	// set block's collider
	collider := components.Collider.Get(entry)
	collider.Hitboxes = assets.GetHitboxes("block")
}

func updateBlocks(ecs *ecs.ECS) {
	components.Block.Each(ecs.World, func(entry *donburi.Entry) {
		block := components.Block.Get(entry)

		// move block
		vel := components.Velocity.Get(entry)
		pos := transform.WorldPosition(entry)

		pos.Y += vel.Y
		transform.SetWorldPosition(entry, pos)

		if block.Durability <= 0 {
			removeBlock(ecs, entry)
		}
	})
}

// removeBlock removes block from ecs
func removeBlock(ecs *ecs.ECS, entry *donburi.Entry) {
	// create fragments
	for i := 0; i < 5; i++ {
		newFragment(ecs, transform.WorldPosition(entry))
	}
	entry.Remove()
}
