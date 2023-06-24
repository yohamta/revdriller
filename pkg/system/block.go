package system

import (
	"math"
	"revdriller/pkg/collision"
	"revdriller/pkg/components"
	"revdriller/pkg/events"
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
	block.Init(blockType)

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

func reverseBlocks(w donburi.World, e events.ReverseBlockBroken) {
	components.Block.Each(w, func(entry *donburi.Entry) {
		block := components.Block.Get(entry)

		// change block type
		block.Init(block.Type.Reverse())

		// set block's animation
		animation := components.Animation.Get(entry)
		animation.Animation = block.Animation()
	})
}

func updateBlocks(ecs *ecs.ECS) {
	components.Block.Each(ecs.World, func(entry *donburi.Entry) {
		block := components.Block.Get(entry)

		// move block
		vel := components.Velocity.Get(entry)
		pos := transform.WorldPosition(entry)

		pos.Y += vel.Y
		transform.SetWorldPosition(entry, pos)

		if block.IsBroken() {
			removeBlock(ecs, entry)
		}
	})
}

func findBlockOn(ecs *ecs.ECS, point dmath.Vec2) (*donburi.Entry, bool) {
	// TODO: make it more efficient
	var found *donburi.Entry
	var distance float64 = math.MaxFloat64
	components.Block.Each(ecs.World, func(entry *donburi.Entry) {
		if collision.Contain(newCollider(entry), point) {
			pos := transform.WorldPosition(entry)
			if tmp := pos.Distance(point); tmp < distance {
				distance = tmp
				found = entry
			}
		}
	})
	return found, found != nil
}

// removeBlock removes block from ecs
func removeBlock(ecs *ecs.ECS, entry *donburi.Entry) {
	block := components.Block.Get(entry)
	// create fragments
	if block.IsBroken() {
		for i := 0; i < 5; i++ {
			newFragment(ecs, transform.WorldPosition(entry))
		}

		// publish reverse block broken event
		if block.Type == components.BlockTypeReverse {
			events.ReverseBlockBrokenEvent.Publish(ecs.World, events.ReverseBlockBroken{ECS: ecs})
		}
	}

	entry.Remove()
}
