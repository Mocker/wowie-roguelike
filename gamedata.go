package main

//GameData ...
type GameData struct {
	ScreenWidth int
	ScreenHeight int
	TileWidth int
	TileHeight int
}

//NewGameData ...
func NewGameData() GameData {
	g := GameData{
		ScreenWidth: 10,
		ScreenHeight: 10,
		TileWidth: 64,
		TileHeight: 64,
	}
	return g
}