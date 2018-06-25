package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

var stdSizes = []Dim2{
	{2, 2},
	{2, 5},
	{5, 2},
	{1, 1},
}

func testDiffW(t *testing.T, f DiffFunc) {
	t.Helper()

	w := randomV(f.NumParam())
	x := randomV(f.NumIn())

	have := MakeM(diffWSize(f))
	f.DiffW(have, w, x)

	want := MakeM(diffWSize(f))
	NumericDiffW(want, f, w, x)

	test.Approxv(t, have.List, want.List, 1e-5)
}

func testDiffX(t *testing.T, f DiffFunc) {
	t.Helper()

	w := randomV(f.NumParam())
	x := randomV(f.NumIn())

	have := MakeM(diffXSize(f))
	f.DiffX(have, w, x)

	want := MakeM(diffXSize(f))
	NumericDiffX(want, f, w, x)

	test.Approxv(t, have.List, want.List, 1e-5)
}

func testGrad(t *testing.T, f *Net, c int) {
	t.Helper()

	randomize(f.w, 1)
	x := randomV(f.NumIn())

	buf := MakeV(f.NumOut())
	have := MakeV(f.NumParam())
	f.Backprop(have, buf, x, c)

	want := MakeV(f.NumParam())
	NumericGrad(want, f, x, c)

	test.Approxv(t, have, want, 1e-5)
}

func randomV(length int) V {
	y := MakeV(length)
	randomize(y, 1)
	return y
}
