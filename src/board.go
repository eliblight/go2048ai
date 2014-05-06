package src

import (
	"fmt"
	"math/rand"
)

func GetStartValue() int {
	if rand.Intn(10) == 0 {
		return 4
	}
	return 2
}

type Tuple [2]int

func (t *Tuple) xy() (x, y int) {
	return t[0], t[1]
}

type TupleVector []Tuple

func (p *TupleVector) append(t Tuple) {
	l := len(*p)
	*p = (*p)[:l+1]
	(*p)[l] = t
}

var direction_map = map[string][]TupleVector{
	"left": {
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
		{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		{{0, 3}, {1, 3}, {2, 3}, {3, 3}}},
	"right": {
		{{3, 0}, {2, 0}, {1, 0}, {0, 0}},
		{{3, 1}, {2, 1}, {1, 1}, {0, 1}},
		{{3, 2}, {2, 2}, {1, 2}, {0, 2}},
		{{3, 3}, {2, 3}, {1, 3}, {0, 3}}},
	"up": {
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
		{{3, 0}, {3, 1}, {3, 2}, {3, 3}}},
	"down": {
		{{0, 3}, {0, 2}, {0, 1}, {0, 0}},
		{{1, 3}, {1, 2}, {1, 1}, {1, 0}},
		{{2, 3}, {2, 2}, {2, 1}, {2, 0}},
		{{3, 3}, {3, 2}, {3, 1}, {3, 0}}}}

type Board struct {
	cells [4][4]int
}

func (b *Board) get(x, y int) int {
	return b.cells[x][y]
}

func (b *Board) set(x, y, v int) {
	b.cells[x][y] = v
}

func (b *Board) print() {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			fmt.Print(b.get(x, y))
			if x < 3 {
				fmt.Print("\t")

			}
		}
		fmt.Println()
	}
}

func (b *Board) getEmptyCells() TupleVector {
	empty_cells := make(TupleVector, 0, 16)
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if b.get(x, y) == 0 {
				empty_cells.append(Tuple{x, y})
			}
		}
	}
	return empty_cells
}

func (b *Board) getRandomEmtpyCell() (int, int) {
	empty_cells := b.getEmptyCells()
	cell := empty_cells[rand.Intn(len(empty_cells))]
	return cell[0], cell[1]
}

func (b *Board) fillRandomEmptyCell() {
	x, y := b.getRandomEmtpyCell()
	b.set(x, y, GetStartValue())
}

func (b *Board) collapse(v TupleVector) (score int, did_something bool) {
	// recursion exit point
	//	{
	//		fmt.Print("v=")
	//		for _, t := range v {
	//			fmt.Print(b.get(t.xy())," ")
	//		}
	//		fmt.Println()
	//	}
	l := len(v)
	if l == 1 {
//		fmt.Println("len=1 => exit", v)
		return 0, false
	}

	// If first emelent zero, look for the first non zero element and move it
	// there. If non found recursion end.
	x0, y0 := v[0].xy()
	v0 := b.get(x0, y0)
	if v0 == 0 {
		for i := 1; i < l; i++ {
			xi, yi := v[i].xy()
			vi := b.get(xi, yi)
			if vi != 0 {
//				fmt.Println("move non zero ", vi)
				b.set(x0, y0, vi)
				b.set(xi, yi, 0)
				score, _ = b.collapse(v[:])
				return score, true
			}
		}
		//		fmt.Println("All zeros => bailout")
		return 0, false
	}
	// ...else look for the next non zero element if its the same as v0
	// combine them
	for i := 1; i < l; i++ {
		xi, yi := v[i].xy()
		vi := b.get(xi, yi)
		if vi == 0 {
			continue
		}
		if v0 == vi {
//			fmt.Println("merge ", vi)
			b.set(x0, y0, vi*2)
			b.set(xi, yi, 0)
			score, _ = b.collapse(v[1:])
			return score + vi*2, true
		}
		break
	}
	return b.collapse(v[1:])
}

func (b *Board) goDir(dir string) (score int, done_something bool) {
	
	done_something = false
	score = 0
	tvv, _ := direction_map[dir]
	for _, tv := range tvv {
		s, d := b.collapse(tv)
		score += s
		done_something = done_something || d
	}
	return score, done_something
}

func (b *Board) left() (score int, done_something bool) {
	return b.goDir("left")
}

func (b *Board) right() (score int, done_something bool) {
	return b.goDir("right")
}

func (b *Board) up() (score int, done_something bool) {
	return b.goDir("up")
}

func (b *Board) down() (score int, done_something bool) {
	return b.goDir("down")
}

