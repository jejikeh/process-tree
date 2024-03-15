package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Circle struct {
	*Obj

	Radius float32
	Color  rl.Color
}

func NewCirlce(x, y, r float32, color rl.Color) *Circle {
	return &Circle{
		Obj:    NewObj(rl.Vector2{X: x, Y: y}, rl.Vector2{X: 10, Y: 10}),
		Radius: r,
		Color:  color,
	}
}

func (c Circle) Render(offset rl.Vector2) {
	pos := rl.Vector2Add(c.Position, offset)

	rl.DrawCircleV(pos, c.Radius, c.Color)
}

func (c Circle) Start() {
}

func (c Circle) Update() {
}
