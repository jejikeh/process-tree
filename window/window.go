package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/entity"
)

type Window struct {
	Title     string
	Size      rl.Vector2
	TargetFPS int32

	Flags uint32
	GUI   *Gui

	EntityManager *entity.EntityManager
}

const WindowTitle = "process tree"

var WindowSize = rl.Vector2{
	X: 640,
	Y: 480,
}

const WindowFPS = 60
const WindowFlags = rl.FlagWindowResizable | rl.FlagMsaa4xHint | rl.FlagWindowHighdpi

func NewWindow() *Window {
	// @Incomplete: get theese values from .conf file
	w := &Window{
		Title:     WindowTitle,
		Size:      WindowSize,
		TargetFPS: WindowFPS,
		Flags:     WindowFlags,
	}

	rl.SetConfigFlags(w.Flags)
	rl.InitWindow(int32(w.Size.X), int32(w.Size.Y), w.Title)
	rl.SetTargetFPS(w.TargetFPS)

	rl.SetWindowMinSize(int(w.Size.X), int(w.Size.Y))

	w.GUI = NewGui()

	w.EntityManager = &entity.EntityManager{
		Entities: make([]entity.Entity, 0),
	}

	return w
}

func (w *Window) Run(update func(), draw func()) {
	for !rl.WindowShouldClose() {
		update()

		rl.BeginDrawing()

		draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
