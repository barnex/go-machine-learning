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
	y := MakeV(f.NumOut())
	f.Eval(y, w, x)

	have := MakeM(diffWSize(f))
	f.DiffW(have, y, w, x)

	want := MakeM(diffWSize(f))
	NumericDiffW(want, f, w, x)

	test.Approxv(t, have.List, want.List, 1e-5)
}

func testDiffX(t *testing.T, f DiffFunc) {
	t.Helper()

	w := randomV(f.NumParam())
	x := randomV(f.NumIn())
	y := MakeV(f.NumOut())
	f.Eval(y, w, x)

	have := MakeM(diffXSize(f))
	f.DiffX(have, y, w, x)

	want := MakeM(diffXSize(f))
	NumericDiffX(want, f, w, x)

	test.Approxv(t, have.List, want.List, 1e-5)
}

func testGrad(t *testing.T, f *Net) {
	t.Helper()

	for c := 0; c < f.NumOut(); c++ {
		randomize(f.w, 1)
		x := randomV(f.NumIn())

		buf := MakeV(f.NumOut())
		have := MakeV(f.NumParam())
		f.Backprop(have, buf, x, c)

		want := MakeV(f.NumParam())
		NumericGrad(want, f, x, c)

		test.Approxv(t, have, want, 1e-5)
	}
}

func randomV(length int) V {
	y := MakeV(length)
	randomize(y, 1)
	return y
}
