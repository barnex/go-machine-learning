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
	Eq(t, len(have), len(want))
	for i := range have {
		Approxf(t, have[i], want[i], tol)
	}
}

func Eq(t *testing.T, have, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have: %v, want: %v", have, want)
	}
}

func Real(t *testing.T, v ...float64) {
	t.Helper()
	for _, v := range v {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			t.Errorf("have: %v", v)
		}
	}
}
