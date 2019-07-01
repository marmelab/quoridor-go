package game

import (
    "container/list"
)

type QueueNode struct {
	Position Position
	Distance int
}

// Path function to find the shortest path between a given source cell to possible destination cells. 
func Path(board Board, fences []Fence, src Position, dest Positions) int {
    boardSize := board.BoardSize
    visited := make([][]bool, boardSize)
    for i := 0; i < boardSize; i++ {
		visited[i] = make([]bool, boardSize)
	}
    visited[src.Column][src.Row] = true;

	q := list.New()
	q.PushBack(QueueNode{src, 0})

    for q.Len() > 0 {
		curr := q.Front().Value.(QueueNode)
        pos := curr.Position
        if (dest.IndexOf(pos) != -1) {
			return curr.Distance
		}
        q.Remove(q.Front())
        ps:= NewPositionSquare(pos)
        positions := [4]Position{ps.EastPosition, ps.NorthPosition, ps.SouthPosition, ps.WestPosition}
        for _, position := range positions {
            if board.IsInBoard(position) && !visited[position.Column][position.Row] && CanMove(pos, position, fences) {
                visited[position.Column][position.Row] = true
                adjPosition := QueueNode{Position{position.Column, position.Row}, curr.Distance + 1 }
                q.PushBack(adjPosition)
            }
        }
    }
    return -1; 
}
