package game

//Pawn defines a pawn on the board
type Pawn struct {
	Position Position `json:"position"`
	Goal Direction `json:"goal"`
}
