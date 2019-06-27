package game

import (
	"errors"

	"quoridor/storage"

	"github.com/lithammer/shortuuid"
)

type Game struct {
	Id    string
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
	id:= shortuuid.New()
	game := Game{id, pawn, board}
	storage.Set(game.Id, game)
	return &game, nil
}
