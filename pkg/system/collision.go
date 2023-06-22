package system

import (
	"revdriller/pkg/collision"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/events"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

var (
	colliders = donburi.NewQuery(filter.Contains(
		components.Collider,
		transform.Transform,
		components.Size,
	))
)

// drawCollider draws collider for debug
func drawCollider(ecs *ecs.ECS, screen *ebiten.Image) {
	if consts.DebugCollision {
		colliders.Each(ecs.World, func(entry *donburi.Entry) {
			collider := components.Collider.Get(entry)
			pos := transform.WorldPosition(entry)
			size := components.Size.Get(entry)
			for _, hb := range collider.Hitboxes {
				x0, x1, y0, y1 := collision.Bounds(ecs, pos, *size, hb)
				vector.DrawFilledRect(
					screen, float32(x0), float32(y0),
					float32(x1-x0), float32(y1-y0),
					consts.ColliderColor, false,
				)
			}
		})
	}
}

var (
	blockColliders = donburi.NewQuery(filter.Contains(
		components.Collider,
		components.Block,
	))
)

func checkCollisions(ecs *ecs.ECS) {
	player, ok := components.Player.First(ecs.World)
	if !ok {
		return
	}
	collider := newCollider(player)
	blockColliders.Each(ecs.World, func(entry *donburi.Entry) {
		if collision.Collide(collider, newCollider(entry)) {
			events.CollideWithBlockEvent.Publish(ecs.World, events.CollideWithBlock{
				ECS:   ecs,
				Block: entry,
			})
		}
	})
}

func newCollider(entry *donburi.Entry) collision.Collider {
	colliderData := components.Collider.Get(entry)
	// TODO: add size data to the transform component
	pos := transform.WorldPosition(entry)
	size := *components.Size.Get(entry)
	return collision.NewCollider(pos, size, colliderData.Hitboxes)
}
