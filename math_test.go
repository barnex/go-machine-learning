package vs

import (
	"math"
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestMath(t *testing.T) {
	test.Eq(t, ArgMax([]float64{1, 2, 3}), 2)
	test.Eqf(t, Dot([]float64{1, 2, 3}, []float64{4, 5, 6}), 32)
	test.Eqf(t, Len2([]float64{3, 4}), 25)
	test.Eqf(t, Len([]float64{3, 4}), 5)
	{
		dst := make([]float64, 2)
		MAdd(dst, []float64{1, 2}, 3, []float64{4, 5})
		test.Eqv(t, dst, []float64{13, 17})
	}
	{
		dst := make([]float64, 2)
		Map(dst, []float64{9, 16}, math.Sqrt)
		test.Eqv(t, dst, []float64{3, 4})
	}
	{
		min, max := MinMax([]float64{1, 2, 3, 4, 5})
		test.Eqf(t, min, 1)
		test.Eqf(t, max, 5)
	}
	test.Eqf(t, Sum([]float64{1, 2, 3}), 6)
	{
		dst := make([]float64, 2)
		Mul(dst, 3, []float64{1, 2})
		test.Eqv(t, dst, []float64{3, 6})
	}
	{
		dst := make([]float64, 2)
		Set(dst, 3)
		test.Eqv(t, dst, []float64{3, 3})
	}
	test.Eqf(t, Sum([]float64{1, 2, 3}), 6)
}
