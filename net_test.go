package vs

import (
	"math"
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestNet_Eval_Unit(t *testing.T) {
	// softmax * unit matrix
	net := NewNet(Unit(2))
	set(net.Params(), 1)

	y := MakeV(2)
	x := V{1, 0}
	net.Eval(y, x)
	exp := math.Exp
	test.Approxv(t, y, V{exp(1) / (exp(1) + exp(0)), exp(0) / (exp(1) + exp(0))}, 1e-12)
}

func TestNet_Backprop(t *testing.T) {
	testGrad(t, NewNet(Unit(5)))
	testGrad(t, NewNet(Unit(5), Unit(5)))
	testGrad(t, NewNet(LU(5, 2)))
	testGrad(t, NewNet(Unit(5), Unit(5), LU(5, 20), Unit(20)))
	testGrad(t, NewNet(LU(2, 5), Unit(5), LU(5, 20)))

	testGrad(t, NewNet(Re(LU(2, 5))))
	testGrad(t, NewNet(Re(LU(2, 5)), Re(LU(5, 20))))
	testGrad(t, NewNet(Re(LU(2, 5)), Re(LU(5, 5)), Re(LU(5, 20))))
}
