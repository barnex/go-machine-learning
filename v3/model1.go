package vs

type Model1 struct {
	params []float64
	w      []Img
	b      []float64
}

const (
	imgRows = 28
	imgCols = 28
	imgPix  = imgRows * imgCols
)

func NewModel1() *Model1 {
	params := make([]float64, len(digits)*(imgPix+1))
	m := &Model1{params: params}
	pos := 0
	m.w = make([]Img, len(digits))
	for d := range digits {
		m.w[d] = ImgFromSlice(params[pos:pos+imgPix], imgRows, imgCols)
		pos += imgPix
	}
	m.b = params[pos:]
	checkSize(len(m.b), len(digits))

	Randomize(m.params, 0.01)
	return m
}

func (m *Model1) Params() []float64 {
	return m.params
}

func (m *Model1) Infer(dst []float64, img Img) {
	checkSize(len(dst), len(m.b))

	for i := range dst {
		dst[i] = DotNorm(m.w[i].List, img.List) + m.b[i]
	}
	SoftMax(dst, dst)
}

func (m *Model1) NumLabels() int { return len(digits) }
