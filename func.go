package vs

// A func is a parametrized vector-to-vector function.
type Func interface {
	NumOut() int   // NumOut returns the number of outputs.
	NumParam() int // NumParam returns the number of parameters.
	NumIn() int    // NumIn retruns the number of inputs.

	// Eval evaluates the function: y = f(w,x)
	// y's size must be equal to NumOut()
	Eval(y V, w, x V)
}

type DiffFunc interface {
	Func

	// DiffW computes the derivates (Jacobian matrix) with respect to w:
	// 	dy[i][j] = (∂f(w,x)[i] / ∂w[j])
	DiffW(dy M, w, x V)

	// DiffW computes the derivates (Jacobian matrix) with respect to x:
	// 	dy[i][j] = (∂f(w,x)[i] / ∂x[j])
	DiffX(dy M, w, x V)
}

func diffWSize(f Func) Dim2 {
	return Dim2{f.NumParam(), f.NumOut()}
}

func diffXSize(f Func) Dim2 {
	return Dim2{f.NumIn(), f.NumOut()}
}

type LossFunc interface {
	NumIn() int
	Eval(x V, c int) float64
}

// TODO: move to numeric, then hide inside diffLossTheta
func FixLabel(f LossFunc, c int) Func {
	return MakeFunc(1, 0, f.NumIn(), func(y V, w, x V) {
		AssureV(y, 1)
		y[0] = f.Eval(x, c)
	})
}

func MakeFunc(nOut, nParam, nIn int, f func(V, V, V)) Func {
	return &makeFunc{nOut, nParam, nIn, f}
}

type makeFunc struct {
	nOut, nParam, nIn int
	f                 func(V, V, V)
}

var _ Func = (*makeFunc)(nil)

func (f *makeFunc) NumOut() int        { return f.nOut }
func (f *makeFunc) NumParam() int      { return f.nParam }
func (f *makeFunc) NumIn() int         { return f.nIn }
func (f *makeFunc) Eval(y V, w V, x V) { f.f(y, w, x) }

// An OutFunc is the final output layer
// TODO: specify
//type OutFunc interface {
//	NumOut() int // NumOut returns the number of outputs.
//	Infer(y V, x V)
//}
