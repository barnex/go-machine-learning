package vs

type Layer struct {
	Y V // Output
	W V // Parameters
	X V // Input
	J M // Jacobian matrix
	F Func
}

func (l *Layer) UpdateY() {
	l.F.Eval(&l.Y, l.W, l.X)
}

func (l *Layer) NumParam() int { return l.F.NumParam() }
