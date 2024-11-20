package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Loading is the loading state
type Loading struct {
}

var t float32

func (d *Loading) Init() {
	t = 0
}

func (d *Loading) Update(g *Game) error {
	t += rl.GetFrameTime() * 500
	if t > 1000 {
		g.Switch("play")
	}
	return nil
}

func (d *Loading) Draw() {
	rl.ClearBackground(rl.Black)
	// draw the ship and bullets

	rl.DrawCircle(400, 300, 100, rl.DarkGray)
	rl.DrawCircleSector(rl.NewVector2(400, 300), 100, 0+t, 300+t, 100, rl.White)
	rl.DrawCircle(400, 300, 50, rl.Black)
}
