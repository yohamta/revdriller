package collision

import (
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	dmath "github.com/yohamta/donburi/features/math"
)

type Hitbox struct {
	Min dmath.Vec2
	Max dmath.Vec2
}

func Bounds(ecs *ecs.ECS, pos, size math.Vec2, hb Hitbox) (x0, x1, y0, y1 float64) {
	x0 = pos.X - size.X/2 + hb.Min.X
	x1 = pos.X - size.X/2 + hb.Max.X
	y0 = pos.Y - size.Y/2 + hb.Min.Y
	y1 = pos.Y - size.Y/2 + hb.Max.Y
	return
}
