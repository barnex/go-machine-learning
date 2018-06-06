package vs

// NumGrad numerically approximates f's gradient with respect to w.
func NumGrad(y T, f Net, w, x []float64) {
	gnx, gny := gradSize(f)
	checkSize(y.Size(0), gnx)
	checkSize(y.Size(1), gny)

	const delta = 1. / (1024 * 1024)

	y1 := make([]float64, f.NumOut())
	y2 := make([]float64, f.NumOut())

	for i := range w {
		backup := w[i]

		w[i] = backup - delta
		f.Eval(y1, w, x)

		w[i] = backup + delta
		f.Eval(y2, w, x)

		w[i] = backup

		for j := 0; j < f.NumOut(); j++ {
			y := y.Row(j)
			y[i] = (y2[j] - y1[j]) / (2 * delta)
		}
	}
}

func gradSize(f Net) (nx, ny int) {
	return f.NumWeight(), f.NumOut()
}
