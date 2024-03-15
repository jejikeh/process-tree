package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/entity"
)

const ThemeSwitcherWidth = 124
const ThemeSwitcherHeight = 180

const buttonOffsetX = 12
const buttonOffsetY = 32

type ThemeSwitcher struct {
	*entity.Obj

	Visible bool
}

func NewThemeSwitcher(x, y float32) *ThemeSwitcher {
	return &ThemeSwitcher{
		Obj: entity.NewObj(rl.Vector2{X: x - ThemeSwitcherWidth/2, Y: y - ThemeSwitcherHeight/2}, rl.Vector2{X: ThemeSwitcherWidth, Y: ThemeSwitcherHeight}),
	}
}

func (t *ThemeSwitcher) Start() {
	t.Visible = false
}

func (t *ThemeSwitcher) Update() {
	if rl.IsKeyPressed(rl.KeySpace) {
		t.Visible = !t.Visible
	}
}

func (t *ThemeSwitcher) Render(offset rl.Vector2) {
	if !t.Visible {
		return
	}

	pos := t.GetRectangle()

	if raygui.WindowBox(pos, "Theme switcher") {
		t.Visible = false
	}

	buttonPos := rl.Rectangle{X: pos.X + buttonOffsetX, Y: pos.Y + buttonOffsetY, Width: 100, Height: 20}

	if raygui.Button(buttonPos, "Terminal") {
		GameInstance.UI.SetStyle(Terminal)
	}

	buttonPos.Y += buttonOffsetY

	if raygui.Button(buttonPos, "Dark") {
		GameInstance.UI.SetStyle(Dark)
	}

}
