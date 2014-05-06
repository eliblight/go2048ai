package src

import (
)

type Game struct {
	score int
	board Board
}



func NewGame() (score int, board *Board) {
	return 0, new(Board)
}

