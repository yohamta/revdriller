package components

import "github.com/yohamta/donburi"

type PlayerData struct {
	AnimationName string
	IsDead        bool
	State         PlayerState
}

type PlayerState string

const (
	PlayerStateIdle  PlayerState = "idle"
	PlayerStateDrill PlayerState = "drill"
)

var Player = donburi.NewComponentType[PlayerData]()

func (p *PlayerData) Animation() string {
	return p.AnimationName + "_" + string(p.State)
}
