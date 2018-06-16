package vs

import "math"

func AvgGradLoss(g []float64, f Net, w []float64, xl []LabeledVec, buf T) {
	g1 := make([]float64, f.NumWeight())
	Set(g, 0)
	for _, xl := range xl {
		GradLoss(g1, f, w, xl, buf)
		Add(g, g, g1)
	}
	Mul(g, 1/float64(len(xl)), g)
}

// GradLoss calculates the gradient of loss(softmax(f)),
// based on f's gradient of raw logits and the chain rule.
func GradLoss(g []float64, f Net, w []float64, xl LabeledVec, buf T) {
	if buf.List() == nil { // TODO: isnil
		buf = MakeT(gradSize(f))
	}

	x := xl.X
	c := xl.Label
	F := make([]float64, f.NumOut())
	f.Eval(F, w, x)
	f.Grad(buf, w, x)

	den := 0.0
	for j := range F {
		den += math.Exp(F[j])
	}
	for i := range g {
		nom := 0.0
		for j := range F {
			nom += math.Exp(F[j]) * buf.Row(j)[i]
		}
		g[i] = -buf.Row(c)[i] + nom/den
	}
}

// NumGradLoss numerically approximates the gradient of loss(softmax(f)),
// based on f.Eval(), f.Grad().
// Intended for testing.
func NumGradLoss(g []float64, f Net, w []float64, xl LabeledVec) {

	CheckSize(len(g), f.NumOut())
	const delta = 1. / (1024 * 1024)

	buf := make([]float64, f.NumOut())
	for i := range g {
		backup := w[i]

		w[i] = backup - delta
		v1 := Loss1(f, w, xl, buf)

		w[i] = backup + delta
		v2 := Loss1(f, w, xl, buf)

		g[i] = (v2 - v1) / (2 * delta)

		w[i] = backup
	}
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
