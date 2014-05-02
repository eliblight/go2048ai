package main

import (
	"fmt"
)

type Board interface {
	Get(x, y int) int
	Set(x, y, v int)
}

func BoardPrint(b Board) {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			fmt.Print(b.Get(x, y))
			if x < 3 {
				fmt.Print("\t")

			}
		}
		fmt.Println()
	}
}

func BoardGetEmpty(b Board) [][2]int {
	var r = make([][2]int, 0, 16)
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if b.Get(x, y) == 0 {
				n := len(r)
				r = r[:n+1]
				r[n] = [2]int{x, y}
			}
		}
	}
	return r
}

type BasicBoard struct {
	cells [4][4]int
}

func (b *BasicBoard) Get(x, y int) int {
	return b.cells[x][y]
}

func (b *BasicBoard) Set(x, y, v int) {
	b.cells[x][y] = v
}

func main() {
	var b *BasicBoard = new(BasicBoard)
	BoardPrint(b)
	empty := BoardGetEmpty(b)
	fmt.Println(empty, len(empty))
}
