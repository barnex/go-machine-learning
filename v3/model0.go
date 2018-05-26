package vs

import "math/rand"

// Model0 does random guessing
type Model0 struct {
}

func NewModel0() *Model0 {
	return &Model0{}
}

func (m *Model0) NumLabels() int {
	return 10
}

func (m *Model0) Infer(dst []float64, img Img) {
	for i := range dst {
		dst[i] = rand.Float64()
	}
	Normalize(dst, dst)
}
