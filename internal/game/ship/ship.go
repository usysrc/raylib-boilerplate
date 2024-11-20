package ship

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/bullet"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
)

var shipImage rl.Texture2D
var shipPos rl.Vector2
var speed float32
var Alive bool
var velocity rl.Vector2
var maxVelocity float32
var snd rl.Sound

func Init() {
	Alive = false
	velocity = rl.Vector2{}
	maxVelocity = 20
	speed = 10
	img := rl.LoadImage("internal/assets/ship.png")
	shipImage = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	shipPos = rl.Vector2{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight()) - 100}
	snd = rl.LoadSound("internal/assets/laser.wav")
}

type GamestateSwitcher interface {
	Switch(to string)
}

func Update(g GamestateSwitcher) {
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		if velocity.Y > 0 {
			velocity.Y = 0
		}
		velocity.Y -= 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		if velocity.Y < 0 {
			velocity.Y = 0
		}
		velocity.Y += 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		if velocity.X > 0 {
			velocity.X = 0
		}
		velocity.X -= 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		if velocity.X < 0 {
			velocity.X = 0
		}
		velocity.X += 1.0 * float32(rl.GetFrameTime()) * speed
	}
	if !rl.IsKeyDown(rl.KeyUp) && !rl.IsKeyDown(rl.KeyW) && !rl.IsKeyDown(rl.KeyDown) && !rl.IsKeyDown(rl.KeyS) {
		velocity.Y *= 0.95
	}
	if !rl.IsKeyDown(rl.KeyLeft) && !rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyRight) && !rl.IsKeyDown(rl.KeyD) {
		velocity.X *= 0.95
	}
	if velocity.X > maxVelocity {
		velocity.X = maxVelocity
	} else if velocity.X < -maxVelocity {
		velocity.X = -maxVelocity
	}
	if velocity.Y > maxVelocity {
		velocity.Y = maxVelocity
	} else if velocity.Y < -maxVelocity {
		velocity.Y = -maxVelocity
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		rl.PlaySound(snd)
		bullet.Create(shipPos.X, shipPos.Y)
	}
	targetPos := rl.Vector2{X: shipPos.X + velocity.X, Y: shipPos.Y + velocity.Y}
	if targetPos.X > 0 && targetPos.X < float32(rl.GetScreenWidth()) {
		shipPos.X += velocity.X
	} else {
		velocity.X = 0
	}
	if targetPos.Y > 0 && targetPos.Y < float32(rl.GetScreenHeight()) {
		shipPos.Y += velocity.Y
	} else {
		velocity.Y = 0
	}
	enemies := enemy.GetEnemies()
	for i := range enemies {
		if rl.CheckCollisionCircles(shipPos, 32, enemies[i].Pos, 20) {
			g.Switch("death")
			return
		}
	}
}

func Draw() {
	//draw image with the shipPos being at the center of the image
	rl.DrawTexturePro(shipImage, rl.Rectangle{X: 0, Y: 0, Width: float32(shipImage.Width), Height: float32(shipImage.Height)}, rl.Rectangle{X: shipPos.X, Y: shipPos.Y, Width: float32(shipImage.Width), Height: float32(shipImage.Height)}, rl.Vector2{X: float32(shipImage.Width) / 2, Y: float32(shipImage.Height) / 2}, 0, rl.White)
}
