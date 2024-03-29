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

func (fences Fences) IndexOf(e Fence) int {
	for index, a := range fences {
		if a.Equals(e) {
			return index
		}
	}
	return -1
}

func (fences Fences) RemoveFence(e Fence) Fences {
	index := fences.IndexOf(e)
	if (index > -1) {
		return fences.Remove(index)
	}
	return fences
}

func (fences Fences) Remove(index int) Fences {
	return append(fences[:index], fences[index+1:]...)
}

//Contains check if at least one fence is inside the fences
func (fences Fences) Contains(fence1 Fence, fence2 Fence) bool {
    return fences.IndexOf(fence1) != -1 || fences.IndexOf(fence2) != -1
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
