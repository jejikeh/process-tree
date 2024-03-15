package entity

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	*Transform

	Text  string
	Color rl.Color
}

func NewText(pos rl.Vector2, text string) *Text {
	return &Text{
		Transform: &Transform{
			Position: pos,
			Size:     rl.Vector2{X: 32, Y: 32},
		},
		Text:  text,
		Color: rl.Black,
	}
}

func (t Text) Update() {
}

func (t Text) Draw() {
	raygui.Button(rl.Rectangle{X: t.Position.X, Y: t.Position.Y, Width: t.Size.X, Height: t.Size.Y}, t.Text)
	// rl.DrawTextEx(Font, t.Text, t.Position, 24, 1, t.Color)
}
