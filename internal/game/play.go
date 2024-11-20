package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/bullet"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
	"github.com/usysrc/raylib-boilerplate/internal/game/ship"
)

// Play is the main game state
type Play struct {
	background rl.Texture2D
}

func (p *Play) Init() {
	img := rl.LoadImage("internal/assets/background.png")
	p.background = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	ship.Init()
	bullet.Init()
	enemy.Init()
}

func (p *Play) Update(g *Game) error {
	ship.Update(g)
	bullet.Update()
	enemy.Update()
	return nil
}

func (p *Play) Draw() {
	// draw the background
	rl.DrawTexture(p.background, 0, 0, rl.White)

	// draw the ship, bullets and enemies
	bullet.Draw()
	enemy.Draw()
	ship.Draw()
}
