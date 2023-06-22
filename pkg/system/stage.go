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

	y := 0.
	x := (consts.Width-consts.CollumnCount*consts.BlockWidth)/2 - consts.BlockWidth/2
	for i := 0; i < consts.CollumnCount; i++ {
		x += consts.BlockWidth
		newBlock(ecs, dmath.NewVec2(float64(x), y), stage.BlockSpeed)
	}
}
