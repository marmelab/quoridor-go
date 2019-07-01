package game

import (
	"testing"
	"quoridor/game"
	"quoridor/service"
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
	newGame, _ := service.CreateGame(&configuration)
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
	_, err := service.CreateGame(&configuration)
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
	_, err := service.CreateGame(&configuration)
	//Then
	if err == nil {
		t.Error("The size must be at least 3")
	}
}

func TestGetGameShouldRetrieveAnExistingGame(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{9}
	newGame, _ := service.CreateGame(&configuration)
	//When
	getGame, _ := service.GetGame(newGame.ID)
	//Then
	if newGame.ID != getGame.ID {
		t.Error("Games are not the same")
	}
}

func TestGetGameShouldRaiseAnExceptionWithAnUnknownGame(t *testing.T) {
	//Given
	setUp()
	//When
	_, err := service.GetGame("12453po")
	//Then
	if err == nil {
		t.Error("the game does not exists, an error should be raised")
	}
}

func TestAddFenceShouldAddTheFenceAtTheRightPlace(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := service.CreateGame(&configuration)
	//When
	updatedGame, err := newGame.AddFence(game.Fence{game.Position{0, 0}, false})
	//Then
	if err != nil {
		t.Errorf("the game should exist: %s", err.Error())
		return
	}
	if len(updatedGame.Fences) != 1 {
		t.Error("the game should contain a new fence")
		return
	}
	actual := updatedGame.Fences[0]
	if !actual.Equals(game.Fence{game.Position{0, 0}, false}) {
		t.Error("the game does not contain the right fence")
		return
	}
}

func TestAddFenceShouldNotBePossibleOnVerticalFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := service.CreateGame(&configuration)
	updatedGame, _ := newGame.AddFence(game.Fence{game.Position{0, 0}, false})
	//When
	_, err := updatedGame.AddFence(game.Fence{game.Position{0, 0}, true})
	//Then
	if err == nil {
		t.Error("It is not possible to add another fence at the same place")
		return
	}
	if err.Error() != "The fence overlaps another one" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestAddFenceShouldNotBePossibleOnHorizontalFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := service.CreateGame(&configuration)
	updatedGame, _ := newGame.AddFence(game.Fence{game.Position{0, 0}, true})
	//When
	_, err := updatedGame.AddFence(game.Fence{game.Position{0, 0}, false})
	//Then
	if err == nil {
		t.Error("It is not possible to add another fence at the same place")
		return
	}
	if err.Error() != "The fence overlaps another one" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestAddFenceShouldNotBePossibleToAddAnHorizontalFenceOnSquareAfterFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{5}
	game1, _ := service.CreateGame(&configuration)
	game2, _ := game1.AddFence(game.Fence{game.Position{0, 0}, true})
	//When
	_, err := game2.AddFence(game.Fence{game.Position{1, 0}, true})
	//Then
	if err == nil {
		t.Error("It should not be possible to add an horizontal fence one square after a fence")
		return
	}
	if err.Error() != "The fence overlaps another one" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestAddFenceShouldNotBePossibleToAddAVerticalFenceOnSquareAfterFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{5}
	game1, _ := service.CreateGame(&configuration)
	game2, _ := game1.AddFence(game.Fence{game.Position{0, 0}, false})
	//When
	_, err := game2.AddFence(game.Fence{game.Position{0, 1}, false})
	//Then
	if err == nil {
		t.Error("It should not be possible to add a vertical fence one square after a fence")
		return
	}
	if err.Error() != "The fence overlaps another one" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestAddFenceShouldNotBePossibleToAddAVerticalFenceBeforeVerticalFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{5}
	game1, _ := service.CreateGame(&configuration)
	game2, _ := game1.AddFence(game.Fence{game.Position{0, 1}, false})
	//When
	_, err := game2.AddFence(game.Fence{game.Position{0, 0}, false})
	//Then
	if err == nil {
		t.Error("It should not be possible to add a vertical fence before a vertical fence")
		return
	}
	if err.Error() != "The fence overlaps another one" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestAddFenceShouldNotBePossibleToAddAnHorizontalFenceBeforeHorizontalFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{5}
	game1, _ := service.CreateGame(&configuration)
	game2, _ := game1.AddFence(game.Fence{game.Position{1, 0}, true})
	//When
	_, err := game2.AddFence(game.Fence{game.Position{0, 0}, true})
	//Then
	if err == nil {
		t.Error("It should not be possible to add an horizontal fence before a horizontal fence")
		return
	}
	if err.Error() != "The fence overlaps another one" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestAddFenceShoulBePossibleToAddAVerticalFenceBetweenTwoHorizontalFences(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{5}
	game1, _ := service.CreateGame(&configuration)
	game2, _ := game1.AddFence(game.Fence{game.Position{0, 0}, true})
	game3, _ := game2.AddFence(game.Fence{game.Position{2, 0}, true})
	//When
	game4, err := game3.AddFence(game.Fence{game.Position{1, 0}, false})
	//Then
	if err != nil {
		t.Errorf("Not the right error: %s", err.Error())
		return
	}
	if len(game4.Fences) != 3 {
		t.Errorf("Three fences should be in the board: %d", len(game4.Fences))
	}
}

func TestAddFenceShoulBePossibleToAddAnHorizontalFenceBetweenTwoVerticalFences(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{5}
	game1, _ := service.CreateGame(&configuration)
	game2, _ := game1.AddFence(game.Fence{game.Position{0, 0}, false})
	game3, _ := game2.AddFence(game.Fence{game.Position{0, 2}, false})
	//When
	game4, err := game3.AddFence(game.Fence{game.Position{0, 1}, true})
	//Then
	if err != nil {
		t.Errorf("Not the right error: %s", err.Error())
		return
	}
	if len(game4.Fences) != 3 {
		t.Errorf("Three fences should be in the board: %d", len(game4.Fences))
	}
}

func TestAddFenceShouldNotBePossibleToAddAFenceWhichClosesTheAccessToTheGoalLine(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	game1, _ := service.CreateGame(&configuration)
	game2, _ := game1.AddFence(game.Fence{game.Position{0, 0}, false})
	//When
	_, err := game2.AddFence(game.Fence{game.Position{0, 1}, true})
	//Then
	if err == nil {
		t.Error("It should not be possible to add a fence which closes the access to the goal line")
		return
	}
	if err.Error() != "No more access to goal line" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}
