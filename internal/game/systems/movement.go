package systems

import "github.com/usysrc/raylib-boilerplate/internal/game/component"

type MovementSystem struct {
	// References to the relevant component managers or storages
	Components *component.Manager
}

func (m *MovementSystem) Update() {
	for e, vel := range m.Components.Velocities {
		if pos, ok := m.Components.Positions[e]; ok {
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
