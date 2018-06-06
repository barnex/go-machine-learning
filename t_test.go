package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

func TestT(t *testing.T) {
	T := MakeT(3, 2)

	list := []float64{1, 2, 3, 4, 5, 6}

	Copy(T.List(), list)

	test.Eqv(t, T.Row(0), []float64{1, 2, 3})
	test.Eqv(t, T.Row(1), []float64{4, 5, 6})

}
