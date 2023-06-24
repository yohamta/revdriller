package system

import (
	"math"
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/events"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func newPlayer(ecs *ecs.ECS) {
	entry := ecs.World.Entry(ecs.Create(
		layers.Player,
		transform.Transform,
		components.Player,
		components.Animation,
		components.Velocity,
		components.Size,
		components.Collider,
	))

	// set player's initial state
	player := components.Player.Get(entry)
	player.State = components.PlayerStateIdle
	player.Power = 1

	// set player's initial animation
	animation := components.Animation.Get(entry)
	animation.Animation = player.Animation()

	// set player's size
	components.Size.SetValue(entry, dmath.NewVec2(16, 24))

	// set player's position
	transform.SetWorldPosition(
		entry, dmath.NewVec2(consts.Width/2, consts.Height/2),
	)

	// set player's collider
	collider := components.Collider.Get(entry)
	collider.Hitboxes = player.Hitboxes()
}

func updatePlayer(ecs *ecs.ECS) {
	if !isGameStarted(ecs) {
		return
	}

	entry, ok := components.Player.First(ecs.World)
	if !ok {
		return
	}

	// get player's position and size
	pos := transform.WorldPosition(entry)
	size := components.Size.Get(entry)

	// apply gravity
	vel := components.Velocity.Get(entry)
	vel.Y += consts.Gravity

	// update position with velocity
	pos.Y += vel.Y

	player := components.Player.Get(entry)

	// if player is defunct, do nothing
	if !player.IsDefunct() {
		// check player's key input
		input := getInput(ecs)
		if input.ButtonA || input.Down || input.Up {
			vel.Y = -consts.PlayerJumpSpeed
			updatePlayerState(entry, components.PlayerStateDrill)
		} else {
			updatePlayerState(entry, components.PlayerStateIdle)
		}

		// adjust player's position on bottom side
		if vel.Y > 0 {
			if block, ok := findBlockOn(ecs, dmath.NewVec2(pos.X, pos.Y+size.Y/2)); ok {
				bp := transform.WorldPosition(block)
				bs := components.Size.Get(block)
				pos.Y = bp.Y - bs.Y/2 - size.Y/2
				vel.Y = 0
			}
		}

		// adjust player's position on top side
		if vel.Y < 0 {
			if block, ok := findBlockOn(ecs, pos); ok {
				bp := transform.WorldPosition(block)
				bs := components.Size.Get(block)
				pos.Y = bp.Y + bs.Y/2 + size.Y/2
			}
		}

		// move horizontally
		if input.Left {
			pos.X -= consts.PlayerJumpSpeed
			pos.X = math.Max(pos.X, size.X/2)

			// adjust player's position on left side
			if block, ok := findBlockOn(ecs, dmath.NewVec2(pos.X-size.X/2, pos.Y)); ok {
				bp := transform.WorldPosition(block)
				bs := components.Size.Get(block)
				if bp.X+bs.X/2 < consts.MaxX-size.X/2 {
					pos.X = bp.X + bs.X/2 + size.X/2
				}
			}
		}

		if input.Right {
			pos.X += consts.PlayerJumpSpeed
			pos.X = math.Min(pos.X, consts.Width-size.X/2)

			// adjust player's position on right side
			if block, ok := findBlockOn(ecs, dmath.NewVec2(pos.X+size.X/2, pos.Y)); ok {
				bp := transform.WorldPosition(block)
				bs := components.Size.Get(block)
				if bp.X-bs.X/2 > consts.MinX+size.X/2 {
					pos.X = bp.X - bs.X/2 - size.X/2
				}
			}
		}
	}

	transform.SetWorldPosition(entry, pos)

	// check if player is dead
	if pos.Y-size.Y/2-consts.DeadBuffer > consts.Height {
		components.Player.Get(entry).IsDead = true
	}
}

func updatePlayerState(entry *donburi.Entry, state components.PlayerState) {
	// update player state
	player := components.Player.Get(entry)
	if player.State == state {
		return
	}

	player.State = state

	// update player animation
	animation := components.Animation.Get(entry)
	animation.Animation = player.Animation()

	// update player collision
	collider := components.Collider.Get(entry)
	collider.Hitboxes = player.Hitboxes()

	// update player size
	components.Size.SetValue(entry, player.Size())
}

// onCollideWithBlock is called when player collide with block
func onCollideWithBlock(w donburi.World, e events.CollideWithDrill) {
	if !e.Block.Valid() {
		return
	}

	entry := components.Player.MustFirst(e.ECS.World)
	player := components.Player.Get(entry)

	// set player's velocity to zero if player is moving upward
	vel := components.Velocity.Get(entry)
	if vel.Y < 0 {
		vel.Y = 0
	}

	block := components.Block.Get(e.Block)

	// if the block is needle, player becomes defunct
	if block.Type == components.BlockTypeNeedle {
		updatePlayerState(entry, components.PlayerStateDefunct)
		return
	}

	if player.State == components.PlayerStateDrill {
		// add damage to block
		block.Damage(player.Power)

		// play sound
		assets.PlaySE(assets.SEAttack)

		// update block animation
		animation := components.Animation.Get(e.Block)
		animation.Animation = block.Animation()

		// spawn fragments
		if block.IsBreakable() && !block.IsItem() {
			size := components.Size.Get(entry)
			pos := transform.WorldPosition(entry)
			drillPos := dmath.NewVec2(pos.X, pos.Y-size.Y/2-1)
			newFragment(e.ECS, drillPos, FragmentTypeSmall)
		}
	}
}
