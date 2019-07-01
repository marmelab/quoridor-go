package game

import (
	"testing"
	"quoridor/game"
)

func TestGetDirectionNorth(t *testing.T) {
	//Given
    //When
    direction := game.GetDirection(game.Position{1, 1}, game.Position{1, 0})
	//Then
	if direction != game.NORTH {
		t.Errorf("get direction should return north: %v", direction)
	}
}

func TestGetDirectionEast(t *testing.T) {
	//Given
    //When
    direction := game.GetDirection(game.Position{1, 1}, game.Position{2, 1})
	//Then
	if direction != game.EAST {
		t.Errorf("get direction should return east: %v", direction)
	}
}

func TestGetDirectionSouth(t *testing.T) {
	//Given
    //When
    direction := game.GetDirection(game.Position{1, 1}, game.Position{1, 2})
	//Then
	if direction != game.SOUTH {
		t.Errorf("get direction should return south: %v", direction)
	}
}

func TestGetDirectionWest(t *testing.T) {
	//Given
    //When
    direction := game.GetDirection(game.Position{1, 1}, game.Position{0, 1})
	//Then
	if direction != game.WEST {
		t.Errorf("get direction should return west: %v", direction)
	}
}

func TestGetDirectionUnknown(t *testing.T) {
	//Given
    //When
    direction := game.GetDirection(game.Position{1, 1}, game.Position{3, 1})
	//Then
	if direction != game.UNKNOWN {
		t.Errorf("get direction should return unknown: %v", direction)
	}
}
