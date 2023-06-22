package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

var (
	// Velocity is a component that represents the velocity of an entity.
	Velocity = donburi.NewComponentType[math.Vec2]()
	Size     = donburi.NewComponentType[math.Vec2]()
)
