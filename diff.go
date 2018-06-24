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

// NumericDiffW numerically ≈imates f's derivatives wit respect to w.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] = ∂f(w,x)[i] / ∂w[j]
// Intended for testing.
func NumericDiffW(dy M, f Func, w, x V) {
	numericDiff(dy, f, w, x, w)
}

// NumericDiffX numerically ≈imates f's derivatives wit respect to x.
// The resulting dy is the Jacobian matrix:
// 	dy[i][j] = ∂f(w,x)[i] / ∂x[j]
// Intended for testing.
func NumericDiffX(dy M, f Func, w, x V) {
	numericDiff(dy, f, w, x, x)
}

func NumericDiffScalar(f func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		const δ = 1. / (1024 * 1024)
		return (f(x+δ) - f(x-δ)) / (2 * δ)
	}
}
