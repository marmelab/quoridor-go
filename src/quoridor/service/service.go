package service

import (
	"errors"

	"quoridor/game"
	"quoridor/storage"

	"github.com/lithammer/shortuuid"
)

// CreateGame create a game with the default configuration
func CreateGame(conf *game.Configuration) (*game.Game, error) {
	boardSize := conf.BoardSize
	board, err := game.NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	lineCenter := (boardSize - 1) / 2
	pawn := game.Pawn{game.Position{0, lineCenter}}	
	id:= shortuuid.New()
	game := game.Game{id, pawn, []game.Fence{}, board}
	storage.Set(game.ID, game)
	return &game, nil
}

// GetGame get the game via its identifier
func GetGame(id string) (game.Game, error) {
	g, found := storage.Get(id)
	if !found {
		return game.Game{}, errors.New("The game does not exist")
	}
	return g.(game.Game), nil
}

//AddFence add the fence on the board
func AddFence(id string, fence game.Fence) (game.Game, error) {
	g, err := GetGame(id)
	if err != nil {
		return game.Game{}, err
	}
	g, errFence := g.AddFence(fence)
	if errFence != nil {
		return game.Game{}, errFence
	}
	storage.Set(g.ID, g)
	return g, nil
}
