package vs

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
	n.wl = make([]V, len(n.f))
	off := 0
	for i, l := range n.f {
		n.wl[i] = n.w[off : off+l.NumParam()]
		off += l.NumParam()
	}
	Assert(off == len(n.w)) // used exactly all parameters

	// allocate hidden layers
	n.h = make([]V, len(n.f))
	for i, l := range n.f {
		n.h[i] = MakeV(l.NumOut())
	}

	return n
}

func (n *Net) Eval(y, x V) {
	// x -> f0 -> h[0]
	// h[0] -> f1 -> h[1]
	// h[1] -> f2 -> h[2]
	// h[2] -> softmax -> y

	// first layer
	n.f[0].Eval(n.h[0], n.wl[0], x)

	// other layers
	for i := 1; i < len(n.f); i++ {
		n.f[i].Eval(n.h[i], n.wl[i], n.h[i-1]) // previous layer's output is this layer's input
	}

	// final softmax
	SoftMax(y, n.h[n.top()])
}

func (n *Net) Backprop(dy, y V, x V, c int) (loss float64) {
	n.Eval(y, x)
	// TODO: dy

	// ∂f(w2, g(w1,x1)) / ∂w
	// = ( ∂f(w2, g(w1,x1))/∂w1 , ∂f(w2, g(w1,x1)) / ∂ w2 )
	//
	//  ∂f(w2, g(w1,x1))/∂w1
	//  = (∂f/∂g)*(∂g/w1)
	return 0
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

func (n *Net) top() int {
	return len(n.f) - 1
}
