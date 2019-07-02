package gamecontroller

import (
	"errors"

	"quoridor/game"
	"quoridor/storage"
)

const (
	UnknownPlayer = 0
	FirstPlayer = 1
) 

type Party struct {
	game game.Game
	players map[string]int
}

func (p Party) isReady() bool {
	return len(p.players) == len(p.game.Pawns)
}

func getParty(id string) (Party, error) {
	p, found := storage.Get(id)
	if !found {
		return Party{}, errors.New("The game does not exist")
	}
	return p.(Party), nil
}

func checkPlayerCanPlay(p Party, playerID string) error {
	player := p.players[playerID]
	if player == UnknownPlayer {
		return errors.New("Forbidden")
	}
	if !p.isReady() {
		return errors.New("Game is not ready")
	}
	if player != p.game.PlayerTurn {
		return errors.New("It is not your turn")
	}
	return nil
}

// CreateGame create a game with the default configuration
func CreateGame(conf game.Configuration, playerID string) (*game.Game, error) {
	game, err := game.NewGame(conf)
	if err != nil {
		return nil, err
	}
	players := make(map[string]int)
	players[playerID] = FirstPlayer
	storage.Set(game.ID, Party{game, players})
	return &game, nil
}

// GetGame get the game via its identifier
func GetGame(id string) (game.Game, error) {
	p, err := getParty(id)
	if err != nil {
		return game.Game{}, err
	}
	return p.game, nil
}

// JoinGame add a new player to the game
func JoinGame(id string, playerID string) error {
	p, err := getParty(id)
	if err != nil {
		return err
	}
	if p.isReady() {
		return errors.New("Game is already set")
	}
	p.players[playerID] = len(p.players) + 1
	storage.Set(p.game.ID, p)
	return nil
}

// AddFence add the fence on the board
func AddFence(id string, fence game.Fence, playerID string) (game.Game, error) {
	p, err := getParty(id)
	if err != nil {
		return game.Game{}, err
	}
	errPlayer := checkPlayerCanPlay(p, playerID)
	if errPlayer != nil {
		return game.Game{}, errPlayer
	}
	g, errFence := p.game.AddFence(fence)
	if errFence != nil {
		return game.Game{}, errFence
	}
	p.game = g
	storage.Set(p.game.ID, p)
	return g, nil
}

// GetFencePossibilities get all the possibiles places where to add a fence
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

// MovePawn move the pawn on the board
func MovePawn(id string, destination game.Position, playerID string) (game.Game, error) {
	p, err := getParty(id)
	if err != nil {
		return game.Game{}, err
	}
	errPlayer := checkPlayerCanPlay(p, playerID)
	if errPlayer != nil {
		return game.Game{}, errPlayer
	}
	g, errPawn := p.game.MovePawn(destination)
	if errPawn != nil {
		return game.Game{}, errPawn
	}
	p.game = g
	storage.Set(p.game.ID, p)
	return g, nil
}
