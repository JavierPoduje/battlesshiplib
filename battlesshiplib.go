package battlesshiplib

const BOARD_WIDTH = 10
const BOARD_HEIGHT = 10

func NewBattleshipGame() BattleshipState {
	return BattleshipState{
		FirstPlayer: Board{
			Width:  BOARD_WIDTH,
			Height: BOARD_HEIGHT,
			Shots:  make([]Coordinate, 0),
			Ships:  NewShips(),
		},
		SecondPlayer: Board{
			Width:  BOARD_WIDTH,
			Height: BOARD_HEIGHT,
			Shots:  make([]Coordinate, 0),
			Ships:  NewShips(),
		},
		IsFirstPlayerTurn: true,
		Phase:             Setup,
	}
}

func NewShips() []Ship {
	return []Ship{
		{
			Length:      2,
			Orientation: Down,
			Origin: Coordinate{
				X: 0,
				Y: 0,
			},
		},
		{
			Length:      3,
			Orientation: Right,
			Origin: Coordinate{
				X: 2,
				Y: 2,
			},
		},
		{
			Length:      3,
			Orientation: Right,
			Origin: Coordinate{
				X: 2,
				Y: 4,
			},
		},
		{
			Length:      4,
			Orientation: Left,
			Origin: Coordinate{
				X: 9,
				Y: 9,
			},
		},
		{
			Length:      5,
			Orientation: Up,
			Origin: Coordinate{
				X: 7,
				Y: 5,
			},
		},
	}
}
