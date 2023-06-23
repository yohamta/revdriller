package system

import (
	"revdriller/pkg/collision"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func newBlock(ecs *ecs.ECS, leftBottom dmath.Vec2, blockType components.BlockType, speed float64) *donburi.Entry {
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
	block.BlockType = blockType
	block.MaxDurability = 10
	block.Durability = block.MaxDurability

	// set block's animation
	animation := components.Animation.Get(entry)
	animation.Animation = block.Animation()

	// set block's speed
	vel := components.Velocity.Get(entry)
	vel.Y = speed

	// set block's size
	width := block.Width()
	height := block.Height()
	components.Size.SetValue(entry, dmath.NewVec2(width, height))

	// set block's position
	pos := dmath.NewVec2(leftBottom.X+width/2, leftBottom.Y-height/2)
	transform.SetWorldPosition(entry, pos)

	// set block's collider
	collider := components.Collider.Get(entry)
	collider.Hitboxes = block.Hitboxes()

	return entry
}

func updateBlocks(ecs *ecs.ECS) {
	components.Block.Each(ecs.World, func(entry *donburi.Entry) {
		block := components.Block.Get(entry)

		// move block
		vel := components.Velocity.Get(entry)
		pos := transform.WorldPosition(entry)

		pos.Y += vel.Y
		transform.SetWorldPosition(entry, pos)

		size := *components.Size.Get(entry)

		if block.Durability <= 0 {
			removeBlock(ecs, entry)
		} else if pos.Y-size.Y/2 > consts.Height {
			removeBlock(ecs, entry)
		}
	})

	// adjust player's position after block movement
	if player, ok := components.Player.First(ecs.World); ok {
		pc := newCollider(player)

		components.Block.Each(ecs.World, func(entry *donburi.Entry) {
			bc := newCollider(entry)
			if collision.Collide(bc, pc) {
				pp := transform.WorldPosition(player)
				ps := *components.Size.Get(player)
				bp := transform.WorldPosition(entry)
				bs := *components.Size.Get(entry)
				pp.Y = bp.Y + bs.Y/2 + ps.Y/2
				transform.SetWorldPosition(player, pp)
			}
		})
	}
}

func findBlockOn(ecs *ecs.ECS, point dmath.Vec2) (*donburi.Entry, bool) {
	// TODO: make it more efficient
	var found *donburi.Entry
	components.Block.Each(ecs.World, func(entry *donburi.Entry) {
		if collision.Contain(newCollider(entry), point) {
			found = entry
		}
	})
	return found, found != nil
}

// removeBlock removes block from ecs
func removeBlock(ecs *ecs.ECS, entry *donburi.Entry) {
	block := components.Block.Get(entry)
	// create fragments
	if block.Durability <= 0 {
		for i := 0; i < 5; i++ {
			newFragment(ecs, transform.WorldPosition(entry))
		}
	}
	entry.Remove()
}
