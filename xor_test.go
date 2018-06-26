package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

// Test feedforward of an XOR network.
// deeplearningbook.org p. 171, adapted to use softmax.
func TestXOR(t *testing.T) {
	l1 := LU(2, 2)
	lu0 := LU(2, 2)
	l0 := Re(lu0)
	net := NewNet(l1, l0)

	// pre-trained weights from deeplearningbook.org p.171
	copyv(lu0.Weights(net.wl[0]).List, V{1, 1, 1, 1})
	copyv(lu0.Biases(net.wl[0]), V{0, -1})
	copyv(l1.Weights(net.wl[1]).List, V{1, -2, -1, 2})
	copyv(l1.Biases(net.wl[1]), V{0, 1})

	cases := []struct {
		x, want V
	}{
		{V{0, 0}, V{0, 1}},
		{V{0, 1}, V{1, 0}},
		{V{1, 0}, V{1, 0}},
		{V{1, 1}, V{0, 1}},
	}

	y := MakeV(net.NumOut())
	for _, c := range cases {
		net.Eval(y, c.x)
		want := MakeV(c.want.Len())
		softmax(want, c.want)
		test.Eqv(t, y, want)
	}

	testGrad(t, net)
}
