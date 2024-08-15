package battlesshiplib

type Ship struct {
	// computed alive -> hits === length
	Length      uint8
	Origin      Coordinate
	Orientation Orientation
	Hits        uint8
}

func (s Ship) ToCoords() []Coordinate {
	coords := make([]Coordinate, s.Length)
	for i := uint8(0); i < s.Length; i++ {
		switch s.Orientation {
		case Up:
			coords[i] = Coordinate{
				X: s.Origin.X,
				Y: s.Origin.Y - i,
			}
		case Down:
			coords[i] = Coordinate{
				X: s.Origin.X,
				Y: s.Origin.Y + i,
			}
		case Left:
			coords[i] = Coordinate{
				X: s.Origin.X - i,
				Y: s.Origin.Y,
			}
		case Right:
			coords[i] = Coordinate{
				X: s.Origin.X + i,
				Y: s.Origin.Y,
			}
		}
	}
	return coords
}
