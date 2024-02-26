package systems

import (
	"math/rand"

	"github.com/usysrc/raylib-boilerplate/internal/game/component"
	"github.com/usysrc/raylib-boilerplate/internal/game/entity"
)

const spawnInterval float64 = 1.35

type Spawn struct {
	Components *component.Manager
	timer      Timer
}

func NewSpawn(components *component.Manager) *Spawn {
	timer := Timer{}
	timer.Init()
	s := &Spawn{Components: components, timer: timer}
	s.Cycle()
	return s
}

func (s *Spawn) Cycle() {
	s.timer.After(rand.Float64()*spawnInterval, func() {
		entity.CreateEnemy(s.Components)
		s.Cycle()
	})
}

func (s *Spawn) Update() {
	s.timer.Update()
}
