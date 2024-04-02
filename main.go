package main

import (
	// gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/treemap/filemap"
)

const (
	WindowWidth  = 640
	WindowHeight = 480
)

const fontPath = "assets/Lora-Regular.ttf"
const fontSpacing = 4

var fontSize float32 = 0.7

var font rl.Font

func main() {
	rl.InitWindow(WindowWidth, WindowHeight, "raygui - button")

	rl.SetTargetFPS(60)

	_, err := filemap.InitTreemap("samples/dvvf")

	if err != nil {
		panic(err)
	}

	font = rl.LoadFont(fontPath)
	rl.GenTextureMipmaps(&font.Texture)
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		rl.DrawRectangleGradientV(0, WindowHeight-font.BaseSize, WindowWidth, font.BaseSize*2, rl.White, rl.Black)
		rl.DrawTextEx(font, "Hello, World", rl.NewVector2(WindowWidth/25, float32(WindowHeight-float32(font.BaseSize)*fontSize)), float32(font.BaseSize)*fontSize, float32(fontSpacing/font.BaseSize), rl.Black)

		drawTreemap()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func drawTreemap() {
	var x0 int32 = WindowWidth / 25
	var x1 int32 = WindowWidth - x0

	var y0 int32 = x0
	var y1 int32 = WindowHeight - y0

	rl.DrawRectangleGradientV(x0, y0, x1-x0, y1-y0, rl.Blue, rl.Red)
}
