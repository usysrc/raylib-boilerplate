package main

import (
	_ "image/png"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game"
)

var myGame *game.Game

func init() {}

func main() {
	myGame = &game.Game{}
	rl.InitWindow(800, 600, "raylib boilerplate")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetTraceLogLevel(rl.LogError)
	myGame.Init()

	for !rl.WindowShouldClose() {
		myGame.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		myGame.Draw()
		rl.EndDrawing()
	}
}
