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

func printMat(mat [][]float32) {
	for _, row := range mat {
		for _, v := range row {
			if v == 0 {
				fmt.Print("     ")
			} else {
				fmt.Printf("%.2f ", v)
			}
		}
		fmt.Println()
	}
}
