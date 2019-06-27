package game

import (
	"testing"
	"quoridor/game"
	"quoridor/storage"
)

func setUp() {
	storage.Init()
}

func TestCreateGame(t *testing.T) {
	//Given
	configuration := game.Configuration{9}
	//When
	newGame, _ := game.CreateGame(configuration)
	//Then
	if newGame.Id == "" {
		t.Error("create a game should define an id")
	}
}

func TestCreateGameShouldNotBePossibleWithEvenNumber(t *testing.T) {
	//Given
	configuration := game.Configuration{8}
	//When
	_, err := game.CreateGame(configuration)
	//Then
	if err == nil {
		t.Error("The size must be an odd number")
	}
}

func TestCreateGameShouldNotBePossibleWithLessThanThree(t *testing.T) {
	//Given
	configuration := game.Configuration{1}
	//When
	_, err := game.CreateGame(configuration)
	//Then
	if err == nil {
		t.Error("The size must be at least 3")
	}
}
