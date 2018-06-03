package img

// Img is a 2D image / matrix.
type Img struct {
	List []float64   // Underlying data as a single list
	Elem [][]float64 // Data as a 2D matrix
}

func New(rows, cols int) Img {
	return FromSlice(make([]float64, cols*rows), rows, cols)

}

func FromSlice(list []float64, rows, cols int) Img {
	//checkSize(len(list), rows*cols)
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
