package system

import (
	"math"
	"math/rand"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/events"
	"revdriller/pkg/layers"

	"github.com/yohamta/donburi"
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
	stage.BlockSpeed = consts.BlockSpeed + float64(level)/10
	stage.BaseLine = 0.0
	stage.WaveCount = 30 + level*5
	stage.ColumnCount = consts.Width / consts.BlockWidth
	stage.PathColumn = rand.Intn(stage.ColumnCount) + 1
}

func getStage(ecs *ecs.ECS) *components.StageData {
	return components.Stage.Get(components.Stage.MustFirst(ecs.World))
}

func generateWave(ecs *ecs.ECS, stage *components.StageData, y float64) {
	path := stage.PathColumn
	nextPath := path + rand.Intn(3) - 1
	nextPath = int(math.Min(math.Max(float64(nextPath), 1), float64(stage.ColumnCount)))
	shouldReverse := stage.ShouldReverse
	nextShouldReverse := shouldReverse
	x := (consts.Width - stage.ColumnCount*consts.BlockWidth) / 2

	// TODO: fix this ugly code
	for i := 0; i < stage.ColumnCount; {
		var blockType components.BlockType

		switch {
		case i == nextPath-1:
			blockType = components.BlockTypeNormal2
			if shouldReverse {
				blockType = blockType.Reverse()
			}
		case i == path-1 && i == nextPath:
			blockType = components.BlockTypeNormal2
			if shouldReverse {
				blockType = blockType.Reverse()
			}
		default:
			blockType = components.RandomBlockType()
		}

		if stage.Reversed {
			blockType = blockType.Reverse()
		}

		if i == path && rand.Float64() < .2 {
			blockType = components.BlockTypeReverse
			nextShouldReverse = !shouldReverse
		}

		if math.Abs(float64(i)-float64(nextPath)) < 2 && rand.Float64() < .3 {
			blockType = components.BlockTypeBomb
		}

		block := newBlock(ecs, dmath.NewVec2(float64(x), y),
			blockType, stage.BlockSpeed)
		size := components.Size.Get(block)
		x += int(size.X)
		i += int(size.X) / consts.BlockWidth
	}

	stage.PathColumn = nextPath
	stage.ShouldReverse = nextShouldReverse
}

func updateStage(ecs *ecs.ECS) {
	if !isGameStarted(ecs) {
		return
	}
	stage := getStage(ecs)
	stage.BaseLine += stage.BlockSpeed
	if stage.BaseLine >= 0 {
		stage.BaseLine -= consts.BlockHeight
		if stage.WaveCount > 0 {
			stage.WaveCount--
			generateWave(ecs, stage, stage.BaseLine)
		}
	}
}

func onReversed(w donburi.World, e events.ReverseBlockBroken) {
	stage := getStage(e.ECS)
	stage.Reversed = !stage.Reversed
}
