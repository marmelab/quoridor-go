package game

type Game struct {
	Pawn  Pawn  `json:"pawn"`
	Board Board `json:"-"`
}

// CreateGame create a game with the default configuration
func CreateGame(conf Configuration) Game {
	lineCenter := (conf.BoardSize - 1) / 2
	pawn := Pawn{Position{0, lineCenter}}
	board := Board{conf.BoardSize}
	return Game{pawn, board}
}
