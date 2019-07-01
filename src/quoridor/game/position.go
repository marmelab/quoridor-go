package game

//Position position on the board
type Position struct {
	Column int `json:"column"`
	Row    int `json:"row"`
}

func (p Position) translateColumn(deltaColumn int) Position {
	p.Column += deltaColumn
	return p
}

func (p Position) translateRow(deltaRow int) Position {
	p.Row += deltaRow
	return p
}

func (p Position) Equals(other Position) bool {
	return other.Column == p.Column && other.Row == p.Row
}

func (p Position) Copy(deltaColumn int, deltaRow int) Position {
	copy := Position{p.Column, p.Row}
	copy = copy.translateColumn(deltaColumn)
	copy = copy.translateRow(deltaRow)
	return copy
}

type Positions []Position

func (positions Positions) IndexOf(e Position) int {
    for index, a := range positions {
        if a.Equals(e) {
            return index
        }
    }
    return -1
 }

type Direction int

 const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
	UNKNOWN
 )

 func getDirection(from Position, to Position) Direction {
    if from.Row == to.Row {
        if from.Column + 1 == to.Column {
            return EAST
        }
        if from.Column - 1 == to.Column {
            return WEST
        }
    }
    if from.Column == to.Column {
        if from.Row - 1 == to.Row {
            return NORTH
        }
        if from.Row + 1 == to.Row {
            return SOUTH
        } 
    }
    return UNKNOWN
}
