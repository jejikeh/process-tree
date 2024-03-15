package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Obj struct {
	Position rl.Vector2
	Size     rl.Vector2
}

func NewObj(position rl.Vector2, size rl.Vector2) *Obj {
	return &Obj{
		Position: position,
		Size:     size,
	}
}

func (o *Obj) GetRectangle() rl.Rectangle {
	return rl.NewRectangle(o.Position.X, o.Position.Y, o.Size.X, o.Size.Y)
}
