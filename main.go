package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/entity"
)

var GameInstance *Game

func main() {
	// sample := "./samples/file/random/"

	// _, err := filetree.InitFileTree(sample)

	// if err != nil {
	// 	panic(err)
	// }

	GameInstance = NewGame(NewMainScene())

	GameInstance.Run()
}

type MainScene struct {
	*entity.Scene
	Circle entity.Circle
}

func NewMainScene() *MainScene {
	m := &MainScene{
		Scene:  entity.NewScene("Main", 0, 0),
		Circle: *entity.NewCirlce(400, 300, 100, rl.Red),
	}

	m.EntityBucket.AddEntity(m.Circle)

	m.EntityBucket.AddEntity(NewThemeSwitcher(GameWidth/2, GameHeight/2-124))

	return m
}

func (m *MainScene) Start() {
	m.EntityBucket.Start()
}

func (m *MainScene) Update() {
	m.EntityBucket.Update()
}

func (m *MainScene) Render(offset rl.Vector2) {
	m.EntityBucket.Render(m.Position)
}
