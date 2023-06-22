package system

import (
	"math"
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
	components.Size.SetValue(entry, dmath.NewVec2(32, 32))

	// set player's position
	transform.SetWorldPosition(
		entry, dmath.NewVec2(consts.Width/2, consts.Height/2),
	)

	// set player's collider
	collider := components.Collider.Get(entry)
	collider.Hitboxes = player.Hitboxes()

	// setup event
	events.CollideWithBlockEvent.Subscribe(ecs.World, onCollideWithBlock)
}

func updatePlayer(ecs *ecs.ECS) {
	if !isGameStart(ecs) {
		return
	}

	entry, ok := components.Player.First(ecs.World)
	if !ok {
		return
	}

	// apply gravity
	vel := components.Velocity.Get(entry)
	vel.Y += consts.Gravity

	// update position
	pos := transform.WorldPosition(entry)
	pos.Y += vel.Y
	transform.SetWorldPosition(entry, pos)

	// get size
	size := components.Size.Get(entry)

	// check if player is dead
	if pos.Y-size.Y/2-consts.DeadBuffer > consts.Height {
		components.Player.Get(entry).IsDead = true
	}

	// check player's key input
	input := getInput(ecs)
	if input.ButtonA || input.Down {
		vel.Y = -consts.PlayerJumpSpeed
		updatePlayerState(entry, components.PlayerStateDrill)
	} else {
		updatePlayerState(entry, components.PlayerStateIdle)
	}

	// move horizontally
	if input.Left {
		pos.X -= consts.PlayerJumpSpeed
		pos.X = math.Max(pos.X, size.X/2)
		transform.SetWorldPosition(entry, pos)
	}
	if input.Right {
		pos.X += consts.PlayerJumpSpeed
		pos.X = math.Min(pos.X, consts.Width-size.X/2)
		transform.SetWorldPosition(entry, pos)
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
func onCollideWithBlock(w donburi.World, e events.CollideWithBlock) {
	entry := components.Player.MustFirst(e.ECS.World)
	adjustPlayerPosition(e.ECS, entry, e.Block)

	// set player's velocity to zero
	vel := components.Velocity.Get(entry)
	vel.Y = 0

	player := components.Player.Get(entry)

	// add damage to block
	if player.State == components.PlayerStateDrill {
		block := components.Block.Get(e.Block)
		block.Damage(player.Power)

		// update block animation
		animation := components.Animation.Get(e.Block)
		animation.Animation = block.Animation()
	}
}

func adjustPlayerPosition(ecs *ecs.ECS, playerEntry, blockEntry *donburi.Entry) {
	bp := transform.WorldPosition(blockEntry)
	bs := components.Size.Get(blockEntry)

	bt := bp.Y - bs.Y/2
	bl := bp.X - bs.X/2
	br := bp.X + bs.X/2
	bb := bp.Y + bs.Y/2

	pp := transform.WorldPosition(playerEntry)
	ps := components.Size.Get(playerEntry)

	pt := pp.Y - ps.Y/2
	pl := pp.X - ps.X/2
	pr := pp.X + ps.X/2
	pb := pp.Y + ps.Y/2

	// TODO: maybe this can be simplified
	if pt < bb && bb < pb && ((pl < bl && bl < pr) || (pl < br && br < pr)) {
		pp.Y = bp.Y + ps.Y/2 + bs.Y/2
	} else if pb > bt && bt > pt && ((pl < bl && bl < pr) || (pl < br && br < pr)) {
		pp.Y = bp.Y - ps.Y/2 - bs.Y/2
	} else if pl < br && br < pr && ((pt < bt && bt < pb) || (pt < bb && bb < pb)) {
		pp.X = bp.X + ps.X/2 + bs.X/2
	} else if pr > bl && bl > pl && ((pt < bt && bt < pb) || (pt < bb && bb < pb)) {
		pp.X = bp.X - ps.X/2 - bs.X/2
	}

	transform.SetWorldPosition(playerEntry, pp)
}
