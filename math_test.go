package main

import (
	"reflect"
	"testing"
)

func TestMath(t *testing.T) {
	testEqf(t, Sum([]float32{1, 2, 3}), 6)
	testEqf(t, Dot([]float32{1, 2, 3}, []float32{4, 5, 6}), 32)

}

func testEqf(t *testing.T, have, want float32) {
	t.Helper()
	testEq(t, have, want)
}

func testEqv(t *testing.T, have, want []float32) {
	t.Helper()
	testEq(t, have, want)
}

func testEq(t *testing.T, have, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have: %v, want: %v", have, want)
	}
}
