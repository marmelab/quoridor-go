package game

import (
	"errors"
	"fmt"

	"quoridor/exception"

	"github.com/lithammer/shortuuid"
)

type Move struct {
	to Position
	jumpLeft Position
	jumpRight Position
}

// Game is the controller
type Game struct {
	ID       string  `json:"id"`
	Over     bool    `json:"over"`
	PawnTurn int     `json:"pawnTurn"`
	Pawns    []Pawn  `json:"pawn"`
	Fences   []Fence `json:"fences"`
	Board    *Board  `json:"board"`
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
	id := shortuuid.New()
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
	g.PawnTurn = g.getNextPawnTurn()
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
	if isHorizontal {
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
	moves := g.GetPossibleMoves()
	if moves.IndexOf(destination) == -1 {
		return Game{}, fmt.Errorf("It is not possible to move to %v", destination)
	}
	g = g.setCurrentPawnPosition(destination)
	over, err := g.isOver()
	if err != nil {
		return Game{}, err
	}
	g.Over = over
	g.PawnTurn = g.getNextPawnTurn()
	return g, nil
}

func (g Game) GetPossibleMoves() Positions {
	positions := Positions{}

	northMove := Move{Position{0, -1}, Position{-1, 0}, Position{1, 0}}
	northPositions := g.getDirectionPossibleMoves(northMove)
	positions = positions.appendIfNotPresent(northPositions)

	eastMove := Move{Position{1, 0}, Position{0, -1}, Position{0, 1}}
	eastPositions := g.getDirectionPossibleMoves(eastMove)
	positions = positions.appendIfNotPresent(eastPositions)

	southMove := Move{Position{0, 1}, Position{-1, 0}, Position{1, 0}}
	southPositions := g.getDirectionPossibleMoves(southMove)
	positions = positions.appendIfNotPresent(southPositions)

	westMove := Move{Position{-1, 0}, Position{0, -1}, Position{0, 1}}
	westPositions := g.getDirectionPossibleMoves(westMove)
	positions = positions.appendIfNotPresent(westPositions)
	return positions
}

func (g Game) getDirectionPossibleMoves(move Move) Positions {
	positions := Positions{}
	from := g.getCurrentPawn().Position
	toPosition, err := g.getPossiblePosition(from, move.to.Column, move.to.Row)
	if err != nil && !exception.MatchGameError(err, exception.OPPONENT) {
		return positions
	}

	if err == nil {
		positions = append(positions, toPosition)
		return positions
	}
	toPosition = from.Copy(move.to.Column, move.to.Row)
	jumpPosition, errJump := g.getPossiblePosition(toPosition, move.to.Column, move.to.Row)
	if errJump == nil {
		positions = append(positions, jumpPosition)
		return positions
	}
	jumpLeftPosition, errLeftJump := g.getPossiblePosition(toPosition, move.jumpLeft.Column, move.jumpLeft.Row)
	if errLeftJump == nil {
		positions = append(positions, jumpLeftPosition)
	}
	jumpRightPosition, errRightJump := g.getPossiblePosition(toPosition, move.jumpRight.Column, move.jumpRight.Row)
	if errRightJump == nil {
		positions = append(positions, jumpRightPosition)
	}
	return positions
}

func (g Game) getPossiblePosition(from Position, col, row int) (Position, error) {
	to := from.Copy(col, row)
	if !g.Board.IsInBoard(to) {
		return Position{}, exception.New(exception.OUTSIDE_BOARD, "Outside")
	}
	if !CanCross(from, to, g.Fences) {
		return Position{}, exception.New(exception.NO_MOVE, "NotCrossable")
	}
	if !isPositionFree(to, g.Pawns) {
		return Position{}, exception.New(exception.OPPONENT, "Opponent")
	}
	return to, nil
}

func isPositionFree(position Position, pawns Pawns) bool {
	return pawns.IndexOf(position) == -1
}

func CanCross(from Position, to Position, fences Fences) bool {
	direction := GetDirection(from, to)
	if direction == UNKNOWN {
		return false
	}
	var fence1, fence2 Fence
	switch direction {
	case EAST:
		fence1 = Fence{Position{from.Column, from.Row - 1}, false}
		fence2 = Fence{Position{from.Column, from.Row}, false}
	case WEST:
		fence1 = Fence{Position{from.Column - 1, from.Row - 1}, false}
		fence2 = Fence{Position{from.Column - 1, from.Row}, false}
	case NORTH:
		fence1 = Fence{Position{from.Column - 1, from.Row - 1}, true}
		fence2 = Fence{Position{from.Column, from.Row - 1}, true}
	case SOUTH:
		fence1 = Fence{Position{from.Column - 1, from.Row}, true}
		fence2 = Fence{Position{from.Column, from.Row}, true}
	default:
		panic("Unknown direction")
	}
	return !fences.Contains(fence1, fence2)
}

func (g Game) isOver() (bool, error) {
	pawn := g.getCurrentPawn()
	if pawn.Goal == EAST {
		return pawn.Position.Column == g.Board.BoardSize-1, nil
	}
	if pawn.Goal == WEST {
		return pawn.Position.Column == 0, nil
	}
	return false, fmt.Errorf("Goal direction not supported %v", pawn.Goal)
}

func (g Game) getNextPawnTurn() int {
	if g.PawnTurn+1 > len(g.Pawns) {
		return 1
	}
	return g.PawnTurn + 1
}

func (g Game) getCurrentPawn() Pawn {
	return g.Pawns[g.PawnTurn-1]
}

func (g Game) setCurrentPawnPosition(newPosition Position) Game {
	g.Pawns[g.PawnTurn-1].Position = newPosition
	return g
}
