package vs

import (
	"log"
	"math"
	"testing"
)

func TestModel0(t *testing.T) {
	model := NewModel0()

	Train(model, LoadLabeledSet("mnist_png/training"))

	testSet := LoadLabeledSet("mnist_png/testing")

	correct := Test(model, testSet)
	success := float64(correct) / float64(len(testSet))
	log.Println("have:", success)
	want := 0.1
	if !approx(success, want, 0.03) {
		t.Errorf("success: have: %v, want: %v", success, want)
	}
}

func approx(have, want, tol float64) bool {
	return math.Abs(have-want) < tol
}
