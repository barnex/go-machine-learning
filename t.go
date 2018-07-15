package vs

import (
	"fmt"
	"io"
)

// T3 is a tensor.
type T3 struct {
	List V
	Size Dim3
}

func MakeT(size Dim3) T3 {
	return Reshape3(make([]float64, size.Len()), size)
}

func Reshape3(list []float64, size Dim3) T3 {
	checkReshape(list, size)
	return T3{list, size}
}

func (t T3) Elem(i int) M {
	return M{
		List: t.List[i*t.stride() : (i+1)*t.stride()],
		Size: Dim2{t.Size[0], t.Size[1]},
	}
}

func (t T3) stride() int { return t.Size[0] * t.Size[1] }

func (t T3) NumElem() int { return t.Size[2] }

func (t T3) PrintTo(w io.Writer) {
	for i := 0; i < t.NumElem(); i++ {
		t.Elem(i).PrintTo(w)
		fmt.Fprintln(w)
	}
}

func (t T3) String() string {
	return printToString(t)
}

func checkReshape(list []float64, d Dim) {
	if len(list) != d.Len() {
		panic(fmt.Sprintf("reshape: %v does not fit %v (len %v)", len(list), d, d.Len()))
	}
}
