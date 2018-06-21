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

func testDiffW(t *testing.T, f Func) {
	t.Helper()

	w := randomV(f.NumParam())
	x := randomV(f.NumIn())

	var have M
	f.DiffW(&have, w, x)

	var want M
	NumericDiffW(&want, f, w, x)

	test.Approxv(t, have.List, want.List, 1e-5)
}

func testDiffX(t *testing.T, f Func) {
	t.Helper()

	w := randomV(f.NumParam())
	x := randomV(f.NumIn())

	var have M
	f.DiffX(&have, w, x)

	var want M
	NumericDiffX(&want, f, w, x)

	test.Approxv(t, have.List, want.List, 1e-5)
}

func randomV(length int) V {
	y := MakeV(length)
	Randomize(y, 1)
	return y
}
