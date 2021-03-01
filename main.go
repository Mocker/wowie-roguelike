package main

import (
	_ "image/png"
    "log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/bytearena/ecs"
)




//IsoTo2D convert isometric x,y to grid x,y
func IsoTo2D(x int, y int) (int, int) {
	gridX := (2 * y + x) / 2
	gridY := (2 * y - x) / 2
	return gridX, gridY
}

//TwoDToIso convert x,y from grid to isometric
func TwoDToIso(x int, y int) (int, int) {
	isoX := x - y
	isoY := (x + y) / 2
	return isoX, isoY
}

// ** Main Game Class **

//Game main game data
type Game struct{
	Map GameMap
	Data GameData
	WindowWidth int
	WindowHeight int
	World *ecs.Manager
	WorldTags map[string]ecs.Tag
}


//NewGame create a new game object and initialize data
// TODO:: refactor NewGame constructor
func NewGame() *Game {
	g := &Game{
		WindowHeight: 1200,
		WindowWidth: 1600,
	}
	world, tags := InitializeWorld()
	g.WorldTags = tags
	g.World = world
	g.Map = NewGameMap()
	g.Data = NewGameData()
	return g
}

//Update is called each tic
func (g *Game) Update() error {
	TryMovePlayer(g)
	return nil
}

//Draw is called each draw cycle and is where we blit
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen, (g.WindowWidth/3), 0)
	// TODO:: walls should be renderable components with a higher depth level instead of an extra set of map tiles
	ProcessRenderables(g, level, screen, (g.WindowWidth/3), 0)
}

//Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) { return  g.WindowWidth, g.WindowHeight }

func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("My Wowie RL")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}