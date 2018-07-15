package vs

import (
	"fmt"
	"io"
)

// M is a matrix.
type M struct {
	List V
	Size Dim2
}

func MakeM(size Dim2) M {
	return Reshape2(make([]float64, size.Len()), size)
}

func Reshape2(list []float64, size Dim2) M {
	checkSize(len(list), size.Len())
	return M{list, size}
}

// Elem returns the i-th row.
func (t M) Elem(i int) V {
	return t.List[i*t.stride() : (i+1)*t.stride()]
}

func (t M) stride() int { return t.Size[0] }

// NumElem returns the number of rows.
func (t M) NumElem() int { return t.Size[1] }

func (t M) PrintTo(w io.Writer) {
	for i := 0; i < t.NumElem(); i++ {
		t.Elem(i).PrintTo(w)
	}
	fmt.Fprintln(w)
}

func (t M) String() string {
	return printToString(t)
}

func mulMV(y V, A M, x V) {
	assureV(y, A.Size[1])
	for i := range y {
		y[i] = A.Elem(i).Dot(x)
	}
}

func mulVM(y V, x V, A M) {
	assureV(y, A.Size[0])
	assureV(x, A.Size[1])
	for i := 0; i < A.Size[0]; i++ {
		y[i] = 0
		for j := 0; j < A.Size[1]; j++ {
			y[i] += x[j] * A.Elem(j)[i]
		}
	}
	/*
	   x: 2
	   A: 2x3 = {3, 2}
	   [x0 x1] [ a00 a01 a02 ]
	           [ a10 a11 a12 ]
	   y: 3
	*/
}
