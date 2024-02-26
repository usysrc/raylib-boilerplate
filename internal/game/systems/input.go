package systems

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/component"
	"github.com/usysrc/raylib-boilerplate/internal/game/entity"
)

type InputSystem struct {
	Components *component.Manager
	ShipEntity component.Entity
}

func (is *InputSystem) Update() {
	velocity, exists := is.Components.Velocities[is.ShipEntity]
	if !exists {
		log.Fatal("entity ship does not have velocity component")
		return
	}

	speed := 16.0
	if rl.IsKeyDown(rl.KeyUp) {
		velocity.Y -= 1.0 * float64(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		velocity.Y += 1.0 * float64(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		velocity.X -= 1.0 * float64(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyDown(rl.KeyRight) {
		velocity.X += 1.0 * float64(rl.GetFrameTime()) * speed
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		is.CreateBullet()
	}

	// apply friction
	velocity.X *= 0.95
	velocity.Y *= 0.95
}

func (is *InputSystem) CreateBullet() {
	bullet := entity.CreateBullet(is.Components)
	bulletPos, exists := is.Components.Positions[bullet]
	if !exists {
		log.Fatal("entity does not have position component")
	}
	shipPos, exists := is.Components.Positions[is.ShipEntity]
	if !exists {
		log.Fatal("entity does not have position component")
	}
	bulletPos.X = shipPos.X
	bulletPos.Y = shipPos.Y

	bulletVelo, exists := is.Components.Velocities[bullet]
	if !exists {
		log.Fatal("entity does not have velocity component")
	}
	bulletVelo.Y = -10
}
