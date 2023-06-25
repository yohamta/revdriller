package components

import (
	"revdriller/pkg/consts"

	"github.com/yohamta/donburi"
)

type GameData struct {
	Stage      int
	Score      int
	AddScore   int
	Life       int
	State      GameState
	HighScores []int
}

type GameState int

const (
	GameStateTitle GameState = iota
	GameStatePlay
	GameStateClear
	GameStateGameOver
)

var Game = donburi.NewComponentType[GameData]()

func (g *GameData) IsFirstPlay() bool {
	return g.Stage == 1 && g.Life == consts.Life
}

func (g *GameData) TotalScore() int {
	return g.Score + g.AddScore
}
