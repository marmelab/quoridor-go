package game

import (
	"errors"
	"fmt"

	"github.com/lithammer/shortuuid"
)

// Game is the controller  
type Game struct {
	ID string `json:"id"`
	Over bool `json:"over"`
	PlayerTurn int `json:"playerTurn"`
	Pawns []Pawn  `json:"pawn"`
	Fences []Fence `json:"fences"`
	Board *Board `json:"board"`
}

// NewGame create a new game depending on the configuration
func NewGame(conf Configuration) (Game, error) {
	boardSize := conf.BoardSize
	board, err := NewBoard(boardSize)
	if err != nil {
		return Game{}, err
	}
	lineCenter := (boardSize - 1) / 2
	pawns := []Pawn{
		Pawn{Position{0, lineCenter}, EAST},
		Pawn{Position{boardSize - 1, lineCenter}, WEST},
	}
	id:= shortuuid.New()
	return Game{id, false, 1, pawns, []Fence{}, board}, nil
}

// AddFence add the fence on the board
func (g Game) AddFence(fence Fence) (Game, error) {
	if g.Over {
		return Game{}, errors.New("Game is over, unable to add a fence")
	}
	positionSquare := NewPositionSquare(fence.NWSquare)
	if g.hasAlreadyAFenceAtTheSamePosition(fence.NWSquare) || g.hasNeighbourFence(fence.Horizontal, positionSquare) {
		return Game{}, errors.New("The fence overlaps another one")
	}
	g, err := g.addFenceIfCrossable(fence)
	if err != nil {
		return Game{}, err
	}
	g.PlayerTurn = g.whoIsNext()
	return g, nil
}

func (g Game) addFenceIfCrossable(fence Fence) (Game, error) {
	if !g.IsCrossable(fence) {
		return Game{}, errors.New("No more access to goal line")
	}
	g.Fences = append(g.Fences, fence)
	return g, nil
}

func (g Game) hasAlreadyAFenceAtTheSamePosition(p Position) bool {
	for i := range g.Fences {
		pos := g.Fences[i].NWSquare
		if pos.Equals(p) {
			return true
		}
	}
	return false
}

func (g Game) hasNeighbourFence(isHorizontal bool, ps PositionSquare) bool {
	if (isHorizontal) {
		for i := range g.Fences {
			fence := g.Fences[i]
			pos := fence.NWSquare
			if fence.Horizontal && (pos.Equals(ps.EastPosition) || pos.Equals(ps.WestPosition)) {
				return true
			}
		}
		return false
	}
	for i := range g.Fences {
		fence := g.Fences[i]
		pos := fence.NWSquare
		if !fence.Horizontal && (pos.Equals(ps.NorthPosition) || pos.Equals(ps.SouthPosition)) {
			return true
		}
	}
	return false
}

// IsCrossable check whether the fence can be added and let a path for all pawns to their goal line
func (g Game) IsCrossable(fence Fence) bool {
	fences := append(g.Fences, fence)
	for i := range g.Pawns {
		pawn := g.Pawns[i]
		destinations := g.getGoalLine(pawn)
		 if Path(*g.Board, fences, pawn.Position, destinations) == -1 {
			return false
		 }
	}
	return true
}

func (g Game) getGoalLine(pawn Pawn) []Position {
	destinations := []Position{}
	var column int
	if pawn.Goal == EAST {
		column = g.Board.BoardSize - 1
	} else if pawn.Goal == WEST {
		column = 0
	}
	for row := 0; row < g.Board.BoardSize; row++ {
		destinations = append(destinations, Position{column, row})
	}
	return destinations
}

// MovePawn move the pawn to the destination
func (g Game) MovePawn(destination Position) (Game, error) {
	if g.Over {
		return Game{}, errors.New("Game is over, unable to move the pawn")
	}
	if !g.Board.IsInBoard(destination) {
		return Game{}, errors.New("The new position is not inside the board")
	}
	from := g.getCurrentPawn().Position
	direction := GetDirection(from, destination)
	if (direction == UNKNOWN) {
		return Game{}, fmt.Errorf("It is not possible to reach %v", destination)
	}
	if !CanMove(from, destination, g.Fences, g.Pawns) {
		return Game{}, fmt.Errorf("It is not possible to move to %v", destination)
	}
	g = g.setCurrentPawnPosition(destination)
	over, err := g.isOver()
	if (err != nil) {
		return Game{}, err
	}
	g.Over = over
	g.PlayerTurn = g.whoIsNext()
	return g, nil
}

func (g Game) isOver() (bool, error) {
	pawn := g.getCurrentPawn()
	if pawn.Goal == EAST {
		return pawn.Position.Column == g.Board.BoardSize - 1, nil
	}
	if pawn.Goal == WEST {
		return pawn.Position.Column == 0, nil
	}
	return false, fmt.Errorf("Goal direction not supported %v", pawn.Goal)
}

func (g Game) whoIsNext() int {
	if g.PlayerTurn + 1 > len(g.Pawns) {
		return 1
	}
	return g.PlayerTurn + 1
}

func (g Game) getCurrentPawn() Pawn {
	return g.Pawns[g.PlayerTurn -1]
}

func (g Game) setCurrentPawnPosition(newPosition Position) Game {
	g.Pawns[g.PlayerTurn -1].Position = newPosition
	return g
}
