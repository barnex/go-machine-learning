package main

import "fmt"

type Mat struct {
	List []float32
	Elem [][]float32
}

func NewMat(rows, cols int) Mat {

	list := make([]float32, cols*rows)
	mat := make([][]float32, rows)
	for iy := range mat {
		mat[iy] = list[iy*cols : (iy+1)*cols]
	}
	return Mat{list, mat}
}

func (m Mat) Rows() int {
	return len(m.Elem)
}

func (m Mat) Cols() int { return len(m.Elem[0]) }

func (m Mat) Len() int { return len(m.List) }

func (m Mat) Print() {
	for _, row := range m.Elem {
		for _, v := range row {
			if v == 0 {
				fmt.Print("     ")
			} else {
				fmt.Printf("%.2f ", v)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m Mat) Render(min, max float32) {
	// see https://en.wikipedia.org/wiki/ANSI_escape_code
	const black = 232
	const white = 255
	for _, row := range m.Elem {
		for _, v := range row {
			col := int(black + ((v-min)/(max-min))*(white-black))
			if col < black {
				col = 12 // blue, underflow
			}
			if col > white {
				col = 9 // red, overflow
			}
			fmt.Printf("\033[48;5;%dm% 3.2f", col, v)
		}
		fmt.Println("\033[m")
	}
	fmt.Println()
}

//  "\033[48;5;245mCOLOR1\033[m"
