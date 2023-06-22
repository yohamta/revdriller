package components

import "github.com/yohamta/donburi"

type PlayerData struct {
	AnimationName string
	IsDead        bool
}

var Player = donburi.NewComponentType[PlayerData]()
