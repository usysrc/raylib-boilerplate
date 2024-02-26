package component

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Position struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type Render struct {
	Image rl.Texture2D
	Z     float64
	Scale float64
}

type Tag struct {
	Name string
}
