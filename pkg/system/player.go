package system

import (
	"math"
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
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
	player.AnimationName = "player1"
	player.State = components.PlayerStateIdle

	// set player's initial animation
	animation := components.Animation.Get(entry)
	animation.Animation = assets.GetAnimation(player.Animation())

	// set player's size
	components.Size.SetValue(entry, dmath.NewVec2(32, 32))

	// set player's position
	transform.SetWorldPosition(
		entry, dmath.NewVec2(consts.Width/2, consts.Height/2),
	)

	// set player's collider
	collider := components.Collider.Get(entry)
	collider.Hitboxes = assets.GetHitboxes("player")
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
		vel.Y = -consts.PlayerSpeed
		updatePlayerState(entry, components.PlayerStateDrill)
	} else {
		updatePlayerState(entry, components.PlayerStateIdle)
	}

	// move horizontally
	if input.Left {
		pos.X -= consts.PlayerSpeed
		pos.X = math.Max(pos.X, size.X/2)
		transform.SetWorldPosition(entry, pos)
	}
	if input.Right {
		pos.X += consts.PlayerSpeed
		pos.X = math.Min(pos.X, consts.Width-size.X/2)
		transform.SetWorldPosition(entry, pos)
	}
}

func updatePlayerState(entry *donburi.Entry, state components.PlayerState) {
	// update player state
	player := components.Player.Get(entry)
	player.State = state

	// update player animation
	animation := components.Animation.Get(entry)
	animation.Animation = assets.GetAnimation(player.Animation())
}
