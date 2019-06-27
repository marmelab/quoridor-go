package game

import (
	"errors"
	"math/rand"

	"quoridor/storage"
)

type Game struct {
	Id    int
	Pawn  Pawn  `json:"pawn"`
	Board Board `json:"-"`
}

// CreateGame create a game with the default configuration
func CreateGame(conf Configuration) (*Game, error) {
	if conf.BoardSize % 2 == 0 {
		return nil, errors.New("The board size must be an odd number")
	} else if conf.BoardSize < 3 {
		return nil, errors.New("The board size must be at least 3");
	}
	lineCenter := (conf.BoardSize - 1) / 2
	pawn := Pawn{Position{0, lineCenter}}
	board := Board{conf.BoardSize}
	game := Game{rand.Int(), pawn, board}
	storage.Set(game.Id, game)
	return &game, nil
}
