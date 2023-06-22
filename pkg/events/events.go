package events

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

type CollideWithBlock struct {
	ECS   *ecs.ECS
	Block *donburi.Entry
}

var CollideWithBlockEvent = events.NewEventType[CollideWithBlock]()
