package game

import (
	"errors"

	"quoridor/storage"

	"github.com/lithammer/shortuuid"
)

// Game is the controller  
type Game struct {
	ID    string `json:"id"`
	Pawn  Pawn  `json:"pawn"`
	Fences []Fence `json:"fences"`
	Board *Board `json:"board"`
}

//AddFence add the fence on the board
func (g Game) AddFence(fence Fence) error {
	positionSquare := NewPositionSquare(fence.NWSquare)
	if g.hasAlreadyAFenceAtTheSamePosition(fence.NWSquare) || g.hasNeighbourFence(fence.Horizontal, positionSquare) {
		return errors.New("The fence overlaps another one")
	}
	fenceBlock := NewFenceBlock(fence.NWSquare)
	g.addFenceWithEdges(fence, fenceBlock)
	return nil
}

func (g Game) addFenceWithEdges(fence Fence, fenceBlock FenceBlock) error {
	var firstEdge, secondEdge Edge 
	if fence.Horizontal {
		firstEdge = Edge{fenceBlock.NWSquare, fenceBlock.SWSquare} //westEdge
		secondEdge = Edge{fenceBlock.NESquare, fenceBlock.SESquare} //eastEdge
	} else {
		firstEdge = Edge{fenceBlock.NWSquare, fenceBlock.NESquare} //northEdge
		secondEdge = Edge{fenceBlock.SWSquare, fenceBlock.SESquare} //southEdge
	}
	err := g.addFenceIfCrossable(fence, firstEdge, secondEdge)
	if err != nil {
		return err
	}
	return nil
}

func (g Game) addFenceIfCrossable(fence Fence, edge1 Edge, edge2 Edge) error {
	if !g.isCrossableForAllPawns() {
		return errors.New("No more access to goal line")
	}
	g.Fences = append(g.Fences, fence)
	return nil
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

// CreateGame create a game with the default configuration
func CreateGame(conf *Configuration) (*Game, error) {
	boardSize := conf.BoardSize
	board, err := NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	lineCenter := (boardSize - 1) / 2
	pawn := Pawn{Position{0, lineCenter}}	
	id:= shortuuid.New()
	game := Game{id, pawn, []Fence{}, board}
	storage.Set(game.ID, game)
	return &game, nil
}

// GetGame get the game via its identifier
func GetGame(id string) (Game, error) {
	game, found := storage.Get(id)
	if !found {
		return Game{}, errors.New("The game does not exist")
	}
	return game.(Game), nil
}

//AddFence add the fence on the board
func AddFence(id string, fence Fence) (Game, error) {
	game, err := GetGame(id)
	if err != nil {
		return Game{}, err
	}
	errFence := game.AddFence(fence)
	if errFence != nil {
		return Game{}, errFence
	}
	storage.Set(game.ID, game)
	return game, nil
}
