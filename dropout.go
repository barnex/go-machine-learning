package vs

import (
	"math/rand"
)

type dropout struct {
	n, m int
	frac float64
	rng  rand.Rand
	drop []bool
}

func Dropout(n, m int, frac float64) *dropout {
	return &dropout{
		n:    n,
		m:    m,
		frac: frac,
		rng:  *rand.New(rand.NewSource(1)),
		drop: make([]bool, n),
	}
}

// Eval implements Func.
func (f *dropout) Eval(y, _, x V) {
	copyv(y, x)
	for i := range f.drop {
		if f.drop[i] {
			y[i] = -1 // TODO: which value?
		}
	}
}

func (f *dropout) NextState() {
	f.Disable()
	for i := 0; i < f.n; i += f.m {
		if f.rng.Float64() < f.frac {
			f.drop[i+f.rng.Intn(f.m)] = true
		}
	}
}

func (f *dropout) Disable() {
	for i := range f.drop {
		f.drop[i] = false
	}
}

// DiffW implements Func.
func (f *dropout) DiffW(dy M, _, _, _ V) {
	assureM(dy, Dim2{f.NumParam(), f.NumOut()}) // == (0,0)
}

// DiffX implements Func.
func (f *dropout) DiffX(dy M, y, θ, x V) {
	assureM(dy, Dim2{f.NumIn(), f.NumOut()})
	checkSize(y.Len(), f.n)
	checkSize(x.Len(), f.n)
	checkSize(θ.Len(), 0)

	for i := range y {
		dyi := dy.Elem(i)
		Set(dyi, 0)
		if !f.drop[i] {
			dyi[i] = 1
		}
	}
}

func (f *dropout) NumOut() int   { return f.n }
func (f *dropout) NumIn() int    { return f.n }
func (f *dropout) NumParam() int { return 0 }
