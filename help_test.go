package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

func testAccuracy(t *testing.T, n *Net, set []LV, want float64) {
	t.Helper()
	have := Accuracy(n, set)
	if have < want {
		t.Errorf("accuracy: have: %v, want: %v", have, want)
	}
}

func testDiffW(t *testing.T, f DiffFunc) {
	t.Helper()

	w := randomV(f.NumParam(), 1234)
	x := randomV(f.NumIn(), 4567)
	y := MakeV(f.NumOut())
	f.Eval(y, w, x)

	have := MakeM(diffWSize(f))
	Set(have.List, 666)
	f.DiffW(have, y, w, x)

	want := MakeM(diffWSize(f))
	numericDiffW(want, f, w, x)

	testApproxM(t, have, want, 1e-5)

}

func testDiffX(t *testing.T, f DiffFunc) {
	t.Helper()

	w := randomV(f.NumParam(), 1234)
	x := randomV(f.NumIn(), 4567)
	y := MakeV(f.NumOut())
	f.Eval(y, w, x)

	have := MakeM(diffXSize(f))
	Set(have.List, 666)
	f.DiffX(have, y, w, x)

	want := MakeM(diffXSize(f))
	numericDiffX(want, f, w, x)

	testApproxM(t, have, want, 1e-5)
}

func testApproxM(t *testing.T, have, want M, tol float64) {
	t.Helper()
	if !test.IsApproxV(have.List, want.List, tol) {
		t.Errorf("have:\n%vwant:\n%v", have, want)
	}
}

func testGrad(t *testing.T, f *Net) {
	t.Helper()

	for c := 0; c < f.NumOut(); c++ {
		Randomize(f.w, 1, 1234)
		x := randomV(f.NumIn(), 1234)

		buf := MakeV(f.NumOut())
		have := MakeV(f.NumParam())
		f.Backprop(have, buf, x, c)

		want := MakeV(f.NumParam())
		numericGrad(want, f, x, c)

		test.Approxv(t, have, want, 1e-5)
	}
}

func randomV(length int, seed int64) V {
	y := MakeV(length)
	Randomize(y, 1, seed)
	return y
}

// numericDiff numerically approximates f's derivatives with respect to coord,
// which must be w, x, or a subslice thereof.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] \≈ ∂f(w,x)[i] / ∂coord[j]
func numericDiff(dy M, f Func, w, x, coord V) {
	assureM(dy, Dim2{len(coord), f.NumOut()})

	const δ = 1. / (1024 * 1024)
	y1 := MakeV(f.NumOut())
	y2 := MakeV(f.NumOut())

	for i := range coord {
		backup := coord[i]

		coord[i] = backup - δ // left
		f.Eval(y1, w, x)

		coord[i] = backup + δ // right
		f.Eval(y2, w, x)

		coord[i] = backup // restore

		for j := 0; j < f.NumOut(); j++ {
			dy := dy.Elem(j)
			dy[i] = (y2[j] - y1[j]) / (2 * δ)
		}
	}
}

// numericDiffW numerically approximates f's derivatives wit respect to w.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] = ∂f(w,x)[i] / ∂w[j]
// Intended for testing.
func numericDiffW(dy M, f Func, w, x V) {
	numericDiff(dy, f, w, x, w)
}

// numericDiffX numerically approximates f's derivatives wit respect to x.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] = ∂f(w,x)[i] / ∂x[j]
// Intended for testing.
func numericDiffX(dy M, f Func, w, x V) {
	numericDiff(dy, f, w, x, x)
}

func numericGrad(dy V, f *Net, x V, c int) {
	assureV(dy, f.NumParam())

	const δ = 1. / (1024 * 1024)
	w := f.w

	yBuf := MakeV(f.NumOut())
	dyBuf := MakeV(f.NumParam())

	for i := range w {
		backup := w[i]

		w[i] = backup - δ // left
		y1 := f.Backprop(dyBuf, yBuf, x, c)

		w[i] = backup + δ // right
		y2 := f.Backprop(dyBuf, yBuf, x, c)

		w[i] = backup // restore

		dy[i] = (y2 - y1) / (2 * δ)
	}
}
