package vs

import "testing"

func TestUnit_DiffW(t *testing.T) {
	testDiffW(t, Unit(2))
}

func TestUnit_DiffX(t *testing.T) {
	testDiffX(t, Unit(2))
}
