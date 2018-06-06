package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestLU(t *testing.T) {
	f := NewLU(2, 4)

	w := make([]float64, 2*4+2)

	// A = [1 2 3 4]
	//     [5 6 7 8]
	// B = [9 10]

	f.Ai(w, 0)[0] = 1
	f.Ai(w, 0)[1] = 2
	f.Ai(w, 0)[2] = 3
	f.Ai(w, 0)[3] = 4

	f.Ai(w, 1)[0] = 5
	f.Ai(w, 1)[1] = 6
	f.Ai(w, 1)[2] = 7
	f.Ai(w, 1)[3] = 8

	f.B(w)[0] = 9
	f.B(w)[1] = 10

	tests := []struct {
		x    []float64
		want []float64
	}{
		{[]float64{1, 0, 0, 0}, []float64{1 + 9, 5 + 10}},
		{[]float64{2, 0, 0, 0}, []float64{2 + 9, 10 + 10}},
		{[]float64{0, 0, 0, 1}, []float64{4 + 9, 8 + 10}},
		{[]float64{0, 1, 1, 0}, []float64{2 + 3 + 9, 6 + 7 + 10}},
		{[]float64{0, 0, 0, 0}, []float64{9, 10}},
	}

	for _, c := range tests {
		have := make([]float64, f.NumOut())
		f.Eval(have, w, c.x)
		test.Eqv(t, have, c.want)
	}
}
