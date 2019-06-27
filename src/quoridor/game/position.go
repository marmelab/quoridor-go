package game

//Position position on the board
type Position struct {
	Column int `json:"column"`
	Row    int `json:"row"`
}

func (p Position) translateColumn(deltaColumn int) {
	p.Column += deltaColumn
}

func (p Position) translateRow(deltaRow int) {
	p.Row += deltaRow
}

func (p Position) Equals(other Position) bool {
	return other.Column == p.Column && other.Row == p.Row
}

func (p Position) Copy(deltaColumn int, deltaRow int) Position {
	copy := Position{p.Column, p.Row}
	copy.translateColumn(deltaColumn)
	copy.translateRow(deltaRow)
	return copy
}
