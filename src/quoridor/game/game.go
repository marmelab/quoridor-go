package game

import (
	"math/rand"

	"quoridor/storage"
)

type Game struct {
	Id    int
	Pawn  Pawn  `json:"pawn"`
	Board Board `json:"-"`
}

// CreateGame create a game with the default configuration
func CreateGame(conf Configuration) Game {
	lineCenter := (conf.BoardSize - 1) / 2
	pawn := Pawn{Position{0, lineCenter}}
	board := Board{conf.BoardSize}
	game := Game{rand.Int(), pawn, board}
	storage.Set(game.Id, game)
	return game
}
