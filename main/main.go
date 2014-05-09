package main

import (
	"github.com/liblight/go2048ai/client"
	"fmt"
	"math/rand"
)

func MyMoveFunc(board client.Board, score int) client.Direction {
	fmt.Println("Score: ",score)
	board.Print()
	var can_move bool

	_, can_move = board.CollapseDirection(client.Down)
	if can_move {
		fmt.Println("Down")
		return client.Down
	}
	_, can_move = board.CollapseDirection(client.Left)
	if can_move {
		fmt.Println("Left")
		return client.Left
	}
	_, can_move = board.CollapseDirection(client.Right)
	if can_move {
		fmt.Println("Right")
		return client.Right
	}
	_, can_move = board.CollapseDirection(client.Up)
	if can_move {
		fmt.Println("Up")
		return client.Up
	}
	panic("Cant move nowhere")		
}

func main() {
	rand.Seed(5)
	board, score := client.NewGame(MyMoveFunc)
	fmt.Println("------ FIN -------")
	fmt.Println("Score: ",score)
	board.Print()
}