package components

import "github.com/yohamta/donburi"

type GameData struct {
	IsGameStart bool
	IsGameOver  bool
	IsGameClear bool
}

var Game = donburi.NewComponentType[GameData]()
