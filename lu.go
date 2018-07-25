package vs

// lu is a linear unit, performing an affine transform:
// 	y = w * x + b
// With parameters [w, b] (weights and biases).
type lu struct {
	nOut, nIn int
}

// LU constructs a linear unit
// with nOut outputs and nIn inputs.
func LU(nOut, nIn int) *lu {
	return &lu{nIn: nIn, nOut: nOut}
}

// Eval implements Func.
func (f *lu) Eval(y V, θ, x V) {
	mulMV(y, f.Weights(θ), x)
	add(y, y, f.Biases(θ))
}

// DiffW implements Func.
func (f *lu) DiffW(dy M, y, θ, x V) {
	assureM(dy, Dim2{f.NumParam(), f.NumOut()})
	for i := 0; i < dy.NumElem(); i++ {
		dyi := dy.Elem(i)
		Set(dyi, 0)
		copyv(f.Weights(dyi).Elem(i), x)
		f.Biases(dyi)[i] = 1
	}
}

// DiffX implements Func.
func (f *lu) DiffX(dy M, y, θ, x V) {
	assureM(dy, Dim2{f.NumIn(), f.NumOut()})
	for i := 0; i < dy.NumElem(); i++ {
		copyv(dy.Elem(i), f.Weights(θ).Elem(i))
	}
}

func (f *lu) Weights(θ V) M {
	return Reshape2(θ[:f.numW()], Dim2{f.nIn, f.nOut})
}

func (f *lu) Biases(w V) V {
	checkSize(w.Len(), f.NumParam())
	return w[f.numW():]
}

func (f *lu) NumOut() int   { return f.nOut }
func (f *lu) NumIn() int    { return f.nIn }
func (f *lu) NumParam() int { return f.nOut*f.nIn + f.nOut }
func (f *lu) numW() int     { return f.nIn * f.nOut }
func (f *lu) numB() int     { return f.nOut }
