package components

import "github.com/yohamta/donburi"

type GameData struct {
	IsGameOver bool
}

var Game = donburi.NewComponentType[GameData]()
