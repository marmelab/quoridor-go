package game

//Pawn defines a pawn on the board
type Pawn struct {
	Position Position `json:"position"`
	Goal Direction `json:"goal"`
}

type Pawns []Pawn

func (pawns Pawns) IndexOf(e Position) int {
	for index, a := range pawns {
		if a.Position.Equals(e) {
			return index
		}
	}
	return -1
}
