package entity

import (
	_ "image/png"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/component"
)

func CreateBullet(cm *component.Manager) component.Entity {
	e := cm.NewEntity()

	cm.Positions[e] = &component.Position{X: 100, Y: 100}
	cm.Velocities[e] = &component.Velocity{}

	img := rl.LoadImage("internal/assets/bullet.png")

	cm.Renders[e] = &component.Render{Image: rl.LoadTextureFromImage(img), Scale: 4, Z: 10}
	cm.Tags[e] = &component.Tag{Name: "Bullet"}

	rl.UnloadImage(img)

	return e
}
