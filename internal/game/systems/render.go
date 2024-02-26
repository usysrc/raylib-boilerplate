package systems

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/usysrc/raylib-boilerplate/internal/game/component"
)

type RenderSystem struct {
	Components *component.Manager
}

type RenderData struct {
	render *component.Render
	entity component.Entity
}

func (r *RenderSystem) Draw() {
	var renderData []RenderData
	for e, render := range r.Components.Renders {
		renderData = append(renderData, RenderData{
			entity: e,
			render: render,
		})
	}

	sort.Slice(renderData, func(i, j int) bool {
		return renderData[i].render.Z < renderData[j].render.Z
	})

	for _, data := range renderData {
		if pos, ok := r.Components.Positions[data.entity]; ok {
			// Draw the entity's image at its position
			//options := &raylib.DrawImageOptions{}
			//options.GeoM.Reset()
			//options.GeoM.Scale(data.render.Scale, data.render.Scale)
			//options.GeoM.Translate(pos.X, pos.Y)

			//rl.DrawTexture(data.render.Image, int32(pos.X), int32(pos.Y), rl.White)
			rl.DrawTextureEx(data.render.Image, rl.Vector2{X: float32(pos.X), Y: float32(pos.Y)}, 0, float32(data.render.Scale), rl.White)
		}
	}
}
