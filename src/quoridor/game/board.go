package game

import (
	"errors"
)

//Board game board 
type Board struct {
	BoardSize int `json:"board-size"`
	Squares []Position `json:"squares"`
}

func NewBoard(boardSize int) (*Board, error) {
	if boardSize % 2 == 0 {
		return nil, errors.New("The board size must be an odd number")
	}
	if boardSize < 3 {
		return nil, errors.New("The board size must be at least 3")
	}
	squares := []Position{}
	for row := 0; row < boardSize; row++ {
		for column := 0; column < boardSize; column++ {
			squares = append(squares, Position{column, row})
		}
	}
	return &Board{boardSize, squares}, nil
}
