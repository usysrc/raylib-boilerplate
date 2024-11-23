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
	rl.SetTraceLogLevel(rl.LogError)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(800, 600, "raylib boilerplate")
	rl.InitAudioDevice()
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetTraceLogLevel(rl.LogError)
	myGame.Init()

	for !myGame.ShouldClose() {
		myGame.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		myGame.Draw()
		rl.EndDrawing()
	}
	rl.CloseAudioDevice()
}
