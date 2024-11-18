package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/bullet"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
	"github.com/usysrc/raylib-boilerplate/internal/game/ship"
)

type Game struct {
	background rl.Texture2D
}

func (g *Game) Init() {
	img := rl.LoadImage("internal/assets/background.png")
	g.background = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	ship.Init()
	bullet.Init()
	enemy.Init()
}

func (g *Game) Update() error {
	ship.Update()
	bullet.Update()
	enemy.Update()
	return nil
}

func (g *Game) Draw() {
	// draw the background
	rl.DrawTexture(g.background, 0, 0, rl.White)
	// draw the ship and bullets
	ship.Draw()
	bullet.Draw()
	enemy.Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
