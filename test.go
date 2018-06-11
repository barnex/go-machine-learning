package vs

func Accuracy(f Net, w []float64, lx []LabeledVec) int {
	var nHit int

	guess := make([]float64, f.NumOut())
	for _, lx := range lx {
		Infer(guess, f, w, lx.X)
		if ArgMax(guess) == lx.Label {
			nHit++
		}
	}
	return nHit
}
