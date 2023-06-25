package game

import (
	"revdriller/assets"
	"revdriller/pkg/components"
	"revdriller/pkg/consts"
	"revdriller/pkg/system"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"golang.org/x/exp/slices"
)

type game struct {
	state          state
	ecs            *ecs.ECS
	stateChangedAt time.Time
	score          int
	stage          int
	life           int
	highscores     []int
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
		state:      stateInit,
		stage:      1,
		score:      0,
		life:       consts.Life,
		ecs:        ecs.NewECS(donburi.NewWorld()),
		highscores: []int{300, 200, 100},
	}
}

// Draw implements ebiten.game.
func (g *game) Draw(screen *ebiten.Image) {
	screen.Clear()

	if g.ecs == nil {
		return
	}

	switch g.state {
	case stateInit:
		// Do nothing.
	case stateStart, statePlaying, stateGameover, stateGameclear:
		g.ecs.Draw(screen)
	}
}

// Layout implements ebiten.game.
func (*game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return consts.Width, consts.Height
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
		if system.IsGameClear(g.ecs) {
			g.changeState(stateGameclear)
		} else if system.IsGameOver(g.ecs) {
			g.changeState(stateGameover)
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
	if g.state == state {
		return
	}

	switch state {
	case stateStart:
	case stateGameclear:
		g.score += system.GetScore(g.ecs)
		g.stage++
	case stateGameover:
		g.score += system.GetScore(g.ecs)
		g.life--
		if g.life <= 0 {
			g.updateHighScore()
			g.reset()
		}
	}

	g.state = state
	g.stateChangedAt = time.Now()
}

func (g *game) reset() {
	g.life = consts.Life
	g.stage = 1
	g.score = 0
}

func (g *game) updateHighScore() {
	g.highscores = append(g.highscores, g.score)
	slices.SortFunc(g.highscores, func(a, b int) bool {
		return a > b
	})
	if len(g.highscores) > consts.HighScoreNum {
		g.highscores = g.highscores[:consts.HighScoreNum]
	}
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

// initStage initializes the stage.
func (g *game) initStage() {
	// Create a new ECS world.
	g.ecs = ecs.NewECS(donburi.NewWorld())

	// Setup systems.
	system.Setup(g.ecs, components.GameData{
		Life:       g.life,
		Stage:      g.stage,
		Score:      g.score,
		HighScores: g.highscores,
	})

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
			panic(err)
		}
	}
}
