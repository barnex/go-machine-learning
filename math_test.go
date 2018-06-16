package vs

import (
	"math"
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestMath(t *testing.T) {
	{
		dst := make(V, 2)
		Add(dst, V{1, 2}, V{4, 5})
		test.Eqv(t, dst, V{5, 7})
	}
	test.Eq(t, ArgMax(V{1, 2, 3}), 2)
	test.Eqf(t, Dot(V{1, 2, 3}, V{4, 5, 6}), 32)
	test.Eqf(t, Len2(V{3, 4}), 25)
	test.Eqf(t, Len(V{3, 4}), 5)
	{
		dst := make(V, 2)
		MAdd(dst, V{1, 2}, 3, V{4, 5})
		test.Eqv(t, dst, V{13, 17})
	}
	{
		dst := make(V, 2)
		Map(dst, V{9, 16}, math.Sqrt)
		test.Eqv(t, dst, V{3, 4})
	}
	{
		min, max := MinMax(V{1, 2, 3, 4, 5})
		test.Eqf(t, min, 1)
		test.Eqf(t, max, 5)
	}
	test.Eqf(t, Sum(V{1, 2, 3}), 6)
	{
		dst := make(V, 2)
		Mul(dst, 3, V{1, 2})
		test.Eqv(t, dst, V{3, 6})
	}
	{
		dst := make(V, 2)
		Set(dst, 3)
		test.Eqv(t, dst, V{3, 3})
	}
	test.Eqf(t, Sum(V{1, 2, 3}), 6)
}
