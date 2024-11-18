package ship

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/bullet"
)

var shipImage rl.Texture2D
var shipPos rl.Vector2
var scale float32
var speed float32

func Init() {
	speed = 100
	img := rl.LoadImage("internal/assets/ship.png")
	shipImage = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	shipPos = rl.Vector2{X: 200, Y: 200}
	scale = 4
}

func Update() {
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
}

func Draw() {
	rl.DrawTextureEx(shipImage, shipPos, 0, scale, rl.White)
}
