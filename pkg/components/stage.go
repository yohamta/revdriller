package components

import "github.com/yohamta/donburi"

type StageData struct {
	Level      int
	BlockSpeed float64
}

var Stage = donburi.NewComponentType[StageData]()
