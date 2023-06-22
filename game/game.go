package game

import (
	"log"
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/entity"
	"revdriller/pkg/layers"
	"revdriller/pkg/system"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type game struct {
	state state
	ecs   *ecs.ECS
}

var (
	setupOnce sync.Once
)

var _ ebiten.Game = &game{}

type state int

const (
	stateInit state = iota
	stateStart
	stateGameover
)

func New() *game {
	return &game{
		state: stateInit,
	}
}

// Draw implements ebiten.game.
func (g *game) Draw(screen *ebiten.Image) {
	screen.Clear()

	switch g.state {
	case stateInit:
		// Do nothing.
	case stateStart:
		g.ecs.Draw(screen)
	}
}

// Layout implements ebiten.game.
func (*game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 240, 320
	// factor := ebiten.DeviceScaleFactor()
	// return int(float64(outsideWidth) * factor), int(float64(outsideHeight) * factor)
}

// Update implements ebiten.game.
func (g *game) Update() error {
	setupOnce.Do(setup)

	switch g.state {
	case stateInit:
		g.initStage()
	case stateStart:
		g.ecs.Update()
		if g.checkGameOver() {
			println("game over")
			g.state = stateGameover
		}
	}

	return nil
}

// checkGameOver checks if the game is over.
func (g *game) checkGameOver() bool {
	game := components.Game.MustFirst(g.ecs.World)
	return components.Game.Get(game).IsGameOver
}

// initStage initializes the stage.
func (g *game) initStage() {
	g.ecs = ecs.NewECS(donburi.NewWorld())

	// Add systems.
	g.ecs.AddSystem(system.UpdateAnimation)
	g.ecs.AddSystem(system.UpdatePlayer)
	g.ecs.AddSystem(system.UpdateGame)

	// Add renderers.
	g.ecs.AddRenderer(layers.Player, system.DrawAnimation)

	// Add entities.
	entity.NewGame(g.ecs)
	entity.NewPlayer(g.ecs)

	// Start the game.
	g.state = stateStart
}

// setup loads assets and initializes the game.
func setup() {
	loadAssets()
}

// loadAssets loads all assets.
func loadAssets() {
	for _, fn := range []func() error{
		assets.Load,
	} {
		if err := fn(); err != nil {
			log.Fatal(err)
		}
	}
}
