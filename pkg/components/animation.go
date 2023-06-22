package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type AnimationData struct {
	Animation *ganim8.Animation
}

var Animation = donburi.NewComponentType[AnimationData]()
