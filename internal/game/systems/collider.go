package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/component"
)

type Collider struct {
	Components *component.Manager
}

func (c *Collider) Update() {
	w := 32.0
	h := 32.0
	for e1, position1 := range c.Components.Positions {
		for e2, position2 := range c.Components.Positions {
			if position1.X > position2.X && position1.X < position2.X+w && position1.Y > position2.Y && position1.Y < position2.Y+h {
				c.Collide(e1, e2)
			}
		}
	}
}

func (c *Collider) Collide(e1, e2 component.Entity) {
	tag1, exist := c.Components.Tags[e1]
	if !exist {
		return
	}
	tag2, exist := c.Components.Tags[e2]
	if !exist {
		return
	}

	if tag1 != tag2 {
		if tag1.Name == "Enemy" && tag2.Name == "Bullet" {
			c.RemoveEntity(e1)
			c.RemoveEntity(e2)
		}
		if tag2.Name == "Enemy" && tag1.Name == "Bullet" {
			c.RemoveEntity(e1)
			c.RemoveEntity(e2)
		}
	}
}

func (c *Collider) RemoveEntity(entity component.Entity) {
	rl.UnloadTexture(c.Components.Renders[entity].Image)
	delete(c.Components.Positions, entity)
	delete(c.Components.Renders, entity)
	delete(c.Components.Velocities, entity)
	delete(c.Components.Tags, entity)
}
