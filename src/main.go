package src

import (
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
		score, keep_going = b.down()
		if keep_going {
			fmt.Println("down")
			continue
		}
		score, keep_going = b.left()
		if keep_going {
			fmt.Println("left")
			continue
		}
		score, keep_going = b.right()
		if keep_going {
			fmt.Println("right")
			continue
		}		 
		score, keep_going = b.up()
		if keep_going {
			fmt.Println("up")
			continue
		}		 
				 
	}
}