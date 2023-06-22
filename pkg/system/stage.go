package system

import (
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
)

func newStage(ecs *ecs.ECS, level int) {
	entry := ecs.World.Entry(ecs.Create(
		layers.Default,
		components.Stage,
	))

	stage := components.Stage.Get(entry)
	stage.Level = level
	stage.BlockSpeed = 1 + float64(level)/10
	stage.BlockInterval = consts.BlockInterval
	stage.Waves = stage.Level * 10

	// TODO: generate only one wave
	y := float64(-consts.BlockHeight / 2)
	for i := 0; i < stage.Waves; i++ {
		generateWave(ecs, stage, y)
		y -= consts.BlockHeight
	}
}

func generateWave(ecs *ecs.ECS, stage *components.StageData, y float64) {
	x := (consts.Width-consts.CollumnCount*consts.BlockWidth)/2 - consts.BlockWidth/2
	for j := 0; j < consts.CollumnCount; j++ {
		x += consts.BlockWidth
		newBlock(ecs, dmath.NewVec2(float64(x), y), stage.BlockSpeed)
	}
}

func updateStage(ecs *ecs.ECS) {
	// TODO: generate new wave
}
