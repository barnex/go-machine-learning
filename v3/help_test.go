package vs

/*

import (
	"math"
	"reflect"
	"testing"
)

var (
	// TODO: load all
	trainingSet []LabeledImg = LoadLabeledSet("mnist_png/training", 5)
	testingSet  []LabeledImg = LoadLabeledSet("mnist_png/testing", 200)
)

func testEqf(t *testing.T, have, want float64) {
	t.Helper()
	testNoNaN(t, have, want)
	testEq(t, have, want)
}

func testApprox(t *testing.T, have, want, tol float64) {
	t.Helper()
	testNoNaN(t, have, want)
	if math.Abs(have-want) > tol {
		t.Errorf("have: %v, want: %v +/- %v", have, want, tol)
	}
}

func testEqv(t *testing.T, have, want []float64) {
	t.Helper()
	testNoNaN(t, have...)
	testNoNaN(t, want...)
	testEq(t, have, want)
}

func testEq(t *testing.T, have, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have: %v, want: %v", have, want)
	}
}

func testNoNaN(t *testing.T, v ...float64) {
	t.Helper()
	for _, v := range v {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			t.Errorf("have: %v", v)
		}
	}
}
*/
