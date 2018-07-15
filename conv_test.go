package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestConv(t *testing.T) {
	src := Reshape2(V{
		0, 0, 0, 0, 0,
		0, 1, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 1,
	}, Dim2{5, 4})

	kern := Reshape2(V{
		1, 2,
		3, 4,
	}, Dim2{2, 2})

	want := Reshape2(V{
		4, 3, 0, 0,
		2, 1, 0, 0,
		0, 0, 0, 4,
	}, Dim2{4, 3})

	{
		dst := MakeM(Dim2{4, 3})
		dst.Elem(0)[2] = 666 // test overwrite of existing data
		ConvMM(dst, kern, src)
		test.Eq(t, dst, want)
	}

	{
		y := MakeV(12)
		w := MakeV(4)
		x := src.List
		copyv(w, kern.List)
		Conv(Dim3{2, 2, 1}, Dim2{5, 4}).Eval(y, w, x)
		test.Eqv(t, y, want.List)
	}
}

// Test raw derivatives by comparing to numerical approximation
func TestConv_Diff(t *testing.T) {
	t.Skip("TODO")
	testDiffW(t, Conv(Dim3{3, 3, 2}, Dim2{5, 4}))
	//	testDiffX(t, Conv(Dim3{3, 3, 2}, Dim2{5, 4}))
}
