package client

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

// Direction enum		
type Direction int

const (
	Left  Direction = iota
	Right           = iota
	Up              = iota
	Down            = iota
)

// Declare a map from direction to its scaning sequence
// as expressed by a set of TupleVectors
var direction_map map[Direction][]TupleVector

func convertArrayToSliceVector(a [4][4]Tuple) []TupleVector {
	
}

func init() {
	var l,r,u,d [4][4]Tuple
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			t := Tuple{i, j}
			l[i][j] = t
			r[3-i][j] = t
			u[j][i] = t
			d[3-j][i] = t
		}
	}
	direction_map[Left] = l
	direction_map[Right] = r
	direction_map[Up] = u
	direction_map[Down] = d
}

type Board struct {
	cells [4][4]int
}

func (b *Board) get(x, y int) int {
	return b.cells[x][y]
}

func (b *Board) set(x, y, v int) {
	b.cells[x][y] = v
}

func (b *Board) Print() {
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

func (b *Board) FillRandomEmptyCell() {
	x, y := b.getRandomEmtpyCell()
	b.set(x, y, GetStartValue())
}

func (b *Board) collapse(v TupleVector) (score int, did_something bool) {
	// single element, recursion exit point
	l := len(v)
	if l == 1 {
		//		fmt.Println("len=1 => exit", v)
		return 0, false
	}

	// If first emelent zero, look for the first non zero element and move it
	// there and re-scan. If none found, recursion end.
	// i.e. 0222 => 2022, collapse(2022)
	x0, y0 := v[0].xy()
	v0 := b.get(x0, y0)
	if v0 == 0 {
		for i := 1; i < l; i++ {
			xi, yi := v[i].xy()
			vi := b.get(xi, yi)
			if vi != 0 {
				b.set(x0, y0, vi)
				b.set(xi, yi, 0)
				score, _ = b.collapse(v[:])
				return score, true
			}
		}
		return 0, false
	}
	// ...so first element is non zero, look for the next non zero element if
	// its the same as v0 combine them and dive in.
	// i.e. 2022 => 4002, collapse (002)
	for i := 1; i < l; i++ {
		xi, yi := v[i].xy()
		vi := b.get(xi, yi)
		if vi == 0 {
			continue
		}
		if v0 == vi {
			b.set(x0, y0, vi*2)
			b.set(xi, yi, 0)
			score, _ = b.collapse(v[1:])
			return score + vi*2, true
		}
		break
	}
	// ...so first element is non zero and next non zero element is different
	// just dive in.
	// i.e. 4022, collapse(022) 
	return b.collapse(v[1:])
}

func (b *Board) CollapseDirection(direction Direction) (score int, done_something bool) {
	tvv, _ := direction_map[direction]
	for _, tv := range tvv {
		s, d := b.collapse(tv)
		score += s
		done_something = done_something || d
	}
	return score, done_something
}

func (b *Board) IsFin() bool {
	// if there is any zero its not done
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if b.get(x, y) == 0 {
				return false
			}
		}
	}
	// Board is full of values, if there are two adject same value
	// cells they can be collapsed. 
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.get(i, j) == b.get(i+1, j) ||
				b.get(i, j) == b.get(i, j+1) {
				return false
			}
		}
	}
	return true
}
