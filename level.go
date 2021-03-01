package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//MapTile ..
// TODO:: move map to its own class
// TODO:: map tile should have multiple layers with their own tags
type MapTile struct {
	PixelX int
	PixelY int
	Blocked bool
	Image *ebiten.Image
	Label string
}

//Level a single level of a dungeon
type Level struct {
	Tiles []MapTile
	Walls []MapTile
}

var assets map[string]*Asset

//NewLevel creates a blank game level
func NewLevel() Level {
	l := Level{}
	assets = GetAssets()
	tiles, walls := l.CreateTiles()
	l.Tiles = tiles
	l.Walls = walls
	return l
}

//GetIndexFromXY get single slice index for an x,y position
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}


//CreateTiles generate our map tiles. returns floors, walls
func (level *Level) CreateTiles() ([]MapTile, []MapTile) {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)
	walls := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)

	// TODO:: this should use the asset atlas .. if it worked
	floorTile, _, err := ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stone_E.png");
	if err != nil {
		log.Fatal(err)
	}
	wallE, _, err := ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_E.png");
	if err != nil {
		log.Fatal(err)
	}
	wallW, _, err := ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_W.png");
	if err != nil {
		log.Fatal(err)
	}
	wallN, _, err := ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_N.png");
	if err != nil {
		log.Fatal(err)
	}
	wallS, _, err := ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_S.png");
	if err != nil {
		log.Fatal(err)
	}

	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			// TODO:: more tiles and isometric facing walls and stuff
			tile := MapTile{
				PixelX: x * gd.TileWidth,
				PixelY: y * gd.TileWidth, // we are ignoring tile height here
				Blocked: false,
				Image: floorTile,
				Label: "floor",
			}
			
			// TODO:: place walls via mapgen instead of just on borders
			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight - 1 {
				tile.Blocked = true
				wall := MapTile{
					PixelX: x * gd.TileWidth,
					PixelY: y * gd.TileWidth, // we are ignoring tile height here
					Blocked: false,
					Label: "wall",
				}
				wall.Blocked = true
				if x == 0 {
					wall.Image = wallE
				} else if x == gd.ScreenWidth - 1 {
					wall.Image = wallW
				} else if y == 0 {
					wall.Image = wallS
				} else {
					wall.Image =wallN
				}
				walls[index] = wall
			}
			tiles[index] = tile
			
		}
	}
	return tiles, walls
}


//DrawLevel is called each draw cycle and is where we blit
func (level *Level) DrawLevel(screen *ebiten.Image, offsetX int, offsetY int) {
	// TODO:: fix isometric rendering..
	gd := NewGameData()
	// first layer is floor
	for x:= 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			isoX, isoY := TwoDToIso(tile.PixelX, tile.PixelY)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(isoX+(offsetX)), float64(isoY-(offsetY)))
			screen.DrawImage(tile.Image, op)
		}
	}
	// loop through walls (and other stuff.. ? )
	for x:= 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			wall := level.Walls[level.GetIndexFromXY(x, y)]
			if wall.Image != nil {
				isoX, isoY := TwoDToIso(wall.PixelX+1, wall.PixelY+1)
				wallYOffset := 0
				if y < 1 {
					wallYOffset = wallYOffset * -1
				}
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(isoX+(offsetX)), float64(isoY-(offsetY)+(wallYOffset)))
				screen.DrawImage(wall.Image, op)
			}
		}
	}
}