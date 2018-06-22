package vs

import "math"

type SoftmaxXEntropy struct {
	nIn int
}

func NewSoftmaxXEntropy(nIn int) *SoftmaxXEntropy {
	return &SoftmaxXEntropy{nIn: nIn}
}

// Eval returns:
// 	  -log(softmax(x)_c)
// 	= -log(exp(x_c) / Σ_j exp(x_j))
// During training, c is the label corresponding to training data x.
func (f *SoftmaxXEntropy) Eval(x V, c int) float64 {
	buf := MakeV(f.NumIn()) // TODO: don't alloc
	SoftMax(buf, x)
	return -math.Log(buf[c])
}

// GradX calculates the gradient of -log(softmax(x)_c) with respect to x:
//
//  dy[i] = ∂[ -log(softmax(x)_c) ] / ∂x[i]
//        = -δ_ic + softmax(x)[i]
//
// During training, c is the label corresponding to training data x.
func (f *SoftmaxXEntropy) GradX(dy *V, x V, c int) {
	AssureV(dy, f.NumIn())
	SoftMax(*dy, x)
	(*dy)[c] -= 1
}

func (f *SoftmaxXEntropy) NumIn() int    { return f.nIn }
func (f *SoftmaxXEntropy) NumParam() int { return 0 }
func (f *SoftmaxXEntropy) NumOut() int   { return 1 }

func (f *SoftmaxXEntropy) FixLabel(l int) Func {
	return &sMXEnNet{
		Label:           l,
		SoftmaxXEntropy: f,
	}
}

type sMXEnNet struct {
	Label int
	*SoftmaxXEntropy
}

func (f sMXEnNet) Eval(y *V, w, x V) {
	AssureV(y, 1)
	CheckSize(len(w), 0)
	(*y)[0] = f.SoftmaxXEntropy.Eval(x, f.Label)
}

func (f sMXEnNet) DiffX(y *M, w, x V) {
	AssureM(y, Dim2{f.NumIn(), 1})
	CheckSize(len(w), 0)
	y0 := y.Row(0)
	f.SoftmaxXEntropy.GradX(&y0, x, f.Label)
}

func (f sMXEnNet) DiffW(y *M, w, x V) {
	AssureM(y, Dim2{0, 1})
	CheckSize(len(w), 0)
	CheckSize(len(x), f.NumIn())
	return
}
