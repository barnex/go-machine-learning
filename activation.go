package vs

// Activation is a layer that does point-wise activation only.
// Intended for testing.
type Activation struct {
	nIn int
	f   ScalarFunc
}

func NewActivation(nIn int, f ScalarFunc) *Activation {
	return &Activation{nIn: nIn, f: f}
}

// Eval implements Func.
func (f *Activation) Eval(y V, θ, x V) {
	AssureV(y, len(x))
	Map(y, x, f.f.Eval)
}

// DiffW implements Func.
func (f *Activation) DiffW(dy M, θ, x V) {
	AssureM(dy, Dim2{0, f.NumOut()})
}

// DiffX implements Func.
// The result is a diagonal matrix,
// thus very inefficient to multiply in a dense way.
func (f *Activation) DiffX(dy M, θ, x V) {

	y := MakeV(f.NumOut())
	f.Eval(y, θ, x) // TODO: remove

	AssureM(dy, Dim2{f.nIn, f.nIn})
	Set(dy.List, 0)
	for i := 0; i < f.nIn; i++ {
		dy.Row(i)[i] = f.f.Diff(y[i]) // TODO: Func.Diff: pass y for re-use
	}
}

func (f *Activation) NumOut() int   { return f.nIn }
func (f *Activation) NumIn() int    { return f.nIn }
func (f *Activation) NumParam() int { return 0 }
