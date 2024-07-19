package battlesshiplib

type Orientation byte

const (
	Up Orientation = iota
	Right
	Down
	Left
)

type Coordinate struct {
	X uint8
	Y uint8
}

type Ship struct {
	// computed alive -> hits === length
	Length      uint8
	Origin      Coordinate
	Orientation Orientation
	Hits        uint8
}

type Board struct {
	Ships  []Ship
	Shots  []Coordinate
	Width  uint8
	Height uint8
}

type Phase byte

const (
	Setup Phase = iota
	Playing
	Finished
)

type BattleshipState struct {
	FirstPlayer       Board
	SecondPlayer      Board
	IsFirstPlayerTurn bool
	Phase             Phase
}
