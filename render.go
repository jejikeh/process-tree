package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Renderer interface {
	Render(proc func())
}

type RaylibRenderer struct {
	Title string
	Size  rl.Vector2

	WindowFlags uint32

	TargetFPS int32
}

func NewRenderer(title string, size rl.Vector2, flags uint32, targetFPS int32) *RaylibRenderer {
	r := &RaylibRenderer{
		Title:       title,
		Size:        size,
		WindowFlags: flags,
	}

	rl.SetConfigFlags(r.WindowFlags)

	rl.InitWindow(int32(r.Size.X), int32(r.Size.Y), r.Title)

	rl.SetTargetFPS(targetFPS)

	rl.SetWindowMinSize(int(r.Size.X), int(r.Size.Y))

	return r
}

func (r RaylibRenderer) Render(proc func()) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		proc()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
