package enemy

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var enemies []Enemy

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
	e.speed = 100
	img := rl.LoadImage("internal/assets/enemy.png")
	e.image = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	e.Pos = rl.Vector2{X: x, Y: y}
	e.scale = 4
	enemies = append(enemies, *e)
	return e
}

func (e *Enemy) Update() {
	velocity := rl.Vector2{}
	e.Pos.Y += e.speed * float32(rl.GetFrameTime())
	e.Pos.X += velocity.X
	e.Pos.Y += velocity.Y
}

func (e *Enemy) Draw() {
	rl.DrawTextureEx(e.image, e.Pos, 0, e.scale, rl.White)
}

func Init() {
	// create empty enemies slice
	enemies = make([]Enemy, 0)
}

func Update() {
	// create new enemy if random number is less than 0.01
	if rl.GetRandomValue(0, 100) < 1 {
		Create(float32(rl.GetRandomValue(0, 800)), 0)
	}
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

func Draw() {
	for i := range enemies {
		enemies[i].Draw()
	}
}

func GetEnemies() []Enemy {
	return enemies
}
