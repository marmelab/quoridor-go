package game

import (
	"quoridor/game"
	"quoridor/storage"
	"testing"
)

func setUp() {
	storage.Init()
}

func checkMoves(t *testing.T, moves game.Positions, expected game.Positions) {
	if len(moves) != len(expected) {
		t.Errorf("Pawn can move in %v squares", len(moves))
		return
	}
	for _, pos := range expected {
		if moves.IndexOf(pos) == -1 {
			t.Errorf("Moves should contain %v but not present in %v", pos, moves)
			return
		}
	}
}

func TestAddFenceShouldAddTheFenceAtTheRightPlace(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	newGame, _ := game.NewGame(configuration)
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
	newGame, _ := game.NewGame(configuration)
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
	newGame, _ := game.NewGame(configuration)
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
	game1, _ := game.NewGame(configuration)
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
	game1, _ := game.NewGame(configuration)
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
	game1, _ := game.NewGame(configuration)
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
	game1, _ := game.NewGame(configuration)
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
	game1, _ := game.NewGame(configuration)
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
	game1, _ := game.NewGame(configuration)
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
	game1, _ := game.NewGame(configuration)
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

func TestMovePawnEast(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	//When
	g1, _ := g.MovePawn(game.Position{1, 1})
	//Then
	if !g1.Pawns[0].Position.Equals(game.Position{1, 1}) {
		t.Error("The pawn should move east")
	}
}

func TestMovePawnNorth(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	//When
	g1, _ := g.MovePawn(game.Position{0, 0})
	//Then
	if !g1.Pawns[0].Position.Equals(game.Position{0, 0}) {
		t.Error("The pawn should move north")
	}
}

func TestMovePawnSouth(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	//When
	g1, _ := g.MovePawn(game.Position{0, 2})
	//Then
	if !g1.Pawns[0].Position.Equals(game.Position{0, 2}) {
		t.Error("The pawn should move south")
	}
}

func TestMovePawnWest(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{0, 0}) // Move Pawn 1
	//When
	g2, _ := g1.MovePawn(game.Position{1, 1}) // Move Pawn 2
	//Then
	if !g2.Pawns[1].Position.Equals(game.Position{1, 1}) {
		t.Error("The pawn should move west")
	}
}

func TestMovePawnOutOfBoard(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	//When
	_, err := g.MovePawn(game.Position{-1, 1})
	//Then
	if err == nil {
		t.Error("It is not possible to move outside of the board")
		return
	}
	if err.Error() != "The new position is not inside the board" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestMovePawnToUnreachablePosition(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	//When
	_, err := g.MovePawn(game.Position{2, 2})
	//Then
	if err == nil {
		t.Error("It is not possible to move to {2 2}")
		return
	}
	if err.Error() != "It is not possible to move to {2 2}" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestMovePawnCannotCrossFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.AddFence(game.Fence{game.Position{1, 0}, false})
	//When
	_, err := g1.MovePawn(game.Position{1, 1})
	//Then
	if err == nil {
		t.Error("It is not possible to move to {1 1}")
		return
	}
	if err.Error() != "It is not possible to move to {1 1}" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestMovePawnCannotBeOnTheSamePositionOfAnotherPawn(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1}) // Move Pawn 1
	//When
	_, err := g1.MovePawn(game.Position{1, 1}) // Move Pawn 2
	//Then
	if err == nil {
		t.Error("It is not possible to move to {1 1}")
		return
	}
	if err.Error() != "It is not possible to move to {1 1}" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestNotOverWhenPawnIsInTheBoard(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	//When
	g1, _ := g.MovePawn(game.Position{1, 1})
	//Then
	if g1.Over == true {
		t.Error("The game is not over, pawn is in the middle of the board")
		return
	}
}

func TestOverWhenPawnArrivesGoalLine(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})  //Move Pawn 1
	g2, _ := g1.MovePawn(game.Position{2, 2}) //Move Pawn 2
	//When
	g3, _ := g2.MovePawn(game.Position{2, 1}) //Move Pawn 1
	//Then
	if g3.Over == false {
		t.Error("The game is over, pawn reaches the goal line")
		return
	}
}

func TestOverNoMoreMove(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})  // Move Pawn 1
	g2, _ := g1.MovePawn(game.Position{2, 2}) // Move Pawn 2
	g3, _ := g2.MovePawn(game.Position{2, 1}) // Move Pawn 1
	//When
	_, err := g3.MovePawn(game.Position{1, 2}) // Move Pawn 2
	//Then
	if err == nil {
		t.Error("The game is over, it is not possible to move the pawn")
		return
	}
	if err.Error() != "Game is over, unable to move the pawn" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestOverNoMoreFenceAddition(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})  // Move Pawn 1
	g2, _ := g1.MovePawn(game.Position{2, 2}) // Move Pawn 2
	g3, _ := g2.MovePawn(game.Position{2, 1}) // Move Pawn 1
	//When
	_, err := g3.AddFence(game.Fence{game.Position{0, 0}, true})
	//Then
	if err == nil {
		t.Error("The game is over, it is not possible to add fences")
		return
	}
	if err.Error() != "Game is over, unable to add a fence" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestJumpPawn(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1}) // Move Pawn 1
	//When
	g2, _ := g1.MovePawn(game.Position{0, 1}) // Move Pawn 2
	//Then
	if !g2.Pawns[1].Position.Equals(game.Position{0, 1}) {
		t.Error("The pawn should jump opponent")
	}
}

func TestJumpLeft(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})                    // Move Pawn 1
	g2, _ := g1.AddFence(game.Fence{game.Position{0, 0}, true}) // Add Fence Pawn 2
	//When
	g3, _ := g2.MovePawn(game.Position{2, 0}) // Jump Left Pawn 1
	//Then
	if !g3.Pawns[0].Position.Equals(game.Position{2, 0}) {
		t.Error("The pawn should jump left opponent")
	}
}

func TestJumpLeftImpossibleFence(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})                    // Move Pawn 1
	g2, _ := g1.AddFence(game.Fence{game.Position{1, 0}, true}) // Add Fence Pawn 2
	//When
	_, err := g2.MovePawn(game.Position{2, 0}) // Jump Left Pawn 1
	//Then
	if err == nil {
		t.Error("It is not possible to move to {2 0}")
		return
	}
	if err.Error() != "It is not possible to move to {2 0}" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}

func TestJumpRight(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})                    // Move Pawn 1
	g2, _ := g1.AddFence(game.Fence{game.Position{0, 0}, true}) // Add Fence Pawn 2
	//When
	g3, _ := g2.MovePawn(game.Position{2, 2}) // Jump Right Pawn 1
	//Then
	if !g3.Pawns[0].Position.Equals(game.Position{2, 2}) {
		t.Error("The pawn should jump right opponent")
	}
}

func TestJumpRightImpossibleRight(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})                    // Move Pawn 1
	g2, _ := g1.AddFence(game.Fence{game.Position{1, 1}, true}) // Add Fence Pawn 2
	//When
	_, err := g2.MovePawn(game.Position{2, 2}) // Jump Right Pawn 1
	//Then
	if err == nil {
		t.Error("It is not possible to move to {2 2}")
		return
	}
	if err.Error() != "It is not possible to move to {2 2}" {
		t.Errorf("Not the right error: %s", err.Error())
	}
}


func TestGetPossibleMoves(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	//When
	moves := g.GetPossibleMoves()
	//Then
	checkMoves(t, moves, game.Positions{game.Position{0, 0}, game.Position{1, 1}, game.Position{0, 2}})
}

func TestGetPossibleMovesWithFenceAndJump(t *testing.T) {
	//Given
	setUp()
	configuration := game.Configuration{3}
	g, _ := game.NewGame(configuration)
	g1, _ := g.MovePawn(game.Position{1, 1})                    // Move Pawn 1
	g2, _ := g1.AddFence(game.Fence{game.Position{1, 1}, true}) // Add Fence Pawn 2
	//When
	moves := g2.GetPossibleMoves()
	//Then
	checkMoves(t, moves, game.Positions{game.Position{1, 0}, game.Position{2, 0}, game.Position{0, 1}})
}
