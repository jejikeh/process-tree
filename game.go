package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/entity"
)

const Title = "Process Tree"

const GameWidth = 800
const GameHeight = 600

type Game struct {
	Scene  entity.Entity
	Render Renderer
	UI     *UI

	Scenes []*entity.Scene
}

func NewGame(scene entity.Entity) *Game {
	g := &Game{
		Scene: scene,
	}

	g.Render = NewRenderer(Title, rl.Vector2{X: GameWidth, Y: GameHeight}, 0, 60)

	g.UI = NewUI()

	return g
}

func (g *Game) Run() {
	start := false
	offset := rl.Vector2{X: 0, Y: 0}

	g.Render.Render(func() {
		if !start {
			g.Scene.Start()
			start = true
		}

		g.Scene.Update()

		g.UI.Render()

		g.Scene.Render(offset)
	})
}
