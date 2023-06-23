package components

import (
	"time"

	"github.com/yohamta/donburi"
)

type StageData struct {
	Level         int
	BlockSpeed    float64
	BlockInterval time.Duration
	Timer         time.Duration
	BaseLine      float64
}

var Stage = donburi.NewComponentType[StageData]()
