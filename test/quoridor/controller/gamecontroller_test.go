package gamecontroller

import (
	"testing"
	"quoridor/controller"
	"quoridor/game"
	"quoridor/storage"
)

func setUp() {
	storage.Init()
}

func TestCreateGame(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	//When
	newGame, _ := gamecontroller.CreateGame(&configuration)
	//Then
	if newGame.ID == "" {
		t.Error("create a game should define an id")
	}
}

func TestCreateGameShouldNotBePossibleWithEvenNumber(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{8}
	//When
	_, err := gamecontroller.CreateGame(&configuration)
	//Then
	if err == nil {
		t.Error("The size must be an odd number")
	}
}

func TestCreateGameShouldNotBePossibleWithLessThanThree(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{1}
	//When
	_, err := gamecontroller.CreateGame(&configuration)
	//Then
	if err == nil {
		t.Error("The size must be at least 3")
	}
}

func TestGetGameShouldRetrieveAnExistingGame(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(&configuration)
	//When
	getGame, _ := gamecontroller.GetGame(newGame.ID)
	//Then
	if newGame.ID != getGame.ID {
		t.Error("Games are not the same")
	}
}

func TestGetGameShouldRaiseAnExceptionWithAnUnknownGame(t *testing.T) {
	//Given
	setUp()
	//When
	_, err := gamecontroller.GetGame("12453po")
	//Then
	if err == nil {
		t.Error("the game does not exists, an error should be raised")
	}
}
