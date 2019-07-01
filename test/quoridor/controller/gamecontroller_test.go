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

func TestAddFencePossibilitiesShouldRaiseAnExceptionWithAnUnknownGame(t *testing.T) {
	//Given
	setUp()
	//When
	_, err := gamecontroller.AddFencePossibilities("12453po")
	//Then
	if err == nil {
		t.Error("the game does not exists, an error should be raised")
	}
}

func TestAddFencePossibilitiesShouldRetrieveAllPossibilities(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(&configuration)
	//When
	fences, _ := gamecontroller.AddFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 128 {
		t.Errorf("Without any fences, there are 128 possibilities but get %v", len(fences))
	}
}

func TestAddFencePossibilitiesShouldRetrieveAllPossibilitiesWithAFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := gamecontroller.CreateGame(&configuration)
	gamecontroller.AddFence(newGame.ID, game.Fence{game.Position{0, 0}, false})
	//When
	fences, _ := gamecontroller.AddFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 125 {
		t.Errorf("With one fences, there are 125 possibilities but get %v", len(fences))
	}
}

func TestAddFencePossibilitiesShouldRetrievePossibilitiesWithoutFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := gamecontroller.CreateGame(&configuration)
	//When
	fences, _ := gamecontroller.AddFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 8 {
		t.Errorf("With one fences, there are 4 possibilities but get %v", len(fences))
	}
}

func TestAddFencePossibilitiesShouldRetrieveAllPossibilitiesWithAFenceWihtoutClosingPath(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := gamecontroller.CreateGame(&configuration)
	gamecontroller.AddFence(newGame.ID, game.Fence{game.Position{0, 0}, false})
	//When
	fences, _ := gamecontroller.AddFencePossibilities(newGame.ID)
	//Then
	if len(fences) != 4 {
		t.Errorf("With one fences, there are 4 possibilities but get %v", len(fences))
	}
}
