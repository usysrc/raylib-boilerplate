package ship

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/usysrc/raylib-boilerplate/internal/game/bullet"
	"github.com/usysrc/raylib-boilerplate/internal/game/enemy"
)

type Ship struct {
	Image       rl.Texture2D
	Pos         rl.Vector2
	Speed       float32
	Alive       bool
	Velocity    rl.Vector2
	MaxVelocity float32
	Sound       rl.Sound
}

func NewShip() *Ship {
	ship := &Ship{
		Alive:       false,
		Velocity:    rl.Vector2{},
		MaxVelocity: 10,
		Speed:       10,
	}
	img := rl.LoadImage("internal/assets/ship.png")
	ship.Image = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	ship.Pos = rl.Vector2{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight()) - 100}
	ship.Sound = rl.LoadSound("internal/assets/laser.wav")
	return ship
}

type GamestateSwitcher interface {
	Switch(to string)
}

func (s *Ship) Update(g GamestateSwitcher) {
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		if s.Velocity.Y > 0 {
			s.Velocity.Y = 0
		}
		s.Velocity.Y -= 1.0 * float32(rl.GetFrameTime()) * s.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		if s.Velocity.Y < 0 {
			s.Velocity.Y = 0
		}
		s.Velocity.Y += 1.0 * float32(rl.GetFrameTime()) * s.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		if s.Velocity.X > 0 {
			s.Velocity.X = 0
		}
		s.Velocity.X -= 1.0 * float32(rl.GetFrameTime()) * s.Speed
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		if s.Velocity.X < 0 {
			s.Velocity.X = 0
		}
		s.Velocity.X += 1.0 * float32(rl.GetFrameTime()) * s.Speed
	}
	if !rl.IsKeyDown(rl.KeyUp) && !rl.IsKeyDown(rl.KeyW) && !rl.IsKeyDown(rl.KeyDown) && !rl.IsKeyDown(rl.KeyS) {
		s.Velocity.Y *= 0.95
	}
	if !rl.IsKeyDown(rl.KeyLeft) && !rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyRight) && !rl.IsKeyDown(rl.KeyD) {
		s.Velocity.X *= 0.95
	}
	if s.Velocity.X > s.MaxVelocity {
		s.Velocity.X = s.MaxVelocity
	} else if s.Velocity.X < -s.MaxVelocity {
		s.Velocity.X = -s.MaxVelocity
	}
	if s.Velocity.Y > s.MaxVelocity {
		s.Velocity.Y = s.MaxVelocity
	} else if s.Velocity.Y < -s.MaxVelocity {
		s.Velocity.Y = -s.MaxVelocity
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		rl.PlaySound(s.Sound)
		bullet.Create(s.Pos.X, s.Pos.Y)
	}
	targetPos := rl.Vector2{X: s.Pos.X + s.Velocity.X, Y: s.Pos.Y + s.Velocity.Y}
	if targetPos.X > 0 && targetPos.X < float32(rl.GetScreenWidth()) {
		s.Pos.X += s.Velocity.X
	} else {
		s.Velocity.X = 0
	}
	if targetPos.Y > 0 && targetPos.Y < float32(rl.GetScreenHeight()) {
		s.Pos.Y += s.Velocity.Y
	} else {
		s.Velocity.Y = 0
	}
	enemies := enemy.GetEnemies()
	for i := range enemies {
		if rl.CheckCollisionCircles(s.Pos, 32, enemies[i].Pos, 20) {
			g.Switch("death")
			return
		}
	}
}

func (s *Ship) Draw() {
	// draw a line from the ship to the top of the screen
	rl.SetLineWidth(20)
	rl.DrawLine(int32(s.Pos.X), int32(s.Pos.Y), int32(s.Pos.X), 0, rl.Color{R: 255, G: 0, B: 0, A: 100})
	// draw image with the shipPos being at the center of the image
	rl.DrawTexturePro(s.Image, rl.Rectangle{X: 0, Y: 0, Width: float32(s.Image.Width), Height: float32(s.Image.Height)}, rl.Rectangle{X: s.Pos.X, Y: s.Pos.Y, Width: float32(s.Image.Width), Height: float32(s.Image.Height)}, rl.Vector2{X: float32(s.Image.Width) / 2, Y: float32(s.Image.Height) / 2}, 0, rl.White)
}

func (s *Ship) GetPos() rl.Vector2 {
	return s.Pos
}
