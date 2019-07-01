package game

// Fence represents a fence on the board
type Fence struct {
	NWSquare Position `json:square`
	Horizontal bool `json:horizontal`
}

func (f Fence) Equals(other Fence) bool {
	return f.NWSquare.Equals(other.NWSquare) && f.Horizontal == other.Horizontal
}

type Fences []Fence

func (fences Fences) indexOf(e Fence) int {
	for index, a := range fences {
		if a.Equals(e) {
			return index
		}
	}
	return -1
}

//Contains check if at least one fence is inside the fences
func (fences Fences) Contains(fence1 Fence, fence2 Fence) bool {
    return fences.indexOf(fence1) != -1 || fences.indexOf(fence2) != -1
}

func CanMove(from Position, to Position, fences Fences) bool {
    direction := GetDirection(from, to)
    if (direction == UNKNOWN) {
        return false
    }
    var fence1, fence2 Fence
    switch direction {
    case EAST:
        fence1 = Fence{Position{from.Column, from.Row -1}, false}
        fence2 = Fence{Position{from.Column, from.Row}, false}
    case WEST:
        fence1 = Fence{Position{from.Column -1, from.Row -1}, false}
        fence2 = Fence{Position{from.Column -1, from.Row}, false}
    case NORTH:
        fence1 = Fence{Position{from.Column -1, from.Row -1}, true}
        fence2 = Fence{Position{from.Column, from.Row -1}, true}
    case SOUTH:
        fence1 = Fence{Position{from.Column -1, from.Row + 1}, true}
        fence2 = Fence{Position{from.Column, from.Row}, true}
    default:
        panic("Unknown direction")
    }
    return !fences.Contains(fence1, fence2)
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
