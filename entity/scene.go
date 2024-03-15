package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Scene struct {
	*Obj
	*EntityBucket

	Title string
}

func NewScene(title string, x, y int32) *Scene {
	return &Scene{
		Obj:          NewObj(rl.Vector2{X: float32(x), Y: float32(y)}, rl.Vector2{X: 0, Y: 0}),
		EntityBucket: NewEntityBucket(),
		Title:        title,
	}
}
