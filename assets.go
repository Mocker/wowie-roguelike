package main

// loads images we will need into a map

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var atlas map[string]*Asset
var atlasLoaded bool

//Asset basic struct that can contain images or other assets
type Asset struct {
	Image *ebiten.Image
}

func init() {
	atlasLoaded = false
}

func loadAssets() {
	atlas :=  make(map[string]*Asset)
	var temp Asset
	// TODO:: this really should not be hardcoded but whatever
	floorTile, _, err := ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stone_E.png");
	if err != nil {
		log.Fatal(err)
	}
	temp = Asset{ Image: floorTile}
	atlas["floor"] = &temp
	var wall *ebiten.Image
	wall, _, err = ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_E.png");
	if err != nil {
		log.Fatal(err)
	}
	temp = Asset{ Image: wall}
	atlas["wallE"] = &temp
	wall, _, err = ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_W.png");
	if err != nil {
		log.Fatal(err)
	}
	temp = Asset{ Image: wall}
	atlas["wallW"] = &temp
	wall, _, err = ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_N.png");
	if err != nil {
		log.Fatal(err)
	}
	temp = Asset{ Image: wall}
	atlas["wallN"] = &temp
	wall, _, err = ebitenutil.NewImageFromFile("assets/kenney-dungeon/Isometric/stoneWall_S.png");
	if err != nil {
		log.Fatal(err)
	}
	temp = Asset{ Image: wall}
	atlas["wallS"] = &temp

}

//GetAssets load a hardcoded list of assets we will need elsewhere in the game
func GetAssets() map[string]*Asset {
	if atlasLoaded == false {
		loadAssets()
		atlasLoaded = true
	}
	return atlas
}