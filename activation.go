package vs

// withActivation wraps a function with an element-wise activation function. E.g.:
type withActivation struct {
	F              DiffFunc
	Activation     func(float64) float64
	DiffActivation func(float64) float64
}

// Re wraps a function with a linear rectifier activation:
// 	Re(f(x)) = max(0, f(x))
// E.g.:
// 	Re(LU(1,2))
// yields a ReLU unit.
func Re(f DiffFunc) DiffFunc {
	return &withActivation{F: f, Activation: re, DiffActivation: step}
}

func LeakyRe(f DiffFunc, leak float64) DiffFunc {
	act := func(x float64) float64 {
		if x > 0 {
			return x
		}
		return leak * x
	}
	diff := func(x float64) float64 {
		if x > 0 {
			return 1
		}
		return leak
	}
	return &withActivation{F: f, Activation: act, DiffActivation: diff}
}

func (f *withActivation) NumOut() int   { return f.F.NumOut() }
func (f *withActivation) NumParam() int { return f.F.NumParam() }
func (f *withActivation) NumIn() int    { return f.F.NumIn() }

func (f *withActivation) Eval(y V, w, x V) {
	f.F.Eval(y, w, x)
	mapf(y, y, f.Activation)
}

func (f *withActivation) DiffX(dy M, y, w, x V) {
	assureM(dy, Dim2{f.NumIn(), f.NumOut()})
	f.F.DiffX(dy, y, w, x)

	for i := 0; i < dy.Size[1]; i++ {
		dyi := dy.Elem(i)
		for j := 0; j < dy.Size[0]; j++ {
			dyi[j] *= f.DiffActivation(y[i])
		}
	}
}

func (f *withActivation) DiffW(dy M, y, w, x V) {
	assureM(dy, Dim2{f.NumParam(), f.NumOut()})
	f.F.DiffW(dy, y, w, x)

	for i := 0; i < dy.Size[1]; i++ {
		dyi := dy.Elem(i)
		for j := 0; j < dy.Size[0]; j++ {
			dyi[j] *= f.DiffActivation(y[i])
		}
	}
}
