package vs

func GradNumerical(dst []float64, m Model, testdata []LabeledImg) {
	params := m.Params()
	checkSize(len(dst), len(params))

	const delta = 1. / (1024 * 1024)
	for i := range dst {
		backup := params[i]

		params[i] = backup - delta
		v1 := Loss(m, testdata)

		params[i] = backup + delta
		v2 := Loss(m, testdata)

		dst[i] = (v2 - v1) / (2 * delta)

		params[i] = backup
	}
}
