package main

import (
	"fmt"
)
type Board interface {
	Get(x, y int) int
	Set(x, y, v int)
}

func PrintBoard(b Board) {
	for y := 0 ; y < 4 ; y++ {
		for x:= 0 ; x < 4 ; x++ {
			fmt.Print(b.Get(x,y))
			if x < 3 {
				fmt.Print("\t")
			}
		}
		if y < 3 {
			fmt.Println()
		}
	}
}

type BasicBoard struct {
	cells [4][4]int 
}

func (b *BasicBoard) Get(x, y int) int{
	return b.cells[x][y]
}

func (b *BasicBoard) Set(x, y, v int) {
	b.cells[x][y] = v
}


func main() {
	var b *BasicBoard = new(BasicBoard)
	PrintBoard(b)
}

