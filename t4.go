package vs

import (
	"fmt"
	"io"
)

// T4 is a rank-4 tensor.
type T4 struct {
	List V
	Size Dim4
}

func MakeT4(size Dim4) T4 {
	return Reshape4(make([]float64, size.Len()), size)
}

func Reshape4(list []float64, size Dim4) T4 {
	checkSize(len(list), size.Len())
	return T4{list, size}
}

func (t T4) Elem(i int) T3 {
	return T3{
		List: t.List[i*t.stride() : (i+1)*t.stride()],
		Size: Dim3{t.Size[0], t.Size[1], t.Size[2]},
	}
}

func (t T4) stride() int { return t.Size[0] * t.Size[1] * t.Size[2] }

func (t T4) NumElem() int { return t.Size[3] }

func (t T4) PrintTo(w io.Writer) {
	for i := 0; i < t.NumElem(); i++ {
		t.Elem(i).PrintTo(w)
		fmt.Fprintln(w)
	}
}

func (t T4) String() string {
	return printToString(t)
}
