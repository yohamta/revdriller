package components

import (
	"revdriller/pkg/consts"

	"github.com/yohamta/donburi"
)

type GameData struct {
	Stage       int
	Score       int
	AddScore    int
	Life        int
	IsGameStart bool
	IsDead      bool
	IsClear     bool
}

var Game = donburi.NewComponentType[GameData]()

func (g *GameData) IsFirstPlay() bool {
	return g.Stage == 1 && g.Life == consts.Life
}
