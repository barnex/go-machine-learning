package vs

import (
	"math"
	"reflect"
	"testing"
)

func testEqf(t *testing.T, have, want float64) {
	t.Helper()
	testEq(t, have, want)
}

func testApprox(t *testing.T, have, want, tol float64) {
	t.Helper()
	if math.Abs(have-want) > tol {
		t.Errorf("have: %v, want: %v +/- %v", have, want, tol)
	}
}

func testEqv(t *testing.T, have, want []float64) {
	t.Helper()
	testEq(t, have, want)
}

func testEq(t *testing.T, have, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have: %v, want: %v", have, want)
	}
}
