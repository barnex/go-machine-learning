package vs

import "testing"

func TestChain(t *testing.T) {
	net := NewNet(NewSoftmaxXEntropy(10), NewLU(10, 36), NewActivation(36, Re), NewLU(36, 36))
	Randomize(net.W, 0.1)
	// ∂f(w2, g(w1,x1)) / ∂w
	// = ( ∂f(w2, g(w1,x1))/∂w1 , ∂f(w2, g(w1,x1)) / ∂ w2 )
	//
	//  ∂f(w2, g(w1,x1))/∂w1
	//  = (∂f/∂g)*(∂g/w1)
}

type OutLayer struct {
	// TODO
}

func Forward(w, x, out OutLayer, net []Layer) V {

	// Outlayer has *SoftmaxXEntropy-like func, output vector and loss scalar
	// Forward updates output and loss

	return V{}
}

func Backprop(w, x, out OutLayer, net []Layer) V {
	// 0: alloc and slice gradient
	// 1: slice weights: theta[layer]
	// 2: Func.Diff(dy, y, ...)
	// 3:

	return V{}
}
