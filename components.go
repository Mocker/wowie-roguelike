package main

import (
	
	"github.com/hajimehoshi/ebiten/v2"
)

//Player basic player data
type Player struct{}

//Position x,y
type Position struct {
	X int
	Y int
}

//Renderable something you can render to the screen
type Renderable struct{
	Image *ebiten.Image
}

//Movable something that can be moved
type Movable struct{}