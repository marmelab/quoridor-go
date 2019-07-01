package game

import (
	"errors"
)

// Game is the controller  
type Game struct {
	ID    string `json:"id"`
	Pawn  Pawn  `json:"pawn"`
	Fences []Fence `json:"fences"`
	Board *Board `json:"board"`
}

//AddFence add the fence on the board
func (g Game) AddFence(fence Fence) (Game, error) {
	positionSquare := NewPositionSquare(fence.NWSquare)
	if g.hasAlreadyAFenceAtTheSamePosition(fence.NWSquare) || g.hasNeighbourFence(fence.Horizontal, positionSquare) {
		return Game{}, errors.New("The fence overlaps another one")
	}
	g, err := g.addFenceIfCrossable(fence)
	if err != nil {
		return Game{}, err
	}
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

func (g Game) IsCrossable(fence Fence) bool {
	fences := append(g.Fences, fence)
    column := g.Board.BoardSize - 1
	destinations := []Position{}
	for row := 0; row < g.Board.BoardSize; row++ {
		destinations = append(destinations, Position{column, row})
	}
	return Path(*g.Board, fences, g.Pawn.Position, destinations) != -1
}
