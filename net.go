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
	net := &Net{
		f: layers,
	}
	net.initParam()
	net.initH()
	return net
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
	return n.topLayer().NumOut()
}

func (n *Net) Params() V {
	return n.w
}

func (n *Net) numParam() int {
	p := 0
	for _, l := range n.f {
		p += l.NumParam()
	}
	return p
}

func (n *Net) top() int {
	return len(n.f) - 1
}

func (n *Net) topLayer() Func {
	return n.f[n.top()]
}

// each layer gets a sub-slice of the net's overall parameters
func (n *Net) initParam() {
	n.w = MakeV(n.numParam())
	n.wl = make([]V, len(n.f))
	off := 0
	for i, l := range n.f {
		n.wl[i] = n.w[off : off+l.NumParam()]
		off += l.NumParam()
	}
	Assert(off == len(n.w)) // used exactly all parameters
}

func (n *Net) initH() {
	n.h = make([]V, len(n.f))
	for i, l := range n.f {
		n.h[i] = MakeV(l.NumOut())
	}
}

//func(n*Net) Validate(y V, x V, c int) (loss float64){
//	panic(0)
//}

//func (n *Net) Loss(x V, c int) float64{
//	return n.LossF.Loss(
//}

//func (n *Net) Eval(x V) {
//	Copy(n.X, x)
//	for _, l := range n.Layers {
//		l.Update()
//	}
//
//	lastHidden := n.Layers[len(n.Layers)-1]
//	n.LossF.Infer(n.Y, lastHidden.Y)
//	//n.Loss = n.LossF.Loss()
//}
//
//func (n *Net) initLayers(out *SoftmaxXEntropy, funcs []Func) {
//	n.LossF = out
//	n.Layers = make([]*Layer, 0, len(funcs))
//	for i := len(funcs) - 1; i >= 0; i-- {
//		L := &Layer{
//			F: funcs[i],
//		}
//		n.Layers = append(n.Layers, L)
//	}
//}
//
//
//func (n *Net) initXY() {
//	// alloc inputs
//	for _, l := range n.Layers {
//		l.X = MakeV(l.F.NumIn())
//	}
//	n.X = n.Layers[0].X
//
//	// output = input of layer above
//	for i := 0; i < len(n.Layers)-1; i++ {
//		n.Layers[i].Y = n.Layers[i+1].X
//	}
//	// last layer
//	last := n.Layers[len(n.Layers)-1]
//	last.Y = MakeV(last.F.NumOut())
//	n.Y = MakeV(last.Y.Len()) // TODO: softmaxxentropy wrongly has numout==1
//}
