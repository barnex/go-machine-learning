package vs

// LU is a linear unit, performing an affine transform:
// 	y = w * x + b
// With parameters [w, b] (weights and biases).
type LU struct {
	nOut, nIn int
}

// NewLU constructs a linear unit
// with nOut outputs and nIn inputs.
func NewLU(nOut, nIn int) *LU {
	return &LU{nIn: nIn, nOut: nOut}
}

// Eval implements Func.
func (f *LU) Eval(y V, θ, x V) {
	MulMV(y, f.Weights(θ), x)
	Add(y, y, f.Biases(θ))
}

// DiffW implements Func.
func (f *LU) DiffW(dy M, θ, x V) {
	AssureM(dy, Dim2{f.NumParam(), f.NumOut()})
	for i := 0; i < dy.Rows(); i++ {
		dyi := dy.Row(i)
		Copy(f.Weights(dyi).Row(i), x)
		f.Biases(dyi)[i] = 1
	}
}

// DiffX implements Func.
func (f *LU) DiffX(dy M, θ, x V) {
	AssureM(dy, Dim2{f.NumIn(), f.NumOut()})
	for i := 0; i < dy.Rows(); i++ {
		Copy(dy.Row(i), f.Weights(θ).Row(i))
	}
}

func (f *LU) Weights(θ V) M {
	return Reshape(θ[:f.numW()], Dim2{f.nIn, f.nOut})
}

func (f *LU) Biases(w V) V {
	CheckSize(w.Len(), f.NumParam())
	return w[f.numW():]
}

func (f *LU) NumOut() int   { return f.nOut }
func (f *LU) NumIn() int    { return f.nIn }
func (f *LU) NumParam() int { return f.nOut*f.nIn + f.nOut }
func (f *LU) numW() int     { return f.nIn * f.nOut }
func (f *LU) numB() int     { return f.nOut }
