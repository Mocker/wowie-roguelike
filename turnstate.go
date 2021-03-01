package main

//TurnState ..
type TurnState int

const (
	//BeforePlayerAction ...
	BeforePlayerAction = iota
	//PlayerTurn ...
	PlayerTurn
	//MonsterTurn ...
	MonsterTurn
)

//GetNextState ...
func GetNextState(state TurnState) TurnState {
	switch state {
		case BeforePlayerAction:
			return PlayerTurn
		case PlayerTurn:
			return MonsterTurn
		case MonsterTurn:
			return BeforePlayerAction
		default:
			return PlayerTurn
	}
}