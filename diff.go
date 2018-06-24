package vs

// numericDiff numerically approximates f's derivatives with respect to coord,
// which must be w, x, or a subslice thereof.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] \≈ ∂f(w,x)[i] / ∂coord[j]
func numericDiff(dy M, f Func, w, x, coord V) {
	AssureM(dy, Dim2{len(coord), f.NumOut()})

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
			dy := dy.Row(j)
			dy[i] = (y2[j] - y1[j]) / (2 * δ)
		}
	}
}

// NumericDiffW numerically approximates f's derivatives wit respect to w.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] = ∂f(w,x)[i] / ∂w[j]
// Intended for testing.
func NumericDiffW(dy M, f Func, w, x V) {
	numericDiff(dy, f, w, x, w)
}

// NumericDiffX numerically approximates f's derivatives wit respect to x.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] = ∂f(w,x)[i] / ∂x[j]
// Intended for testing.
func NumericDiffX(dy M, f Func, w, x V) {
	numericDiff(dy, f, w, x, x)
}

func NumericGrad(dy V, f *Net, x V, c int) {
	AssureV(dy, f.NumParam())

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
