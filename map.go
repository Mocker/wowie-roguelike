package main

//GameMap holds all level and world data
type GameMap struct {
	Dungeons []Dungeon
	CurrentLevel Level
}

//NewGameMap create a new set of maps for the game
func NewGameMap() GameMap {
	// TODO:: create multiple levels and dungeons. currently just set to a single level
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)
	d := Dungeon{
		Name: "My first dungeon",
		Levels: levels,
	}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	gm := GameMap{
		Dungeons: dungeons,
		CurrentLevel: l,
	}
	return gm
}