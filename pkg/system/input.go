package system

import (
	"revdriller/pkg/components"
	"revdriller/pkg/layers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

func newInput(ecs *ecs.ECS) {
	_ = ecs.World.Entry(ecs.Create(
		layers.Default,
		components.Input,
	))
}

func updateInput(ecs *ecs.ECS) {
	input := getInput(ecs)

	input.Reset()

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		input.Down = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		input.Up = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		input.Left = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		input.Right = true
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyZ) {
		input.ButtonA = true
	}
}

func getInput(ecs *ecs.ECS) *components.InputData {
	return components.Input.Get(
		components.Input.MustFirst(ecs.World),
	)
}
