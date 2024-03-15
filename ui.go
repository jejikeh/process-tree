package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIStyle string

const (
	Terminal UIStyle = "assets/themes/terminal/style_terminal.txt.rgs"
	Dark     UIStyle = "assets/themes/dark/style_dark.txt.rgs"
)

type UI struct {
	Style UIStyle
	Font  rl.Font
}

func NewUI() *UI {
	ui := &UI{
		Font: rl.LoadFont("assets/fonts/Martel-Regular.ttf"),
	}

	ui.SetStyle(Terminal)

	return ui
}

func (ui *UI) SetStyle(style UIStyle) {
	ui.Style = style
	raygui.LoadStyle(string(ui.Style))
}

func (ui *UI) Render() {
	switch ui.Style {
	case Terminal:
		rl.ClearBackground(rl.Black)
	case Dark:
		rl.ClearBackground(rl.Gray)
	}

	raygui.SetFont(ui.Font)
}
