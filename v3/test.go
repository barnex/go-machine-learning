package vs

import "math"

func Loss(m Model, testdata []LabeledImg) float64 {
	loss := 0.0

	infer := make([]float64, m.NumLabels())
	for _, img := range testdata {
		m.Infer(infer, img.Img)

		for _, v := range infer {
			assert(v > 0)
		}
		assert(approxEq(Sum(infer), 1, 1e-6))

		loss += -math.Log(infer[img.Label])
	}
	return loss / float64(len(testdata))
}

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
