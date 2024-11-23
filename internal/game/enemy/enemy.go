package enemy

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var enemies []Enemy
var enemyTexture rl.Texture2D

type Enemy struct {
	image rl.Texture2D
	Pos   rl.Vector2
	scale float32
	speed float32
	Alive bool
}

func Create(x, y float32) *Enemy {
	e := &Enemy{}
	e.Alive = true
	e.speed = float32(rl.GetRandomValue(50, 100))
	e.image = enemyTexture
	e.Pos = rl.Vector2{X: x, Y: y}
	e.scale = 1
	enemies = append(enemies, *e)
	return e
}

func (e *Enemy) Update() {
	velocity := rl.Vector2{}
	e.Pos.Y += e.speed * float32(rl.GetFrameTime())
	e.Pos.X += velocity.X
	e.Pos.Y += velocity.Y
	if e.Pos.Y > 600 {
		e.Alive = false
	}
}

type Position interface {
	GetPos() rl.Vector2
}

func (e *Enemy) Draw(ship Position) {
	shipPos := ship.GetPos()
	// check if enemy is above the ship
	if rl.CheckCollisionCircleRec(e.Pos, 20, rl.Rectangle{X: shipPos.X, Y: 0, Width: 2, Height: float32(rl.GetScreenHeight())}) {
		// draw a unfilled rectangle around the enemy
		rl.DrawRectangleLines(int32(e.Pos.X-20), int32(e.Pos.Y-20), 40, 40, rl.Red)
	}

	rl.DrawTexturePro(e.image, rl.Rectangle{X: 0, Y: 0, Width: float32(e.image.Width), Height: float32(e.image.Height)}, rl.Rectangle{X: e.Pos.X, Y: e.Pos.Y, Width: float32(e.image.Width), Height: float32(e.image.Height)}, rl.Vector2{X: float32(e.image.Width) / 2, Y: float32(e.image.Height) / 2}, 0, rl.White)
}

func Init() {
	img := rl.LoadImage("internal/assets/enemy.png")
	enemyTexture = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)

	// create empty enemies slice
	enemies = make([]Enemy, 0)
}

func Update() {
	// create new enemy if random number is less than 0.01
	if rl.GetRandomValue(0, 100) < 1 {
		Create(float32(rl.GetRandomValue(0, 800)), 0)
	}
	// update all enemies
	for i := 0; i < len(enemies); {
		enemies[i].Update()
		if enemies[i].Alive {
			i++
		} else {
			enemies[i] = enemies[len(enemies)-1]
			enemies = enemies[:len(enemies)-1]
		}
	}
}

func Draw(ship Position) {
	rl.DrawText(fmt.Sprintf("Enemies: %d", len(enemies)), 10, 10, 20, rl.White)
	for i := range enemies {
		enemies[i].Draw(ship)
	}
}

func GetEnemies() []Enemy {
	return enemies
}
