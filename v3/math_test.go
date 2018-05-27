package vs

import (
	"testing"
)

func TestMath(t *testing.T) {
	testEqf(t, Sum([]float64{1, 2, 3}), 6)
	//	testEqf(t, Dot([]float64{1, 2, 3}, []float64{4, 5, 6}), 32./3)

}
