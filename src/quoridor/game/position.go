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
