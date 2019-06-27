package game

import (
	"errors"
)

//Board game board 
type Board struct {
	BoardSize int
}

func NewBoard(boardSize int) (*Board, error) {
	if boardSize % 2 == 0 {
		return nil, errors.New("The board size must be an odd number")
	}
	if boardSize < 3 {
		return nil, errors.New("The board size must be at least 3");
	}
	return &Board{boardSize}, nil
}
