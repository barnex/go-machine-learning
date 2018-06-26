package vs

type unit struct {
	n int
}

func Unit(n int) DiffFunc {
	return &unit{n: n}
}

func (f *unit) Eval(y V, w, x V) {
	assureV(w, f.n)
	assureV(x, f.n)
	for i := range y {
		y[i] = w[i] * x[i]
	}
}

func (f *unit) DiffW(dy M, y, w, x V) {
	assureM(dy, Dim2{f.NumParam(), f.NumOut()})
	set(dy.List, 0)
	for i := 0; i < dy.Rows(); i++ {
		dy.Row(i)[i] = x[i]
	}
}

func (f *unit) DiffX(dy M, y, w, x V) {
	assureM(dy, Dim2{f.NumIn(), f.NumOut()})
	set(dy.List, 0)
	for i := 0; i < dy.Rows(); i++ {
		dy.Row(i)[i] = w[i]
	}
}

func (f *unit) NumOut() int   { return f.n }
func (f *unit) NumIn() int    { return f.n }
func (f *unit) NumParam() int { return f.n }
