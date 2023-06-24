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
	getGame(ecs).IsGameStart = true

	assets.PlayBGM(assets.BGMMain)
}

func newGame(ecs *ecs.ECS, stage, life int) {
	_ = ecs.World.Entry(ecs.Create(
		layers.System,
		components.Game,
	))

	game := getGame(ecs)
	game.Stage = stage
	game.Life = life
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
			getGame(ecs).IsDead = true
		}
		playerPos := transform.WorldPosition(playerEntry)
		playerSize := components.Size.Get(playerEntry)

		if playerPos.Y+playerSize.Y/2 < 0 {
			getGame(ecs).IsClear = true
		}
	}
}

func drawGame(ecs *ecs.ECS, screen *ebiten.Image) {
	game := getGame(ecs)

	// if the game is not started, draw title or stage screen.
	if !game.IsGameStart {
		if game.IsFirstPlay() {
			spr := assets.GetSprite("img/title.png")
			ganim8.DrawSprite(screen, spr, 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
		} else {
			spr := assets.GetSprite("img/stage.png")
			ganim8.DrawSprite(screen, spr, 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
			// draw stage number
			drawNumber(screen, game.Stage, consts.Width/2+50, consts.Height/2-85)
			// draw life
			drawNumber(screen, game.Life, consts.Width/2+80, consts.Height/2-29)
		}
	}

	// if the player is dead, draw lose or game over.
	if game.IsDead {
		spr := assets.GetSprite("img/messages.png")
		if game.Life > 1 {
			ganim8.DrawSprite(screen, spr, 0,
				consts.Width/2, consts.Height/2,
				0, 1, 1, .5, .5)
		} else {
			// Gameover
			ganim8.DrawSprite(screen, spr, 2,
				consts.Width/2, consts.Height/2,
				0, 1, 1, .5, .5)
		}
	}

	// if the game is clear, draw game clear.
	if game.IsClear {
		spr := assets.GetSprite("img/messages.png")
		ganim8.DrawSprite(screen, spr, 1,
			consts.Width/2, consts.Height/2,
			0, 1, 1, .5, .5)
	}
}
