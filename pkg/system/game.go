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
	getGame(ecs).State = components.GameStatePlay

	assets.PlayBGM(assets.BGMMain)
}

func GetScore(ecs *ecs.ECS) int {
	return getGame(ecs).AddScore
}

func IsGameOver(ecs *ecs.ECS) bool {
	return getGame(ecs).State == components.GameStateGameOver
}

func IsGameClear(ecs *ecs.ECS) bool {
	return getGame(ecs).State == components.GameStateClear
}

func newGame(ecs *ecs.ECS, game components.GameData) {
	_ = ecs.World.Entry(ecs.Create(
		layers.System,
		components.Game,
	))

	components.Game.SetValue(
		components.Game.MustFirst(ecs.World),
		game,
	)
}

func isGameStarted(ecs *ecs.ECS) bool {
	return getGame(ecs).State == components.GameStatePlay
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
			getGame(ecs).State = components.GameStateGameOver
		}
		playerPos := transform.WorldPosition(playerEntry)
		playerSize := components.Size.Get(playerEntry)

		if playerPos.Y+playerSize.Y/2 < 0 {
			getGame(ecs).State = components.GameStateClear
		}
	}
}

func drawGame(ecs *ecs.ECS, screen *ebiten.Image) {
	game := getGame(ecs)

	switch game.State {
	case components.GameStateTitle:
		if game.IsFirstPlay() {
			// draw title
			ganim8.DrawSprite(screen, assets.GetSprite("img/title.png"), 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
			// draw high scores
			for i, score := range game.HighScores {
				line := float64(consts.Height/2 + 110 + i*30)
				ganim8.DrawSprite(screen, assets.GetSprite("img/highscore_position.png"), i, 20, line, 0, 1, 1, .0, .0)
				drawNumberR(screen, score, consts.Width-22, line, assets.GetSprite(consts.SmallNumberSprite))
			}
		} else {
			// draw stage number
			ganim8.DrawSprite(screen, assets.GetSprite("img/stage.png"), 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
			drawNumberL(screen, game.Stage, consts.Width/2+50, consts.Height/2-85, assets.GetSprite(consts.LargeNumberSprite))
			// draw life
			drawNumberL(screen, game.Life, consts.Width/2+80, consts.Height/2-29, assets.GetSprite(consts.LargeNumberSprite))
			// draw score
			drawNumberR(screen, game.TotalScore(), consts.Width-20, 10, assets.GetSprite(consts.SmallNumberSprite))
		}
	case components.GameStatePlay:
		// draw score
		drawNumberR(screen, game.TotalScore(), consts.Width-20, 10, assets.GetSprite(consts.SmallNumberSprite))
	case components.GameStateClear:
		// draw stage clear
		ganim8.DrawSprite(screen, assets.GetSprite("img/messages.png"), 1, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
	case components.GameStateGameOver:
		if game.Life > 1 {
			// draw continue
			ganim8.DrawSprite(screen, assets.GetSprite("img/messages.png"), 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
		} else {
			// Gameover
			ganim8.DrawSprite(screen, assets.GetSprite("img/messages.png"), 2, consts.Width/2, consts.Height/2-50, 0, 1, 1, .5, .5)
			// draw high score
			ganim8.DrawSprite(screen, assets.GetSprite("img/highscore.png"), 0, 20, consts.Height/2-10, 0, 1, 1, .0, .5)
			// draw score
			drawNumberR(screen, game.TotalScore(), consts.Width-20, consts.Height/2+25, assets.GetSprite(consts.LargeNumberSprite))
		}
	}
}
