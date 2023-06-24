package components

import (
	"time"

	"github.com/yohamta/donburi"
)

type StageData struct {
	Level         int
	BlockSpeed    float64
	Timer         time.Duration
	BaseLine      float64
	WaveCount     int
	PathColumn    int
	ShouldReverse bool
	Reversed      bool
}

var Stage = donburi.NewComponentType[StageData]()
