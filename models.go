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

// TODO: computed player-alive -> every-ship-hits === every-ship-length

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

type PlayerSetup byte
const (
	IsFirstPlayerReady PlayerSetup = iota
	IsSecondPlayerReady
)

type BattleshipState struct {
	FirstPlayer       Board
	SecondPlayer      Board
	IsFirstPlayerTurn bool
	Phase             Phase
}
