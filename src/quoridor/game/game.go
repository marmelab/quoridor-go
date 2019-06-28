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
	fenceBlock := NewFenceBlock(fence.NWSquare)
	g, err :=g.addFenceWithEdges(fence, fenceBlock)
	if err != nil {
		return Game{}, err
	}
	return g, nil
}

func (g Game) addFenceWithEdges(fence Fence, fenceBlock FenceBlock) (Game, error) {
	var firstEdge, secondEdge Edge 
	if fence.Horizontal {
		firstEdge = Edge{fenceBlock.NWSquare, fenceBlock.SWSquare} //westEdge
		secondEdge = Edge{fenceBlock.NESquare, fenceBlock.SESquare} //eastEdge
	} else {
		firstEdge = Edge{fenceBlock.NWSquare, fenceBlock.NESquare} //northEdge
		secondEdge = Edge{fenceBlock.SWSquare, fenceBlock.SESquare} //southEdge
	}
	g, err := g.addFenceIfCrossable(fence, firstEdge, secondEdge)
	if err != nil {
		return Game{}, err
	}
	return g, nil
}

func (g Game) addFenceIfCrossable(fence Fence, edge1 Edge, edge2 Edge) (Game, error) {
	if !g.isCrossableForAllPawns() {
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

func (g Game) isCrossableForAllPawns() bool {
	return true
}

func (g Game) isCrossable(pawn Pawn) bool {
	return true
}
