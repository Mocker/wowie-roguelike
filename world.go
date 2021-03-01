package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// The world is a collection of all entities in the game
// using the entity component system to manage them

var position *ecs.Component
var renderable *ecs.Component

//GetPositionComponent ..
func GetPositionComponent() *ecs.Component {
	return position
}

//GetRenderableComponent ..
func GetRenderableComponent() *ecs.Component {
	return renderable
}

//InitializeWorld ...
func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()

	// TODO:: remove sample adding a component to our world and put it somewhere smarter
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/kenney-dungeon/Characters/Male/Male_0_Idle0.png")
	if err != nil {
		log.Fatal(err)
	}
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: 2,
			Y: 2,
		})
	players := ecs.BuildTag(player, position)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	return manager, tags
}
