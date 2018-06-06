package vs

import (
	. "github.com/barnex/vectorstream"
)

type MNIST1 struct{}

var _ Net = &MNIST1{}

var digits = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

const (
	numOut  = 10
	numIn   = 28 * 28
	numBias = 10
)

func (f *MNIST1) NumOut() int {
	return numOut
}

func (f *MNIST1) NumWeight() int {
	return numOut*numIn + numBias
}

func (f *MNIST1) NumIn() int {
	return numIn
}

func (f *MNIST1) Eval(dst, w, x []float64) {
	NetCheckSize(f, dst, w, x)
	B := f.Bias(w)
	for i := range dst {
		W := f.Weight(w, i)
		dst[i] = Dot(W, x) + B[i]
	}
}

func (f *MNIST1) Bias(w []float64) []float64 {
	CheckSize(len(w), f.NumWeight())
	b := w[numIn*numOut:]
	CheckSize(len(b), numBias)
	return b
}

func (f *MNIST1) Weight(w []float64, i int) []float64 {
	return w[i*numIn : (i+1)*numIn]
}
