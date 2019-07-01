package game

import (
	"testing"
	"quoridor/game"
)

func TestNewBoard(t *testing.T) {
	//Given
	//When
	_, err := game.NewBoard(9)
	//Then
	if err != nil {
		t.Errorf("create a board should not raise an exception: %v", err.Error())
	}
}

func TestNewBoardShouldNotBePossibleWithEvenNumber(t *testing.T) {
	//Given
	//When
	_, err := game.NewBoard(8)
	//Then
	if err == nil {
		t.Error("The size must be an odd number")
	}
}

func TestNewBoardShouldNotBePossibleWithLessThanThree(t *testing.T) {
	//Given
	//When
	_, err := game.NewBoard(1)
	//Then
	if err == nil {
		t.Error("The size must be at least 3")
	}
}

func TestIsInBoardShouldReturnTrueIfPositionIsInside(t *testing.T) {
	//Given
	board, _ := game.NewBoard(9)
	//When
	inside := board.IsInBoard(game.Position{0, 1})
	//Then
	if inside == false {
		t.Error("The position should be inside the board")
	}
}

func TestIsInBoardShouldReturnFalseIfPositionIsOutside(t *testing.T) {
	//Given
	board, _ := game.NewBoard(9)
	//When
	inside := board.IsInBoard(game.Position{-1, 1})
	//Then
	if inside == true {
		t.Error("The position should not be inside the board")
	}
}
