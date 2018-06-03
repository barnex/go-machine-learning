package vs

type MNIST1 struct {
}

var _ Net = &MNIST1{}

func (m *MNIST1) Eval(dst, w, x []float64) {
	m.Linear(dst, w, x)
	SoftMax(dst, dst)
}

func (m *MNIST1) Linear(dst, w, x []float64) {
	const nLabel = 10
	const nInput = 28 * 28
	const nBias = 10

	B := w[(nInput+1)*nLabel:]
	for i := range dst {
		W := w[i*nInput : (i+1)*nInput]
		dst[i] = Dot(W, x) + B[i]
	}
}
