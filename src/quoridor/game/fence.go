package game

// Fence represents a fence on the board
type Fence struct {
	NWSquare Position `json:square`
	Horizontal bool `json:horizontal`
}

func (f Fence) Equals(other Fence) bool {
	return f.NWSquare.Equals(other.NWSquare) && f.Horizontal == other.Horizontal
}

type PositionSquare struct {
	NorthPosition Position
	EastPosition Position
	SouthPosition Position
	WestPosition Position   
}

func NewPositionSquare(center Position) PositionSquare {
	northPosition := center.Copy(0, -1)
	eastPosition := center.Copy(1, 0)
	southPosition := center.Copy(0, 1)
	westPosition := center.Copy(-1, 0)
	return PositionSquare{northPosition, eastPosition, southPosition, westPosition}
}
