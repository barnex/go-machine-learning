package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestM(t *testing.T) {
	m := MakeM(Dim2{3, 2})

	test.Eq(t, m.Len(), 6)

	list := []float64{1, 2, 3, 4, 5, 6}
	copyv(m.List, list)

	test.Eqv(t, m.Row(0), []float64{1, 2, 3})
	test.Eqv(t, m.Row(1), []float64{4, 5, 6})
}

func TestMV(t *testing.T) {
	/*
		[ 1 2 3 ] [1]
		[ 4 5 6 ] [0]
		          [0]
	*/
	A := Reshape(V{1, 2, 3, 4, 5, 6}, Dim2{3, 2})
	y := MakeV(2)

	mulMV(y, A, V{1, 0, 0})
	test.Eqv(t, y, V{1, 4})

	mulMV(y, A, V{0, 1, 0})
	test.Eqv(t, y, V{2, 5})

	mulMV(y, A, V{0, 0, 1})
	test.Eqv(t, y, V{3, 6})
}

func TestVM(t *testing.T) {
	/*
		[1 0] [ 1 2 3 ]
		      [ 4 5 6 ]
	*/
	A := Reshape(V{1, 2, 3, 4, 5, 6}, Dim2{3, 2})
	y := MakeV(3)

	mulVM(y, V{1, 0}, A)
	test.Eqv(t, y, V{1, 2, 3})

	mulVM(y, V{0, 1}, A)
	test.Eqv(t, y, V{4, 5, 6})
}

func TestMVPanic(t *testing.T) {
	A := Reshape(V{1, 2, 3, 4, 5, 6}, Dim2{3, 2})
	y := MakeV(2)
	test.Panic(t, func() { mulMV(y, A, V{1, 2}) }) // need 3-component vector, not 2
}
