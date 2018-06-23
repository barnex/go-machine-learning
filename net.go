package vs

type Net struct {
	Loss   float64
	LossF  OutFunc
	Layers []Layer
	W      V // all weights, shared over the layers
	X      V // net input, shared with the first layer
}

// NewNet constructs a feedforward chain.
// The layers are specified output-to-input:
// 	NewNet(output, top_hidden, ..., bottom_hidden, input)
func NewNet(out OutFunc, layers ...Func) *Net {
	net := &Net{}
	net.initLayers(layers)
	net.initParam()
	return net
}

func (n *Net) initLayers(funcs []Func) {
	n.Layers = make([]Layer, 0, len(funcs))
	for i := len(funcs) - 1; i >= 0; i-- {
		L := Layer{
			F: funcs[i],
		}
		n.Layers = append(n.Layers, L)
	}
}

// each layer gets a sub-slice of the net's overall parameters
func (n *Net) initParam() {
	n.W = MakeV(n.NumParam())
	off := 0
	for _, l := range n.Layers {
		l.W = n.W[off : off+l.NumParam()]
		off += l.NumParam()
	}
}

func (n *Net) NumParam() int {
	p := 0
	for _, l := range n.Layers {
		p += l.NumParam()
	}
	return p
}
