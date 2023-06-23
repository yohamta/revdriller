package system

import (
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/layers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/ganim8/v2"
)

func StartGame(ecs *ecs.ECS) {
	newStage(ecs, 1)

	getGame(ecs).IsGameStart = true
}

func newGame(ecs *ecs.ECS) {
	_ = ecs.World.Entry(ecs.Create(
		layers.System,
		components.Game,
	))
}

func isGameStarted(ecs *ecs.ECS) bool {
	return getGame(ecs).IsGameStart
}

func getGame(ecs *ecs.ECS) *components.GameData {
	return components.Game.Get(
		components.Game.MustFirst(ecs.World),
	)
}

func updateGame(ecs *ecs.ECS) {
	playerEntry, ok := components.Player.First(ecs.World)
	if ok {
		if components.Player.Get(playerEntry).IsDead {
			getGame(ecs).IsGameOver = true
		}
		playerPos := transform.WorldPosition(playerEntry)
		playerSize := components.Size.Get(playerEntry)

		if playerPos.Y+playerSize.Y/2 < 0 {
			getGame(ecs).IsGameClear = true
		}
	}
}

func drawGame(ecs *ecs.ECS, screen *ebiten.Image) {
	if getGame(ecs).IsGameOver {
		// draw game over
		spr := assets.GetSprite("img/messages.png")
		ganim8.DrawSprite(screen, spr, 0,
			consts.Width/2, consts.Height/2,
			0, 1, 1, .5, .5)
	}

	if getGame(ecs).IsGameClear {
		// draw game clear
		spr := assets.GetSprite("img/messages.png")
		ganim8.DrawSprite(screen, spr, 1,
			consts.Width/2, consts.Height/2,
			0, 1, 1, .5, .5)
	}
}
