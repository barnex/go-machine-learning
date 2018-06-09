package vs

// A Model combines a Net with set weights.
type Model struct {
	F Net
	W []float64
}

func NewModel(f Net) *Model {
	return &Model{
		F: f,
		W: make([]float64, f.NumWeight()),
	}
}

func (m *Model) Loss(xl []LabeledVec) float64 {
	return Loss(m.F, m.W, xl, nil)
}

func (m *Model) Infer(y, x []float64) {
	m.Eval(y, x)
	SoftMax(y, y)
}

func (m *Model) Eval(y, x []float64) {
	m.F.Eval(y, m.W, x)
}

func (m *Model) Grad(y T, x []float64) {
	m.F.Grad(y, m.W, x)
}
