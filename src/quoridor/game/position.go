package game

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
