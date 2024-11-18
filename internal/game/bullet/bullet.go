package bullet

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
)

var bullets []Bullet

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
	b.scale = 4
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
		if rl.CheckCollisionRecs(rl.Rectangle{X: b.Pos.X, Y: b.Pos.Y, Width: 16, Height: 16}, rl.Rectangle{X: enemies[i].Pos.X, Y: enemies[i].Pos.Y, Width: 16, Height: 16}) {
			b.Alive = false
			enemies[i].Alive = false
		}
	}
}

func (b *Bullet) Draw() {
	rl.DrawTextureEx(b.image, b.Pos, 0, b.scale, rl.Black)
}

func Init() {
	// create empty bullet slice
	bullets = make([]Bullet, 0)
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
	for i := range bullets {
		bullets[i].Draw()
	}
}
