package vs

import (
	"math"
)

type Net struct {
	f  []DiffFunc
	h  []V
	j  []M
	w  V // all weights, shared over the layers
	wl []V
}

// NewNet constructs a feedforward chain,
// with a final softmax layer (implicit).
func NewNet(layers ...DiffFunc) *Net {
	n := &Net{f: layers}

	// allocate parameters
	n.w = MakeV(n.NumParam())
	n.wl = n.sliceParams(n.w)

	// allocate hidden layers
	n.h = make([]V, len(n.f))
	for i, l := range n.f {
		n.h[i] = MakeV(l.NumOut())
	}

	return n
}

func (n *Net) Eval(y, x V) {
	n.f[0].Eval(n.h[0], n.wl[0], x) // first layer

	for i := 1; i < len(n.f); i++ { // other layers
		n.f[i].Eval(n.h[i], n.wl[i], n.h[i-1]) // previous layer's output is this layer's input
	}

	softmax(y, n.h[n.top()]) // final softmax
	/*
	 x -> f0 -> h[0]
	 h[0] -> f1 -> h[1]
	 h[1] -> f2 -> h[2]
	 h[2] -> softmax -> y
	*/
}

func (n *Net) Backprop(dy, y V, x V, c int) (loss float64) {

	// y = softmax(f3(w3, f2(w2, ...f0(w0, x))))
	n.Eval(y, x)

	// J = grad_x( -log( softmax( y )_c)) , y = h[top]
	J := MakeV(n.NumOut())
	copyv(J, y)
	J[c] -= 1

	dyl := n.sliceParams(dy)
	for i := n.top(); i > 0; i-- {
		f := n.f[i]
		wl := n.wl
		hl := n.h
		//dy := dyl[i]

		// jacobian: weights to outputs
		JW := MakeM(Dim2{f.NumParam(), f.NumOut()})
		f.DiffW(JW, wl[i], hl[i-1])

		// this layer's contribution to the gradient
		mulVM(dyl[i], J, JW)

		// Chain the layer below
		JX := MakeM(Dim2{f.NumIn(), f.NumOut()})
		f.DiffX(JX, wl[i], hl[i-1])
		J2 := MakeV(f.NumIn())
		mulVM(J2, J, JX)
		J = J2
	}

	{
		f := n.f[0]
		JW := MakeM(Dim2{f.NumParam(), f.NumOut()})
		f.DiffW(JW, n.wl[0], x)
		mulVM(dyl[0], J, JW)
	}

	return -math.Log(y[c])
}

func (n *Net) NumOut() int {
	return n.f[n.top()].NumOut()
}

func (n *Net) Params() V {
	return n.w
}

func (n *Net) NumParam() int {
	p := 0
	for _, l := range n.f {
		p += l.NumParam()
	}
	return p
}

func (n *Net) NumIn() int {
	return n.f[0].NumIn()
}

func (N *Net) sliceParams(w V) []V {
	wl := make([]V, len(N.f))
	off := 0
	for i, f := range N.f {
		wl[i] = w[off : off+f.NumParam()]
		off += f.NumParam()
	}
	assert(off == len(w)) // used exactly all parameters
	return wl
}

func (n *Net) top() int {
	return len(n.f) - 1
}
