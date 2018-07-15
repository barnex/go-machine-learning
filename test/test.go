// Package test provides test utilities.
package test

import (
	"math"
	"reflect"
	"testing"
)

func Eqf(t *testing.T, have, want float64) {
	t.Helper()
	Real(t, have, want)
	Eq(t, have, want)
}

func Eqv(t *testing.T, have, want []float64) {
	t.Helper()
	Real(t, have...)
	Real(t, want...)
	Eq(t, have, want)
}

func Approxf(t *testing.T, have, want, tol float64) {
	t.Helper()
	Real(t, have, want)
	if math.Abs(have-want) > tol {
		t.Errorf("have: %v, want: %v +/- %v", have, want, tol)
	}
}

func Approxv(t *testing.T, have, want []float64, tol float64) {
	t.Helper()

	if len(have) != len(want) {
		t.Errorf("vector size mismatch: have len: %v, want: %v", len(have), len(want))
	}

	if !IsApproxV(have, want, tol) {
		t.Errorf("have: %v, want: %v", have, want)
	}
}

func IsApproxV(have, want []float64, tol float64) bool {
	ok := true
	for i := range have {
		if math.Abs(have[i]-want[i]) > tol {
			ok = false
			break
		}
	}
	return ok
}

func Eq(t *testing.T, have, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have: %v, want: %v", have, want)
	}
}

// Real fails the test if v is not a real number
// (NaN or infinity).
func Real(t *testing.T, v ...float64) {
	t.Helper()
	for _, v := range v {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			t.Errorf("have: %v", v)
		}
	}
}

// Panic fails the test if f does not panic.
func Panic(t *testing.T, f func()) {
	defer func() {
		if p := recover(); p == nil {
			t.Errorf("want panic")
		}
	}()
	f()
}
