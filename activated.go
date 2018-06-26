package vs

type Activated struct {
	F          Func
	Activation ScalarFunc
}

var _ Func = (*Activated)(nil)

func (f *Activated) NumOut() int   { return f.F.NumOut() }
func (f *Activated) NumParam() int { return f.F.NumParam() }
func (f *Activated) NumIn() int    { return f.F.NumIn() }

func (f *Activated) Eval(y V, w, x V) {
	f.F.Eval(y, w, x)
	mapf(y, y, f.Activation.Eval)
}

func (f *Activated) DiffW(dy M, w, x V) {
	//	f.F.DiffW(dy, w, x)
	//
	//	for i := 0; i < dy.Size[1]; i++ {
	//		dyi := dy.Row(i)
	//		for j := 0; j < dy.Size[0]; j++ {
	//			dyi[j] *= f.Activation.Diff(y[i])
	//		}
	//	}
}

func (f *Activated) DiffX(dy *M, w, x V) {

}

func ReLU(nOut, nIn int) *Activated {
	return &Activated{F: LU(nOut, nIn), Activation: Re}
}
