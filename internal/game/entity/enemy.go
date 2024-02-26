package entity

import (
	_ "image/png"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/component"
)

func CreateEnemy(cm *component.Manager) component.Entity {
	e := cm.NewEntity()

	cm.Positions[e] = &component.Position{X: rand.Float64() * 800.0, Y: -32}

	img := rl.LoadImage("internal/assets/enemy.png")

	cm.Velocities[e] = &component.Velocity{X: 0, Y: 1}
	cm.Renders[e] = &component.Render{Image: rl.LoadTextureFromImage(img), Scale: 4}
	cm.Tags[e] = &component.Tag{Name: "Enemy"}

	rl.UnloadImage(img)

	return e
}
