package vs

type Model1 struct {
	params []float64
	w      []Img
	b      []float64
}

func NewModel1() *Model1 {
	return &Model1{}
}

func (m *Model1) NumLabels() int {
	return 10
}

func (m *Model1) Infer(dst []float64, img Img) {

}
