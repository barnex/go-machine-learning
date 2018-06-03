package vs

func Accuracy(m Net, w []float64, lx []LabeledVec) int {
	var nHit int

	guess := make([]float64, m.NumOut())
	for _, lx := range lx {
		Infer(guess, m, w, lx.X)
		if ArgMax(guess) == lx.Label {
			nHit++
		}
	}
	return nHit
}
