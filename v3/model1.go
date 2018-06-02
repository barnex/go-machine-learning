package vs

import "math"

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
	return m
}

func (m *Model1) Grad(dst []float64, x LabeledImg) {
	checkSize(len(dst), len(m.Params()))
	c := x.Label

	nL := m.NumLabels()
	f := make([]float64, nL)
	m.Infer(f, x.Img)
	for p := range dst {
		term1 := m.Partial(x, c, p)

		term2 := 0.0
		term3 := 0.0
		for j := 0; j < nL; j++ {
			expFj := math.Exp(f[j])
			term2 += expFj * m.Partial(x, j, p)
			term3 += expFj
		}

		dst[p] = term1 - term2/term3
	}
}

func (m *Model1) Partial(x LabeledImg, d, p int) float64 {
	if d != x.Label {
		return 0
	}

	return x.Img.List[p]
}

func (m *Model1) Infer(dst []float64, img Img) {
	m.Raw(dst, img)
	SoftMax(dst, dst)
}

func (m *Model1) Raw(dst []float64, img Img) {
	checkSize(len(dst), len(m.b))
	N := float64(len(dst))
	for i := range dst {
		dst[i] = Dot(m.w[i].List, img.List)/N + m.b[i]
	}
}

func (m *Model1) Params() []float64 {
	return m.params
}

func (m *Model1) NumLabels() int { return len(digits) }
