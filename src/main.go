package main

import (
	"src/board"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(5)
	var b *Board = new(Board)
	keep_going := true
	score := 0
	tscore := 0
	for keep_going {
		tscore += score
		fmt.Println("score", tscore)
		b.fillRandomEmptyCell()
		b.print()
		score, keep_going = b.CollapseDirection(Down)
		if keep_going {
			fmt.Println("Down")
			continue
		}
		score, keep_going = b.CollapseDirection(Left)
		if keep_going {
			fmt.Println("left")
			continue
		}
		score, keep_going = b.CollapseDirection(Right)
		if keep_going {
			fmt.Println("right")
			continue
		}	
		
		score, keep_going = b.CollapseDirection(Up)
		if keep_going {
			fmt.Println("up")
			continue
		}		 
				 
	}
}