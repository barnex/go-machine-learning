package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

// Test raw derivatives by comparing to numerical approximation
func TestMaxPool_Diff(t *testing.T) {
	testDiffW(t, MaxPool1D(10, 2))

	testDiffX(t, MaxPool1D(3, 2))
	testDiffX(t, MaxPool1D(30, 5))
	testDiffX(t, MaxPool1D(30, 6))
}

func TestMaxPool_Eval(t *testing.T) {
	f := MaxPool1D(3, 2)
	have := MakeV(3)
	x := V{
		1, 2,
		4, 3,
		5, 6,
	}
	f.Eval(have, nil, x)
	test.Eqv(t, have, V{2, 4, 6})
}
