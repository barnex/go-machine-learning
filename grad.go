package vs

// GradLoss calculates the gradient of loss(softmax(f)),
// based on f's gradient of raw logits and the chain rule.
func GradLoss(g []float64, f Net, w, x []float64, buf T) {

}

// NumGradLoss numerically approximates the gradient of loss(softmax(f)),
// based on f.Eval()
// Intended for testing.
func NumGradLoss(g []float64, f Net, w, x []float64) {

}

// NumGrad numerically approximates f's gradient with respect to w.
// The result is stored in g.
// Intended for testing f's analytical Grad() implementation.
func NumGrad(g T, f Net, w, x []float64) {
	gnx, gny := gradSize(f)
	CheckSize(g.Size(0), gnx)
	CheckSize(g.Size(1), gny)

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
			g := g.Row(j)
			g[i] = (y2[j] - y1[j]) / (2 * delta)
		}
	}
}

func gradSize(f Net) (nx, ny int) {
	return f.NumWeight(), f.NumOut()
}
