package vs

import (
	"fmt"
	"math"
)

type Net struct {
	f  []DiffFunc
	h  []V // buffer for intermediate results
	j  []M // buffer for backprop jacobians
	w  V   // all weights, shared over the layers
	wl []V // weights per layer
}

// NewNet constructs a feedforward chain, with a final softmax function.
// The layers are specified from output to input.
// E.g.:
// 	NewNet(LU(2, 5), ReLU(5, 20))
// Applies LU(ReLU(x)), has 2 outputs and 20 inputs.
func NewNet(layers ...DiffFunc) *Net {
	// reverse layers
	l := make([]DiffFunc, len(layers))
	for i := range l {
		l[i] = layers[len(l)-1-i]
	}
	checkLayers(l)
	n := &Net{f: l}

	// allocate parameters
	numP := 0
	for _, l := range n.f {
		numP += l.NumParam()
	}
	n.w = MakeV(numP)
	n.wl = n.sliceParams(n.w)

	// allocate hidden layers
	n.h = make([]V, len(n.f))
	for i, l := range n.f {
		n.h[i] = MakeV(l.NumOut())
	}

	return n
}

func checkLayers(l []DiffFunc) {
	for i := 1; i < len(l); i++ {
		if l[i].NumIn() != l[i-1].NumOut() {
			panic(fmt.Sprintf("net size mismatch: layer %v has %v inputs, but layer %v has %v outputs", i, l[i].NumIn(), i-1, l[i-1].NumOut()))
		}
	}
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
	n.Eval(y, x)

	// softmax jacobian J = grad_x( loss )
	J := MakeV(n.NumOut())
	copyv(J, y)
	J[c] -= 1

	// chain rule
	dyl := n.sliceParams(dy)
	for i := n.top(); i > 0; i-- {
		f := n.f[i]
		wl := n.wl
		hl := n.h

		// JW: weights to outputs
		JW := MakeM(Dim2{f.NumParam(), f.NumOut()})
		f.DiffW(JW, hl[i], wl[i], hl[i-1])
		// this layer's contribution to the gradient
		mulVM(dyl[i], J, JW)

		// JX: inputs to outputs
		JX := MakeM(Dim2{f.NumIn(), f.NumOut()})
		f.DiffX(JX, hl[i], wl[i], hl[i-1])
		J2 := MakeV(f.NumIn())
		mulVM(J2, J, JX)
		J = J2
	}

	// input layer
	{
		f := n.f[0]
		JW := MakeM(Dim2{f.NumParam(), f.NumOut()})
		f.DiffW(JW, n.h[0], n.wl[0], x)
		mulVM(dyl[0], J, JW)
	}

	return -math.Log(y[c])
}

func (n *Net) NumOut() int         { return n.f[n.top()].NumOut() }
func (n *Net) NumIn() int          { return n.f[0].NumIn() }
func (n *Net) NumParam() int       { return len(n.w) }
func (n *Net) Params() V           { return n.w }
func (n *Net) LParams(layer int) V { return n.wl[layer] }

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
