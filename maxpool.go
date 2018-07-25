package vs

type maxPool1D struct {
	nOut, width int
}

func MaxPool1D(nOut, width int) *maxPool1D {
	return &maxPool1D{nOut: nOut, width: width}
}

// Eval implements Func.
func (f *maxPool1D) Eval(y, _, x V) {
	checkSize(y.Len(), f.nOut)
	checkSize(x.Len(), f.NumIn())

	X := Reshape2(x, Dim2{f.width, f.NumOut()}) // width * numOut = numIn
	for i := range y {
		y[i] = maxv(X.Elem(i))
	}
}

// DiffW implements Func.
func (f *maxPool1D) DiffW(dy M, _, _, _ V) {
	assureM(dy, Dim2{f.NumParam(), f.NumOut()}) // == (0,0)
}

// DiffX implements Func.
func (f *maxPool1D) DiffX(dy M, y, θ, x V) {
	assureM(dy, Dim2{f.NumIn(), f.NumOut()})
	Set(dy.List, 0)

	X := Reshape2(x, Dim2{f.width, f.NumOut()})
	for i := 0; i < f.NumOut(); i++ {
		dyi := Reshape2(dy.Elem(i), X.Size)
		dyi.Elem(i)[argmax(X.Elem(i))] = 1
	}
}

// TODO: testDiff: randomize before eval, fail with separate message or so

func (f *maxPool1D) NumOut() int   { return f.nOut }
func (f *maxPool1D) NumIn() int    { return f.nOut * f.width }
func (f *maxPool1D) NumParam() int { return 0 }
