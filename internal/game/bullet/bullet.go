package bullet

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
)

var bullets []Bullet
var score int
var explosion rl.Sound

type Bullet struct {
	image rl.Texture2D
	Pos   rl.Vector2
	scale float32
	speed float32
	Alive bool
}

func Create(x, y float32) *Bullet {
	b := &Bullet{}
	b.Alive = true
	b.speed = 500
	img := rl.LoadImage("internal/assets/bullet.png")
	b.image = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	b.Pos = rl.Vector2{X: x, Y: y}
	b.scale = 1
	bullets = append(bullets, *b)
	return b
}

func (b *Bullet) Update() {
	// move bullet up
	velocity := rl.Vector2{}
	b.Pos.Y -= b.speed * float32(rl.GetFrameTime())
	b.Pos.X += velocity.X
	b.Pos.Y += velocity.Y

	// check collision with enemies
	enemies := enemy.GetEnemies()
	for i := range enemies {
		if rl.CheckCollisionCircles(b.Pos, 5, enemies[i].Pos, 20) {
			b.Alive = false
			enemies[i].Alive = false
			rl.SetSoundPitch(explosion, 1.0+float32(rl.GetRandomValue(-10, 10))/100)
			rl.PlaySound(explosion)
			score++
		}
	}
}

func (b *Bullet) Draw() {
	rl.DrawTexturePro(b.image, rl.Rectangle{X: 0, Y: 0, Width: float32(b.image.Width), Height: float32(b.image.Height)}, rl.Rectangle{X: b.Pos.X, Y: b.Pos.Y, Width: float32(b.image.Width), Height: float32(b.image.Height)}, rl.Vector2{X: float32(b.image.Width) / 2, Y: float32(b.image.Height) / 2}, 0, rl.White)
}

func Init() {
	// create empty bullet slice
	bullets = make([]Bullet, 0)

	// load explosion sound
	explosion = rl.LoadSound("internal/assets/explosion.wav")

	// set score to 0
	score = 0
}

func Update() {
	for i := 0; i < len(bullets); {
		bullets[i].Update()
		if bullets[i].Alive {
			i++
		} else {
			bullets[i] = bullets[len(bullets)-1]
			bullets = bullets[:len(bullets)-1]
		}
	}
}

func Draw() {
	rl.DrawText(fmt.Sprintf("Score: %d", score), 10, 32, 20, rl.White)
	for i := range bullets {
		bullets[i].Draw()
	}
}
