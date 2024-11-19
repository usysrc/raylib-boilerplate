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
	if rl.IsKeyPressed(rl.KeySpace) {
		g.Switch("play")
	}
	return nil
}

func (d *Death) Draw() {
	// draw the ship and bullets
	rl.DrawText("You died", 200, 200, 20, rl.White)
}
