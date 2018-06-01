package vs

func Test(m Model, testdata []LabeledImg) (correct int) {
	assert(len(testdata) > 0)

	var nHit int
	guess := make([]float64, m.NumLabels())
	for _, limg := range testdata {
		assert(len(limg.Img.List) > 0)
		m.Infer(guess, limg.Img)
		if ArgMax(guess) == limg.Label {
			nHit++
		}
	}
	return nHit
}
