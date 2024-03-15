package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	Title     string
	Size      rl.Vector2
	TargetFPS int32

	Flags uint32
	GUI   *Gui
}

const WindowTitle = "process tree"

var WindowSize = rl.Vector2{
	X: 640,
	Y: 480,
}

const WindowFPS = 60
const WindowFlags = rl.FlagWindowHighdpi | rl.FlagWindowResizable

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

	w.GUI = NewGui()

	return w
}

func (w *Window) Run(run func()) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		run()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
