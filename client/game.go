package client

import (
)

type Game struct {
	score int
	board Board
}

type GetMove func(board Board, score int) Direction

func NewGame(get_move_fn GetMove) (board *Board, score int) {
	score = 0
	board = new(Board)
	board.FillRandomEmptyCell()
	for ; !board.IsFin() ; {
		move_score, to_spawn := board.CollapseDirection(get_move_fn(*board,score))
		score += move_score
		if (to_spawn) {
			board.FillRandomEmptyCell()
		}
	}
	return board, score 
}

