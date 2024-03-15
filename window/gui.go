package window

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Font rl.Font

type Gui struct {
	Style string
}

const GuiStyle = "./assets/themes/cherry/cherry.rgs"

func NewGui() *Gui {
	// @Incomplete: Get theese values from .conf file
	g := &Gui{
		Style: GuiStyle,
	}

	gui.LoadStyle(g.Style)

	return g
}
