package collision

import (
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	dmath "github.com/yohamta/donburi/features/math"
)

// Hitbox is a rectangle that represents a collision area.
// The origin is the center of the entity.
type Hitbox struct {
	Min dmath.Vec2
	Max dmath.Vec2
}

// Bounds returns the bounds of the hitbox.
func Bounds(ecs *ecs.ECS, pos, size math.Vec2, hb Hitbox) (x0, x1, y0, y1 float64) {
	x0 = pos.X - size.X/2 + hb.Min.X
	x1 = pos.X - size.X/2 + hb.Max.X
	y0 = pos.Y - size.Y/2 + hb.Min.Y
	y1 = pos.Y - size.Y/2 + hb.Max.Y
	return
}

// Collider is a struct that represents a collision area.
type Collider struct {
	Pos      dmath.Vec2
	Size     dmath.Vec2
	Hitboxes []Hitbox
}

// NewCollider returns a new collider.
func NewCollider(pos, size dmath.Vec2, hitboxes []Hitbox) Collider {
	return Collider{pos, size, hitboxes}
}

// Collide returns true if two colliders collide.
func Collide(c1, c2 Collider) bool {
	for _, hb1 := range c1.Hitboxes {
		for _, hb2 := range c2.Hitboxes {
			if collide(hb1, c1.Pos, c1.Size, hb2, c2.Pos, c2.Size) {
				return true
			}
		}
	}
	return false
}

// collide returns true if two hitboxes collide.
func collide(hb1 Hitbox, pos1, size1 dmath.Vec2, hb2 Hitbox, pos2, size2 dmath.Vec2) bool {
	x0, x1, y0, y1 := Bounds(nil, pos1, size1, hb1)
	x2, x3, y2, y3 := Bounds(nil, pos2, size2, hb2)
	return x0 < x3 && x1 > x2 && y0 < y3 && y1 > y2
}
