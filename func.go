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

// A DiffFunc is parametrized, differentiable vector-to-vector function.
type DiffFunc interface {
	Func

	// DiffW computes the derivates (Jacobian matrix) with respect to w:
	// 	dy[i][j] = (∂f(w,x)[i] / ∂w[j])
	DiffW(dy M, y, w, x V)

	// DiffW computes the derivates (Jacobian matrix) with respect to x:
	// 	dy[i][j] = (∂f(w,x)[i] / ∂x[j])
	DiffX(dy M, y, w, x V)
}

func diffWSize(f Func) Dim2 {
	return Dim2{f.NumParam(), f.NumOut()}
}

func diffXSize(f Func) Dim2 {
	return Dim2{f.NumIn(), f.NumOut()}
}

// An OutFunc is the final output layer
// TODO: specify
//type OutFunc interface {
//	NumOut() int // NumOut returns the number of outputs.
//	Infer(y V, x V)
//}
