package vs

import (
	"testing"
)

//func TestNet_Eval_Unit(t *testing.T) {
//	net := NewNet()
//	//set(net.Params(), 1)
//
//	y := MakeV(2)
//	x := V{1, 0}
//	net.Eval(y, x)
//	exp := math.Exp
//	test.Approxv(t, y, V{exp(1) / (exp(1) + exp(0)), exp(0) / (exp(1) + exp(0))}, 1e-12)
//}

func TestNet_Backprop_Unit(t *testing.T) {
	net := NewNet(Unit(2))
	testGrad(t, net, 0)
	//testGrad(t, net, 1)
}

func TestNet_Backprop_LU(t *testing.T) {
	t.Skip("TODO")
	net := NewNet(NewLU(5, 2))
	testGrad(t, net, 0)
}
