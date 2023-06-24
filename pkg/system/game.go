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

func newGame(ecs *ecs.ECS, stage, life, score int) {
	_ = ecs.World.Entry(ecs.Create(
		layers.System,
		components.Game,
	))

	game := getGame(ecs)
	game.Stage = stage
	game.Life = life
	game.Score = score
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

	// TODO: refactor these codes

	// if the game is not started, draw title or stage screen.
	if !game.IsGameStart {
		if game.IsFirstPlay() {
			ganim8.DrawSprite(screen, assets.GetSprite("img/title.png"), 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
		} else {
			ganim8.DrawSprite(screen, assets.GetSprite("img/stage.png"), 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
			// draw stage number
			drawNumberL(screen, game.Stage, consts.Width/2+50, consts.Height/2-85, assets.GetSprite("img/numbers.png"))
			// draw life
			drawNumberL(screen, game.Life, consts.Width/2+80, consts.Height/2-29, assets.GetSprite("img/numbers.png"))
		}
	}

	// if the player is dead, draw lose or game over.
	if game.IsDead {
		if game.Life > 1 {
			ganim8.DrawSprite(screen, assets.GetSprite("img/messages.png"), 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
		} else {
			// Gameover
			ganim8.DrawSprite(screen, assets.GetSprite("img/messages.png"), 2, consts.Width/2, consts.Height/2-50, 0, 1, 1, .5, .5)
			// draw high score
			ganim8.DrawSprite(screen, assets.GetSprite("img/highscore.png"), 0, 20, consts.Height/2-10, 0, 1, 1, .0, .5)
			// draw score
			drawNumberR(screen, game.Score, consts.Width-20, consts.Height/2+25, assets.GetSprite("img/numbers.png"))
		}
	} else if game.IsClear {
		// if the game is clear, draw game clear.
		ganim8.DrawSprite(screen, assets.GetSprite("img/messages.png"), 1, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
	}

	if (game.IsGameStart || game.IsClear) && !(game.IsDead && game.Life == 1) {
		// draw score
		drawNumberR(screen, game.Score+game.AddScore, consts.Width-20, 10, assets.GetSprite("img/numbers_small.png"))
	}
}
