package game

import (
	"testing"
	"quoridor/game"
)

func TestPathShouldFindAPathWithoutFences(t *testing.T) {
    //Given
    board, _ := game.NewBoard(3)
    fences := []game.Fence{}
    src := game.Position{0, 1}
    destinations := []game.Position{game.Position{2, 0}, game.Position{2, 1}, game.Position{2, 2}}
	//When
	d := game.Path(*board, fences, src, destinations)
	//Then
	if d == -1 {
		t.Error("A path exists without fences")
	}
}

func TestPathShouldFindAPathWithVerticalFences(t *testing.T) {
    //Given
    board, _ := game.NewBoard(3)
    fences := []game.Fence{
        game.Fence{game.Position{0, 0}, false},
        game.Fence{game.Position{1, 1}, false}}
    src := game.Position{0, 1}
    destinations := []game.Position{game.Position{2, 0}, game.Position{2, 1}, game.Position{2, 2}}
	//When
	d := game.Path(*board, fences, src, destinations)
	//Then
	if d == -1 {
		t.Error("A path still exists with vertical fences")
	}
}

func TestPathShouldNotFindAPathWithFences(t *testing.T) {
    //Given
    board, _ := game.NewBoard(3)
    fences := []game.Fence{
        game.Fence{game.Position{0, 0}, false},
        game.Fence{game.Position{0, 1}, true}}
    src := game.Position{0, 1}
    destinations := []game.Position{game.Position{2, 0}, game.Position{2, 1}, game.Position{2, 2}}
	//When
	d := game.Path(*board, fences, src, destinations)
	//Then
	if d != -1 {
		t.Error("No more path")
	}
}
