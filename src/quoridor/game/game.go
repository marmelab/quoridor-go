package game

import (
	"errors"

	"quoridor/storage"

	"github.com/lithammer/shortuuid"
)

// Game is the controller  
type Game struct {
	Id    string
	Pawn  Pawn  `json:"pawn"`
	Board *Board `json:"board"`
}

// CreateGame create a game with the default configuration
func CreateGame(conf *Configuration) (*Game, error) {
	boardSize := conf.BoardSize
	board, err := NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	lineCenter := (boardSize - 1) / 2
	pawn := Pawn{Position{0, lineCenter}}	
	id:= shortuuid.New()
	game := Game{id, pawn, board}
	storage.Set(game.Id, game)
	return &game, nil
}

// GetGame get the game via its identifier
func GetGame(id string) (Game, error) {
	game, found := storage.Get(id)
	if !found {
		return Game{}, errors.New("The game does not exist")
	}
	return game.(Game), nil
}
