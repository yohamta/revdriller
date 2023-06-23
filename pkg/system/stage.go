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
	stage.BaseLine = 0.0
}

func getStage(ecs *ecs.ECS) *components.StageData {
	return components.Stage.Get(components.Stage.MustFirst(ecs.World))
}

func generateWave(ecs *ecs.ECS, stage *components.StageData, y float64) {
	x := -consts.BlockWidth / 2
	for x < consts.Width+consts.BlockWidth/2 {
		block := newBlock(ecs, dmath.NewVec2(float64(x), y),
			components.RandomBlockType(), stage.BlockSpeed)
		size := components.Size.Get(block)
		x += int(size.X)
	}
}

func updateStage(ecs *ecs.ECS) {
	if !isGameStarted(ecs) {
		return
	}
	stage := getStage(ecs)
	stage.BaseLine += stage.BlockSpeed
	if stage.BaseLine >= 0 {
		stage.BaseLine -= consts.BlockHeight
		generateWave(ecs, stage, stage.BaseLine)
	}
}
