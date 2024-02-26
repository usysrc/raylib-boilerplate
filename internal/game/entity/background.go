package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/component"
)

func CreateBackground(cm *component.Manager) component.Entity {
	e := cm.NewEntity()

	cm.Positions[e] = &component.Position{X: 0, Y: 0}

	img := rl.LoadImage("internal/assets/background.png")

	cm.Renders[e] = &component.Render{Image: rl.LoadTextureFromImage(img), Scale: 1, Z: -1000}

	rl.UnloadImage(img)

	return e
}
