package events

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
	dmath "github.com/yohamta/donburi/features/math"
)

type CollideWithDrill struct {
	ECS   *ecs.ECS
	Block *donburi.Entry
}

var CollideWithDrillEvent = events.NewEventType[CollideWithDrill]()

type ReverseBlockBroken struct {
	ECS *ecs.ECS
}

var ReverseBlockBrokenEvent = events.NewEventType[ReverseBlockBroken]()

type BombBlockBroken struct {
	Point dmath.Vec2
	ECS   *ecs.ECS
}

var BombBlockBrokenEvent = events.NewEventType[BombBlockBroken]()
