package vs

import "fmt"

// Img is a 2D image / matrix.
type Img struct {
	List []float64   // Underlying data as a single list
	Elem [][]float64 // Data as a 2D matrix
}

func NewImg(rows, cols int) Img {
	return ImgFromSlice(make([]float64, cols*rows), rows, cols)

}

func ImgFromSlice(list []float64, rows, cols int) Img {
	checkSize(len(list), rows*cols)
	mat := make([][]float64, rows)
	for iy := range mat {
		mat[iy] = list[iy*cols : (iy+1)*cols]
	}
	return Img{list, mat}
}

// Rows is the number of elements in the first (outer, y) dimension.
func (m Img) Rows() int {
	return len(m.Elem)
}

// Cols is the number of elements in the second (inner, x) dimension.
func (m Img) Cols() int {
	return len(m.Elem[0])
}

// Len returns the total number of elements
func (m Img) Len() int {
	return len(m.List)
}

func (m Img) Print() {
	for _, row := range m.Elem {
		for i, v := range row {
			if i != 0 {
				fmt.Print(" ")
			}
			if v == 0 {
				fmt.Print("     ")
			} else {
				fmt.Printf("%.2f", v)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m Img) Render(min, max float64) {
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
