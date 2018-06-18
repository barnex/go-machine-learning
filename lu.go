package vs

// LU is a linear unit, performing an affine transform:
// 	y = a * x + b
type LU struct {
	nOut, nIn int
}

// NewLU constructs a linear unit
// with nOut outputs and nIn inputs.
func NewLU(nOut, nIn int) *LU {
	return &LU{nIn: nIn, nOut: nOut}
}

func (f *LU) NumOut() int   { return f.nOut }
func (f *LU) NumIn() int    { return f.nIn }
func (f *LU) NumParam() int { return f.nOut*f.nIn + f.nOut }
func (f *LU) numW() int     { return f.nIn * f.nOut }
func (f *LU) numB() int     { return f.nOut }

func (f *LU) Weights(theta V) M {
	return Reshape(theta[:f.numW()], Dim2{f.nIn, f.nOut})
}

func (f *LU) Biases(w V) V {
	CheckSize(w.Len(), f.NumParam())
	return w[f.numW():]
}

func (f *LU) Eval(y *V, theta, x V) {
	MulMV(y, f.Weights(theta), x)
	Add(y, *y, f.Biases(theta))
}

func (f *LU) Diff(y *M, theta, x V) {
	AssureM(y, diffSize(f))
	for j := 0; j < y.Size[1]; j++ {
		g := y.Row(j)
		w := f.Weights(g)
		Copy(w.Row(j), x)
		f.Biases(g)[j] = 1
	}
}
