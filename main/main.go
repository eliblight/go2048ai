package main

import (
	"github.com/liblight/go2048ai/client"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(5)
	var b *client.Board = new(client.Board)
	keep_going := true
	score := 0
	tscore := 0
	for keep_going {
		tscore += score
		fmt.Println("score", tscore)
		b.FillRandomEmptyCell()
		b.Print()
		score, keep_going = b.CollapseDirection(client.Down)
		if keep_going {
			fmt.Println("Down")
			continue
		}
		score, keep_going = b.CollapseDirection(client.Left)
		if keep_going {
			fmt.Println("left")
			continue
		}
		score, keep_going = b.CollapseDirection(client.Right)
		if keep_going {
			fmt.Println("right")
			continue
		}	
		
		score, keep_going = b.CollapseDirection(client.Up)
		if keep_going {
			fmt.Println("up")
			continue
		}		 
	}
}