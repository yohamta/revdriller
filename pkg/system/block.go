package system

import (
	"math"
	"revdriller/assets"
	"revdriller/pkg/collision"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
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

		// generate effect
		if block.IsReversible() {
			if block.Width() == consts.BlockWidth*2 {
				newEffect(e.ECS, entry, "reverse_effect", dmath.NewVec2(-consts.BlockWidth/2, 0))
				newEffect(e.ECS, entry, "reverse_effect", dmath.NewVec2(+consts.BlockWidth/2, 0))
			} else {
				newEffect(e.ECS, entry, "reverse_effect", dmath.NewVec2(0, 0))
			}
		}
	})
}

func updateBlocks(ecs *ecs.ECS) {
	if !isGameStarted(ecs) {
		return
	}
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

		// remove block if it is out of screen
		if pos.Y-block.Height()/2 > consts.Height {
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

func bombBlock(w donburi.World, e events.BombBlockBroken) {
	// remove blocks around the point
	for i := 0; i < 5; i++ {
		x := e.Point.X - 2*consts.BlockWidth + float64(i)*consts.BlockWidth
		for j := 0; j < 5; j++ {
			y := e.Point.Y - 2*consts.BlockHeight + float64(j)*consts.BlockHeight
			components.Block.Each(w, func(entry *donburi.Entry) {
				if collision.Contain(newCollider(entry), dmath.NewVec2(x, y)) {
					block := components.Block.Get(entry)

					// remove block
					block.ForceBreak()

					// generate explosion effect
					pos := transform.WorldPosition(entry)
					newEffect(e.ECS, nil, "explosion_effect", pos)
				}
			})
		}
	}
}

// removeBlock removes block from ecs
func removeBlock(ecs *ecs.ECS, entry *donburi.Entry) {
	block := components.Block.Get(entry)
	// create fragments
	if block.IsBroken() {
		for i := 0; i < 10; i++ {
			newFragment(ecs, transform.WorldPosition(entry), FragmentTypeLarge)
		}

		// play sound
		assets.PlaySE(assets.SEBreak)

		// publish reverse block broken event
		switch block.Type {
		case components.BlockTypeReverse:
			if !block.Invalidated {
				events.ReverseBlockBrokenEvent.Publish(ecs.World, events.ReverseBlockBroken{ECS: ecs})
			}
		case components.BlockTypeBomb:
			events.BombBlockBrokenEvent.Publish(ecs.World, events.BombBlockBroken{
				ECS: ecs, Point: transform.WorldPosition(entry),
			})
		}

		// add score
		game := getGame(ecs)
		game.AddScore += block.Score
	}

	transform.RemoveRecursive(entry)
}
