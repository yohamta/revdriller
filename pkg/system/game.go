package system

import (
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/layers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func newGame(ecs *ecs.ECS) {
	_ = ecs.World.Entry(ecs.Create(
		layers.Default,
		components.Game,
	))
}

func updateGame(ecs *ecs.ECS) {
	game := components.Game.MustFirst(ecs.World)

	player, ok := components.Player.First(ecs.World)
	if ok {
		if components.Player.Get(player).IsDead {
			components.Game.Get(game).IsGameOver = true
		}
	}
}

func drawGame(ecs *ecs.ECS, screen *ebiten.Image) {
	game := components.Game.Get(
		components.Game.MustFirst(ecs.World),
	)

	if game.IsGameOver {
		spr := assets.GetSprite("img/messages.png")
		ganim8.DrawSprite(screen, spr, 0,
			consts.Width/2, consts.Height/2,
			0, 1, 1, .5, .5)
	}
}
