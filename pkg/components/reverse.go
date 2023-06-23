package components

import (
	"time"

	"github.com/yohamta/donburi"
)

type ReverseData struct {
	IsReversed bool
	ReversedAt time.Time
}

var Reverse = donburi.NewComponentType[ReverseData]()

func (r *ReverseData) Reverse() {
	r.IsReversed = true
	r.ReversedAt = time.Now()
}
