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
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
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
	_, err := gamecontroller.CreateGame(configuration, "azerty")
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
	_, err := gamecontroller.CreateGame(configuration, "azerty")
	//Then
	if err == nil {
		t.Error("The size must be at least 3")
	}
}

func TestGetGameShouldRetrieveAnExistingGame(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
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

func TestGetFencePossibilitiesShouldRaiseAnExceptionWithAnUnknownGame(t *testing.T) {
	//Given
	setUp()
	//When
	_, err := gamecontroller.GetFencePossibilities("12453po")
	//Then
	if err == nil {
		t.Error("the game does not exists, an error should be raised")
	}
}

func TestGetFencePossibilitiesShouldRetrieveAllPossibilities(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	//When
	fences, _ := gamecontroller.GetFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 128 {
		t.Errorf("Without any fences, there are 128 possibilities but get %v", len(fences))
	}
}

func TestGetFencePossibilitiesShouldRetrieveAllPossibilitiesWithAFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	gamecontroller.JoinGame(newGame.ID, "qsdfgh")
	gamecontroller.AddFence(newGame.ID, game.Fence{game.Position{0, 0}, false}, "azerty")
	//When
	fences, _ := gamecontroller.GetFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 125 {
		t.Errorf("With one fences, there are 125 possibilities but get %v", len(fences))
	}
}

func TestGetFencePossibilitiesShouldRetrievePossibilitiesWithoutFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	//When
	fences, _ := gamecontroller.GetFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 8 {
		t.Errorf("With one fences, there are 4 possibilities but get %v", len(fences))
	}
}

func TestGetFencePossibilitiesShouldRetrieveAllPossibilitiesWithAFenceWihtoutClosingPath(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	gamecontroller.JoinGame(newGame.ID, "qsdfgh")
	gamecontroller.AddFence(newGame.ID, game.Fence{game.Position{0, 0}, false}, "azerty")
	//When
	fences, _ := gamecontroller.GetFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 4 {
		t.Errorf("With one fences, there are 4 possibilities but get %v", len(fences))
	}
}

func TestAddFenceNotPossibleWithoutAnOpponent(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	//When
	_, err := gamecontroller.AddFence(newGame.ID, game.Fence{game.Position{0, 0}, false}, "azerty")
	//Then
	if err == nil {
		t.Error("It is not possible to play without an opponent")
	}
	if err.Error() != "Game is not ready" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestMovePawnNotPossibleWithoutAnOpponent(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	//When
	_, err := gamecontroller.MovePawn(newGame.ID, game.Position{1, 2}, "azerty")
	//Then
	if err == nil {
		t.Error("It is not possible to play without an opponent")
	}
	if err.Error() != "Game is not ready" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestJoinGameShouldAddTheOpponent(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	//When
	err := gamecontroller.JoinGame(newGame.ID, "qsdfgh")
	//Then
	if err != nil {
		t.Errorf("It should be possible to join the game: %s", err.Error())
	}
}

func TestJoinGameShouldNotMoreThanExpectedOpponents(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	gamecontroller.JoinGame(newGame.ID, "qsdfgh")
	//When
	err := gamecontroller.JoinGame(newGame.ID, "wxcvbn")
	//Then
	if err == nil {
		t.Error("It is not possible to join more than expected")
	}
	if err.Error() != "Game is already set" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestAddFenceNotPossibleWithAnUnkownPlayer(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	gamecontroller.JoinGame(newGame.ID, "qsdfgh")
	//When
	_, err := gamecontroller.AddFence(newGame.ID, game.Fence{game.Position{0, 0}, false}, "spy")
	//Then
	if err == nil {
		t.Error("It is not possible to play with an unknown player")
	}
	if err.Error() != "Forbidden" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestMovePawnNotPossibleWithAnUnkownPlayer(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(configuration, "azerty")
	gamecontroller.JoinGame(newGame.ID, "qsdfgh")
	//When
	_, err := gamecontroller.MovePawn(newGame.ID, game.Position{1, 2}, "spy")
	//Then
	if err == nil {
		t.Error("It is not possible to play with an unknown player")
	}
	if err.Error() != "Forbidden" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}
