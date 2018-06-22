package vs

import "testing"

func TestChain(t *testing.T) {
	//L0 := Layer{
	//	X: MakeV(20),
	//	W: MakeV(10*20 + 10),
	//	Y: MakeV(10),
	//	F: NewLU(10, 20),
	//}

	//L1 := Layer{
	//	X: L0.Y,
	//	W: MakeV(5*10 + 5),
	//	Y: MakeV(5),
	//	F: NewLU(5, 10),
	//}

	// ∂f(w2, g(w1,x1)) / ∂w
	// = ( ∂f(w2, g(w1,x1))/∂w1 , ∂f(w2, g(w1,x1)) / ∂ w2 )
	//
	//  ∂f(w2, g(w1,x1))/∂w1
	//  = (∂f/∂x2)*(∂g/w1)
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
