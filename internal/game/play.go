package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/bullet"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
	"github.com/usysrc/raylib-boilerplate/internal/game/particle"
	"github.com/usysrc/raylib-boilerplate/internal/game/ship"
)

// Play is the main game state
type Play struct {
	background rl.Texture2D
	Ship       *ship.Ship
}

func (p *Play) Init() {
	img := rl.LoadImage("internal/assets/background.png")
	p.background = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	p.Ship = ship.NewShip()
	bullet.Init()
	enemy.Init()
	particle.Init()
}

func (p *Play) Update(g *Game) error {
	p.Ship.Update(g)
	bullet.Update()
	enemy.Update()
	particle.Update()
	return nil
}

func (p *Play) Draw() {
	// draw the background
	rl.DrawTexture(p.background, 0, 0, rl.White)

	// draw the ship, bullets and enemies
	bullet.Draw()
	enemy.Draw(p.Ship)
	p.Ship.Draw()
	particle.Draw()
}
