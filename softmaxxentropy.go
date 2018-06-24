package vs

import "math"

type SoftmaxXEntropy struct {
	nIn int
}

func NewSoftmaxXEntropy(nIn int) *SoftmaxXEntropy {
	return &SoftmaxXEntropy{nIn: nIn}
}

func (f *SoftmaxXEntropy) Eval(y, x V) {
	SoftMax(y, x)
}

// Loss returns:
// 	  -log(softmax(x)_c)
// 	= -log(exp(x_c) / Σ_j exp(x_j))
// During training, c is the label corresponding to training data x.
func (f *SoftmaxXEntropy) Loss(x V, c int) float64 {
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
func (f *SoftmaxXEntropy) GradX(dy V, x V, c int) {
	AssureV(dy, f.NumIn())
	SoftMax(dy, x)
	dy[c] -= 1
}

func (f *SoftmaxXEntropy) NumIn() int    { return f.nIn }
func (f *SoftmaxXEntropy) NumParam() int { return 0 }
func (f *SoftmaxXEntropy) NumOut() int   { return 1 }
