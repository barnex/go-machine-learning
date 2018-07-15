package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestT4_Elem(t *testing.T) {
	m := MakeT4(Dim4{4, 3, 2, 2})

	test.Eq(t, m.List.Len(), 48)

	list := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,

		13, 14, 15, 16,
		17, 18, 19, 20,
		21, 22, 23, 24,

		11, 2, 3, 4,
		15, 6, 7, 8,
		19, 10, 11, 12,

		113, 14, 15, 16,
		117, 18, 19, 20,
		121, 22, 23, 24,
	}
	copyv(m.List, list)

	test.Eq(t, m.Elem(0).Elem(0), Reshape2([]float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
	}, Dim2{4, 3}))

	test.Eq(t, m.Elem(0).Elem(1), Reshape2([]float64{
		13, 14, 15, 16,
		17, 18, 19, 20,
		21, 22, 23, 24,
	}, Dim2{4, 3}))

	test.Eq(t, m.Elem(1).Elem(0), Reshape2([]float64{
		11, 2, 3, 4,
		15, 6, 7, 8,
		19, 10, 11, 12,
	}, Dim2{4, 3}))

	test.Eq(t, m.Elem(1).Elem(1), Reshape2([]float64{
		113, 14, 15, 16,
		117, 18, 19, 20,
		121, 22, 23, 24,
	}, Dim2{4, 3}))
}
