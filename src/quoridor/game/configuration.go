package game

//Configuration options to create a game
type Configuration struct {
	BoardSize int `json:boardSize`
	NumberOfFencesPerPawnPlayer int `json:numberOfFencesPerPlayer`
}
