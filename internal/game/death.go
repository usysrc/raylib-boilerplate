package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Death is the game over state
type Death struct {
}

func (d *Death) Init() {
}

func (d *Death) Update(g *Game) error {
	if rl.IsKeyPressed(rl.KeyX) {
		g.Switch("play")
	}
	return nil
}

func (d *Death) Draw() {
	rl.ClearBackground(rl.Black)
	// draw the ship and bullets
	rl.DrawText("YOU DIED!", 200, 200, 64, rl.White)

	rl.DrawText("Press x to restart", 200, 270, 16, rl.White)
}
