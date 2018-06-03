package vs

import (
	"fmt"
	"testing"
)

func TestMNIST1_Untrained(t *testing.T) {
	m := &MNIST1{}
	w := make([]float64, m.NumWeight())

	for i := 0; i < m.NumOut(); i++ {
		W := m.Weight(w, i)
		Randomize(W, 0.1)
	}

	fmt.Println(Accuracy(m, w, trainSet()), "/", len(trainSet()))
}
