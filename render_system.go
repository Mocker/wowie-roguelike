package main

import (
	//"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

//ProcessRenderables loop through renderable components and.. render them?
func ProcessRenderables(g *Game, level Level, screen *ebiten.Image, offsetX int, offsetY int) {
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		if result.Components[position] == nil || result.Components[renderable] == nil {
			continue;
		}
		pos := result.Components[position].(*Position)
		img := result.Components[renderable].(Renderable).Image

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		isoX, isoY := TwoDToIso(tile.PixelX, tile.PixelY)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(isoX+offsetX), float64(isoY-offsetY))
		screen.DrawImage(img, op)
	}
}