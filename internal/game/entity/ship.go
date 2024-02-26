package entity

import (
	_ "image/png"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/component"
)

func CreateShip(cm *component.Manager) component.Entity {
	e := cm.NewEntity()

	cm.Positions[e] = &component.Position{X: 400, Y: 500}
	cm.Velocities[e] = &component.Velocity{}

	img := rl.LoadImage("internal/assets/ship.png")

	cm.Renders[e] = &component.Render{Image: rl.LoadTextureFromImage(img), Scale: 4}
	cm.Tags[e] = &component.Tag{Name: "Ship"}

	rl.UnloadImage(img)

	return e
}
