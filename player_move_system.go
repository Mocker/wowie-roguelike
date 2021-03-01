package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

//TryMovePlayer ...
func TryMovePlayer(g *Game) {
	players := g.WorldTags["players"]
	x := 0
	y := 0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	
	

	if x != 0 || y != 0 {
		level := g.Map.CurrentLevel
		for _, result := range g.World.Query(players) {
			pos := result.Components[position].(*Position)
			if 0 > pos.X+x || g.Data.ScreenWidth <= pos.X+x {
				x = 0
			}
			if 0 > pos.Y+y || g.Data.ScreenHeight <= pos.Y+y {
				y = 0
			}
			index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
			tile := level.Tiles[index]

			if tile.Blocked != true {
				pos.X += x 
				pos.Y += y
			}
		}
	}
	
}