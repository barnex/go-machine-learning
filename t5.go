package vs

import (
	"fmt"
	"io"
)

// T5 is a rank-5 tensor.
type T5 struct {
	List V
	Size Dim5
}

func MakeT5(size Dim5) T5 {
	return Reshape5(make([]float64, size.Len()), size)
}

func Reshape5(list []float64, size Dim5) T5 {
	checkSize(len(list), size.Len())
	return T5{list, size}
}

func (t T5) Elem(i int) T4 {
	return T4{
		List: t.List[i*t.stride() : (i+1)*t.stride()],
		Size: Dim4{t.Size[0], t.Size[1], t.Size[2], t.Size[3]},
	}
}

func (t T5) stride() int { return t.Size[0] * t.Size[1] * t.Size[2] * t.Size[3] }

func (t T5) NumElem() int { return t.Size[4] }

func (t T5) PrintTo(w io.Writer) {
	for i := 0; i < t.NumElem(); i++ {
		t.Elem(i).PrintTo(w)
	}
	fmt.Fprintln(w)
}

func (t T5) String() string {
	return printToString(t)
}
