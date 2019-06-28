package game

// Fence represents a fence on the board
type Fence struct {
	NWSquare Position `json:square`
	Horizontal bool `json:horizontal`
}

type FenceBlock struct {
	NWSquare Position
	NESquare Position
	SWSquare Position
	SESquare Position   
}

func NewFenceBlock(nwSquare Position) FenceBlock {
	neSquare := nwSquare.Copy(1, 0)
	swSquare := nwSquare.Copy(0, 1)
	seSquare := nwSquare.Copy(1, 1)
	return FenceBlock{nwSquare, neSquare, swSquare, seSquare}
}

type PositionSquare struct {
	NorthPosition Position
	EastPosition Position
	SouthPosition Position
	WestPosition Position   
}

func NewPositionSquare(center Position) PositionSquare {
	northPosition := center.Copy(0, -1);
	eastPosition := center.Copy(1, 0);
	southPosition := center.Copy(0, 1);
	westPosition := center.Copy(-1, 0);
	return PositionSquare{northPosition, eastPosition, southPosition, westPosition}
}

type Edge struct {
	First Position
    Second Position
}
