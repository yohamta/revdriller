package components

import (
	"time"

	"github.com/yohamta/donburi"
)

type StageData struct {
	Level         int
	BlockSpeed    float64
	Timer         time.Duration
	WaveCount     int
	PathColumn    int
	ShouldReverse bool
}

var Stage = donburi.NewComponentType[StageData]()
