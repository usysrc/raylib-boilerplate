package particle

import rl "github.com/gen2brain/raylib-go/raylib"

type Particle struct {
	Pos      rl.Vector2
	Dir      rl.Vector2
	Alive    bool
	lifetime float32
	life     float32
	speed    float32
}

var particles []Particle
var particleTexture rl.Texture2D

func Create(pos rl.Vector2, lifetime float32, dir rl.Vector2) {
	p := Particle{}
	p.lifetime = lifetime
	p.life = 0
	p.Alive = true
	p.Pos = pos
	p.Dir = dir
	p.speed = 500
	particles = append(particles, p)
}

func (p *Particle) Update() {
	p.life += rl.GetFrameTime()
	p.Pos.X += p.Dir.X * rl.GetFrameTime() * p.speed
	p.Pos.Y += p.Dir.Y * rl.GetFrameTime() * p.speed

	if p.life >= p.lifetime {
		p.Alive = false
	}
	if p.Pos.X > 800 || p.Pos.Y > 600 || p.Pos.X < 0 || p.Pos.Y < 0 {
		p.Alive = false
	}
}

func (p *Particle) Draw() {
	// fade out particle
	alpha := uint8(rl.Lerp(255, 0, p.life/p.lifetime))
	rl.DrawTexturePro(particleTexture, rl.Rectangle{X: 0, Y: 0, Width: float32(particleTexture.Width), Height: float32(particleTexture.Height)}, rl.Rectangle{X: p.Pos.X, Y: p.Pos.Y, Width: float32(particleTexture.Width), Height: float32(particleTexture.Height)}, rl.Vector2{X: float32(particleTexture.Width) / 2, Y: float32(particleTexture.Height) / 2}, 0, rl.Color{R: 255, G: 255, B: 255, A: alpha})
}

func Init() {
	particleTexture = rl.LoadTexture("internal/assets/particle.png")

	particles = make([]Particle, 0)
}

func Update() {
	for i := 0; i < len(particles); {
		particles[i].Update()
		if particles[i].Alive {
			i++
		} else {
			particles[i] = particles[len(particles)-1]
			particles = particles[:len(particles)-1]
		}
	}
}

func Draw() {
	for i := range particles {
		particles[i].Draw()
	}
}
