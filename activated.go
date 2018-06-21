package vs

//type Activated struct {
//	F          Func
//	Activation ScalarFunc
//}
//
//var _ Func = (*Activated)(nil)
//
//func (f *Activated) NumOut() int   { return f.F.NumOut() }
//func (f *Activated) NumParam() int { return f.F.NumParam() }
//func (f *Activated) NumIn() int    { return f.F.NumIn() }
//
//func (f *Activated) Eval(y *V, w, x V) {
//	f.F.Eval(y, w, x)
//	Map(*y, *y, f.Activation.Eval)
//}
//
//func (f *Activated) DiffW(dy *M, w, x V) {
//	var y M
//	f.F.Eval(&y, w, x)
//	f.F.DiffW(dy, w, x)
//
//	for i := 0; i < dy.Rows(); i++ {
//		dyi := dy.Row(i)
//
//	}
//}
//
//func (f *Activated) DiffX(dy *M, w, x V) {
//	f.F.DiffX(dy, w, x)
//	Map(dy.List, dy.List, f.Activation.Diff)
//}
//
//func NewReLU(nOut, nIn int) *Activated {
//	return &Activated{F: NewLU(nOut, nIn), Activation: Re}
//}
