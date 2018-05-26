package vs

func Test(m Model, testdata []LabeledImg) (correct int) {
	var nHit int
	guess := make([]float64, m.NumLabels())
	for _, limg := range testdata {
		m.Infer(guess, limg.Img)
		if ArgMax(guess) == limg.Label {
			nHit++
		}
	}
	return nHit
}
