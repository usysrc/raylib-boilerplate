package entity

import (
	_ "image/png"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/component"
)

var zCounter component.Counter

func CreateEnemy(cm *component.Manager) component.Entity {
	e := cm.NewEntity()

	cm.Positions[e] = &component.Position{X: rand.Float64() * 640, Y: -64}

	img := rl.LoadImage("internal/assets/enemy.png")

	cm.Velocities[e] = &component.Velocity{X: 0, Y: 1}
	cm.Renders[e] = &component.Render{Image: rl.LoadTextureFromImage(img), Scale: 4, Z: float64(zCounter.NextID())}
	cm.Tags[e] = &component.Tag{Name: "Enemy"}

	rl.UnloadImage(img)

	return e
}
