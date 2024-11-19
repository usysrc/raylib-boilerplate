package ship

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/bullet"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
)

var shipImage rl.Texture2D
var shipPos rl.Vector2
var scale float32
var speed float32
var Alive bool

func Init() {
	Alive = false
	speed = 200
	img := rl.LoadImage("internal/assets/ship.png")
	shipImage = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	shipPos = rl.Vector2{X: 200, Y: 200}
	scale = 4
}

type GamestateSwitcher interface {
	Switch(to string)
}

func Update(g GamestateSwitcher) {
	velocity := rl.Vector2{}
	if rl.IsKeyDown(rl.KeyUp) {
		velocity.Y -= 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		velocity.Y += 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		velocity.X -= 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyRight) {
		velocity.X += 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		bullet.Create(shipPos.X, shipPos.Y)
	}
	shipPos.X += velocity.X
	shipPos.Y += velocity.Y
	enemies := enemy.GetEnemies()
	for i := range enemies {
		if rl.CheckCollisionRecs(rl.Rectangle{X: shipPos.X, Y: shipPos.Y, Width: 16, Height: 16}, rl.Rectangle{X: enemies[i].Pos.X, Y: enemies[i].Pos.Y, Width: 16, Height: 16}) {
			g.Switch("death")
			return
		}
	}
}

func Draw() {
	rl.DrawTextureEx(shipImage, shipPos, 0, scale, rl.White)
}
