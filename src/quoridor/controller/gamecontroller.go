package gamecontroller

import (
	"errors"

	"quoridor/game"
	"quoridor/storage"

	"github.com/lithammer/shortuuid"
)

// CreateGame create a game with the default configuration
func CreateGame(conf *game.Configuration) (*game.Game, error) {
	boardSize := conf.BoardSize
	board, err := game.NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	lineCenter := (boardSize - 1) / 2
	pawn := game.Pawn{game.Position{0, lineCenter}}	
	id:= shortuuid.New()
	game := game.Game{id, pawn, []game.Fence{}, board}
	storage.Set(game.ID, game)
	return &game, nil
}

// GetGame get the game via its identifier
func GetGame(id string) (game.Game, error) {
	g, found := storage.Get(id)
	if !found {
		return game.Game{}, errors.New("The game does not exist")
	}
	return g.(game.Game), nil
}

//AddFence add the fence on the board
func AddFence(id string, fence game.Fence) (game.Game, error) {
	g, err := GetGame(id)
	if err != nil {
		return game.Game{}, err
	}
	g, errFence := g.AddFence(fence)
	if errFence != nil {
		return game.Game{}, errFence
	}
	storage.Set(g.ID, g)
	return g, nil
}


//AddFencePossibilities get all the possibiles places where to add a fence
func GetFencePossibilities(id string) ([]game.Fence, error) {
	g, err := GetGame(id)
	if err != nil {
		return []game.Fence{}, err
	}
	allPossibilities := addAllPossibilities(g.Board.BoardSize)
	possibilities := removeFences(allPossibilities, g)
	return possibilities, nil
}

func addAllPossibilities(boardSize int) []game.Fence  {
	numberOfIntersections := boardSize - 1;
	actions := []game.Fence{}
	for row := 0; row < numberOfIntersections; row++ {
		for column := 0; column < numberOfIntersections; column++ {
			actions = append(actions, game.Fence{game.Position{column, row}, true})
			actions = append(actions, game.Fence{game.Position{column, row}, false})
		}
	}
	return actions
}

func removeFences(allPossibilities []game.Fence, g game.Game) game.Fences {
	var possibilities game.Fences
	possibilities = append(allPossibilities[0:0], allPossibilities...)
	var fences game.Fences
	fences = g.Fences
	for _, element := range allPossibilities {
		if fences.IndexOf(element) > -1 {
			possibilities = possibilities.RemoveFence(element)
			positionSquare := game.NewPositionSquare(element.NWSquare)
			if element.Horizontal {
				rightFence := game.Fence{positionSquare.WestPosition, true}
				possibilities = possibilities.RemoveFence(rightFence)
				leftFence := game.Fence{positionSquare.EastPosition, true}
				possibilities = possibilities.RemoveFence(leftFence)
				oppositeFence := game.Fence{element.NWSquare, false}
				possibilities = possibilities.RemoveFence(oppositeFence)
			} else {
				upFence := game.Fence{positionSquare.NorthPosition, false}
				possibilities = possibilities.RemoveFence(upFence)
				downFence := game.Fence{positionSquare.SouthPosition, false}
				possibilities = possibilities.RemoveFence(downFence)
				oppositeFence := game.Fence{element.NWSquare, true}
				possibilities = possibilities.RemoveFence(oppositeFence)
			}
		} else {
			if (!g.IsCrossable(element)) {
				possibilities = possibilities.RemoveFence(element)
			}
		}
	}
	return possibilities
}

//MovePawn move the pawn on the board
func MovePawn(id string, to game.Position) (game.Game, error) {
	g, err := GetGame(id)
	if err != nil {
		return game.Game{}, err
	}
	g, errPawn := g.MovePawn(to)
	if errPawn != nil {
		return game.Game{}, errPawn
	}
	storage.Set(g.ID, g)
	return g, nil
}
