package vs

// GradLogits numerically approximates component i of f's gradient with respect to w.
func GradLogits(y []float64, f Net, w, x []float64, i int) {
	netCheckSize(f, y, w, x)
	const delta = 1. / (1024 * 1024)

	y1 := y
	y2 := make([]float64, len(y))

	backup := w[i]

	w[i] = backup - delta
	f.Eval(y1, w, x)

	w[i] = backup + delta
	f.Eval(y2, w, x)

	w[i] = backup

	for j := range y {
		y[j] = (y2[j] - y1[j]) / (2 * delta)
	}
}
