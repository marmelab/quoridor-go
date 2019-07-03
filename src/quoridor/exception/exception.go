package exception

type ErrorType int

const (
	OUTSIDE_BOARD ErrorType = iota
	NO_MOVE
	OPPONENT
)

func New(errorType ErrorType, message string) error {
	return &GameError{errorType, message}
}

type GameError struct {
	ErrorType ErrorType
	Message   string
}

func (e *GameError) Error() string {
	return e.Message
}

func MatchGameError(err error, errorType ErrorType) bool {
	if err, ok := err.(*GameError); ok {
		return err.ErrorType == errorType
	}
	return false
}
