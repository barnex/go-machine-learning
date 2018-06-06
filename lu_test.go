package vs

import "testing"

func TestLU(t *testing.T) {
	f := NewLU(2, 3)

	w := make([]float64, 2*3+2)

	// A = [1 2 3]
	//     [4 5 6]
	// B = [7 8]

	f.Ai(w, 0)[0] = 1
	f.Ai(w, 0)[1] = 2
	f.Ai(w, 0)[2] = 3
	f.Ai(w, 1)[0] = 4
	f.Ai(w, 1)[1] = 5
	f.Ai(w, 1)[2] = 6
	f.B(w)[0] = 7
	f.B(w)[1] = 8
}
