package vs

// LU is a linear unit, performing:
// 	y = a * x + b
type LU struct {
	numOut, numIn int
}

func NewLU(numOut, numIn int) *LU {
	return &LU{numIn: numIn, numOut: numOut}
}

var _ Net = (*LU)(nil)

func (f *LU) NumOut() int {
	return f.numOut
}

func (f *LU) NumWeight() int {
	return f.numOut*f.numIn + f.numOut
}

func (f *LU) NumIn() int {
	return f.numIn
}

func (f *LU) Ai(w []float64, i int) []float64 {
	return w[i*f.numIn : (i+1)*f.numIn]
}

func (f *LU) B(w []float64) []float64 {
	CheckSize(len(w), f.NumWeight())
	b := w[(f.numIn)*(f.numOut):]
	CheckSize(len(b), f.numOut)
	return b
}

func (f *LU) Eval(y, w, x []float64) {
	NetCheckSize(f, y, w, x)

	B := f.B(w)
	for i := range y {
		Ai := f.Ai(w, i)
		y[i] = Dot(Ai, x) + B[i]
	}
}

func (f *LU) Grad(y T, w, x []float64) {

}
