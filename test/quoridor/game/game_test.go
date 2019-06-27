package game

import (
	"testing"
	"quoridor/game"
)

func TestCreateGame(t *testing.T) {
	//Given
	configuration := game.Configuration{9}
	//When
	newGame := game.CreateGame(configuration)
	//Then
	if newGame.Id == 0 {
		t.Error("create a game should define an id")
	}
}
