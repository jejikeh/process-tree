package entity

import rl "github.com/gen2brain/raylib-go/raylib"

var Font rl.Font

type Text struct {
	*Transform

	Text  string
	Color rl.Color
}

func NewText(pos rl.Vector2, text string) *Text {
	return &Text{
		Transform: &Transform{
			Position: pos,
		},
		Text:  text,
		Color: rl.Black,
	}
}

func (t Text) Update() {
}

func (t Text) Draw() {
	rl.DrawTextEx(Font, t.Text, t.Position, 24, 1, t.Color)
}
