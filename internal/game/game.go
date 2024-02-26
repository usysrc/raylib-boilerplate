package game

import (
	"errors"

	"github.com/usysrc/raylib-boilerplate/internal/game/component"
	"github.com/usysrc/raylib-boilerplate/internal/game/entity"
	"github.com/usysrc/raylib-boilerplate/internal/game/systems"
)

var ErrTerminated = errors.New("errTerminated")

type Game struct {
	cm             *component.Manager
	movementSystem *systems.MovementSystem
	renderSystem   *systems.RenderSystem
	inputSystem    *systems.InputSystem
	spawnSystem    *systems.Spawn
	colliderSystem *systems.Collider
}

func (g *Game) Init() {
	// Create component manager
	g.cm = component.NewManager()

	// Create entities
	entity.CreateBackground(g.cm)
	ship := entity.CreateShip(g.cm)

	// Create systems
	g.movementSystem = &systems.MovementSystem{Components: g.cm}
	g.renderSystem = &systems.RenderSystem{Components: g.cm}
	g.inputSystem = &systems.InputSystem{Components: g.cm, ShipEntity: ship}
	g.spawnSystem = systems.NewSpawn(g.cm)
	g.colliderSystem = &systems.Collider{Components: g.cm}
}

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.movementSystem.Update()
	g.colliderSystem.Update()
	g.spawnSystem.Update()
	return nil
}

func (g *Game) Draw() {
	g.renderSystem.Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
