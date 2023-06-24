package system

import (
	"math/rand"
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
	stage.BlockSpeed = consts.BlockSpeed + float64(level)*0.02
	stage.WaveCount = 10 + (level-1)*3
	stage.PathColumn = rand.Intn(consts.BlockColumnNum)

	for i := 0; i < stage.WaveCount; i++ {
		generateWave(ecs, stage, -float64(i)*consts.BlockHeight)
	}
}

func getStage(ecs *ecs.ECS) *components.StageData {
	return components.Stage.Get(components.Stage.MustFirst(ecs.World))
}

func generateWave(ecs *ecs.ECS, stage *components.StageData, y float64) {
	// TODO: there seems to be a bug that there's no path to the goal
	path := stage.PathColumn
	switch rand.Intn(4) {
	case 0:
		for i := 0; i < consts.BlockColumnNum; i++ {
			if i == path {
				genBlock(ecs, i, y, components.BlockTypeNormal)
			} else if i == path-1 {
				genBlock(ecs, i, y, components.RandomBlockType().Shorten())
			} else {
				if genBlock(ecs, i, y, components.RandomBlockType()) {
					i++
				}
			}
		}
	case 1:
		if path < consts.BlockColumnNum-1 {
			nextPath := path + 1
			for i := 0; i < consts.BlockColumnNum; i++ {
				if i == path && i == nextPath-1 {
					genBlock(ecs, i, y, components.BlockTypeNormal2)
					i++
				} else {
					if genBlock(ecs, i, y, components.RandomBlockType()) {
						i++
					}
				}
			}
			stage.PathColumn = nextPath
		} else {
			for i := 0; i < consts.BlockColumnNum; i++ {
				if i == path {
					genBlock(ecs, i, y, components.BlockTypeNormal)
				} else {
					genBlock(ecs, i, y, components.RandomBlockType().Shorten())
				}
			}
		}
	case 2:
		if path > 0 {
			nextPath := path - 1
			for i := 0; i < consts.BlockColumnNum; i++ {
				if i == path-1 {
					genBlock(ecs, i, y, components.BlockTypeNormal2)
					i++
				} else {
					genBlock(ecs, i, y, components.RandomBlockType().Shorten())
				}
			}
			stage.PathColumn = nextPath
		} else {
			for i := 0; i < consts.BlockColumnNum; i++ {
				if i == path {
					genBlock(ecs, i, y, components.BlockTypeNormal)
				} else {
					genBlock(ecs, i, y, components.RandomBlockType().Shorten())
				}
			}
		}
	case 3:
		for i := 0; i < consts.BlockColumnNum; i++ {
			if i == path {
				genBlock(ecs, i, y, components.BlockTypeReverse)
			} else {
				genBlock(ecs, i, y, components.RandomBlockType().Shorten())
			}
		}
		stage.ShouldReverse = !stage.ShouldReverse
	}
}

func genBlock(ecs *ecs.ECS, col int, line float64, blockType components.BlockType) bool {
	stage := getStage(ecs)
	x := float64(consts.Margin + col*consts.BlockWidth)
	if col == consts.BlockColumnNum-1 {
		blockType = blockType.Shorten()
	}
	if stage.ShouldReverse {
		blockType = blockType.Reverse()
	}
	entry := newBlock(ecs, dmath.NewVec2(x, line), blockType, stage.BlockSpeed)

	block := components.Block.Get(entry)

	if block.IsItem() {
		block.Score = stage.Level * 500
	} else {
		block.Score = stage.Level * 100
	}

	return blockType.IsLong()
}
