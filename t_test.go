package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestT_Elem(t *testing.T) {
	m := MakeT(Dim3{4, 3, 2})

	test.Eq(t, m.List.Len(), 24)

	list := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,

		13, 14, 15, 16,
		17, 18, 19, 20,
		21, 22, 23, 24,
	}
	copyv(m.List, list)

	test.Eq(t, m.Elem(0), Reshape2([]float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
	}, Dim2{4, 3}))

	test.Eq(t, m.Elem(1), Reshape2([]float64{
		13, 14, 15, 16,
		17, 18, 19, 20,
		21, 22, 23, 24,
	}, Dim2{4, 3}))
}
