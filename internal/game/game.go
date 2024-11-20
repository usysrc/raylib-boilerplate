package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Gamestater interface {
	Init()
	Update(g *Game) error
	Draw()
}

type Game struct {
	shouldCloseFlag bool
	state           Gamestater
}

func (g *Game) Init() {
	g.Switch("loading")
	g.shouldCloseFlag = false
}

func (g *Game) Update() error {
	g.state.Update(g)
	if rl.IsKeyPressed(rl.KeyEscape) {
		g.Close()
	}
	return nil
}

func (g *Game) Draw() {
	g.state.Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func (g *Game) ShouldClose() bool {
	return g.shouldCloseFlag
}

func (g *Game) Close() {
	g.shouldCloseFlag = true
}

func (g *Game) Switch(to string) {
	switch to {
	case "play":
		g.state = &Play{}
	case "death":
		g.state = &Death{}
	case "loading":
		g.state = &Loading{}
	}
	g.state.Init()
}
