package vs

// A func is a differentiable, parametrized vector-to-vector function.
// Funcs are chained to together to form a differentiable network.
type Func interface {
	NumOut() int   // NumOut returns the number of outputs.
	NumParam() int // NumParam returns the number of parameters.
	NumIn() int    // NumIn retruns the number of inputs.

	// Eval evaluates the function:
	// 	*y = f(w,x)
	// If y==nil, it is allocated. Otherwise the size must fit.
	Eval(y *V, w, x V)

	// DiffW computes the derivates (Jacobian matrix) with respect to w:
	// 	dy[i][j] = (∂f(w,x)[i] / ∂w[j])
	// If dy==nil, it is allocated. Otherwise the size must fit.
	DiffW(dy *M, w, x V)

	// DiffW computes the derivates (Jacobian matrix) with respect to x:
	// 	dy[i][j] = (∂f(w,x)[i] / ∂x[j])
	// If dy==nil, it is allocated. Otherwise the size must fit.
	DiffX(dy *M, w, x V)
}

// An OutFunc is the final output layer
// TODO: specify
type OutFunc interface {
}
