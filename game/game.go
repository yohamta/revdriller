package game

import (
	"log"
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/system"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

type game struct {
	state          state
	ecs            *ecs.ECS
	stateChangedAt time.Time
}

var (
	setupOnce sync.Once
)

var _ ebiten.Game = &game{}

type state int

const (
	stateInit state = iota
	stateStart
	statePlaying
	stateGameover
	stateGameclear
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
	case stateStart, statePlaying, stateGameover, stateGameclear:
		g.ecs.Draw(screen)
	}

	if g.state == stateStart {
		spr := assets.GetSprite("img/title.png")
		ganim8.DrawSprite(screen, spr, 0, consts.Width/2, consts.Height/2, 0, 1, 1, .5, .5)
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
		if g.checkStart() && g.stateDuration() > consts.StateDuration {
			system.StartGame(g.ecs)
			g.changeState(statePlaying)
		}
	case statePlaying:
		g.ecs.Update()
		if g.checkGameOver() {
			g.changeState(stateGameover)
		}
		if g.checkGameClear() {
			g.changeState(stateGameclear)
		}
	case stateGameclear, stateGameover:
		if g.checkRestart() {
			g.changeState(stateInit)
		}
	}

	return nil
}

func (g *game) stateDuration() time.Duration {
	return time.Since(g.stateChangedAt)
}

func (g *game) changeState(state state) {
	g.state = state
	g.stateChangedAt = time.Now()
}

func (g *game) checkStart() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyZ)
}

func (g *game) checkRestart() bool {
	if g.stateDuration() < consts.StateDuration {
		return false
	}
	return ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyZ)
}

// checkGameOver checks if the game is over.
func (g *game) checkGameOver() bool {
	game := components.Game.MustFirst(g.ecs.World)
	return components.Game.Get(game).IsGameOver
}

func (g *game) checkGameClear() bool {
	game := components.Game.MustFirst(g.ecs.World)
	return components.Game.Get(game).IsGameClear
}

// initStage initializes the stage.
func (g *game) initStage() {
	// Create a new ECS world.
	g.ecs = ecs.NewECS(donburi.NewWorld())

	// Setup systems.
	system.Setup(g.ecs)

	// Start the game.
	g.changeState(stateStart)
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
