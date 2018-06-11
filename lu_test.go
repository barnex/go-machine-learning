package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

// A trivial learning test: learn the identity function (3x3).
func TestLU_Train_Identity(t *testing.T) {
	size := Dim{3, 3}

	trainSet := []LabeledVec{
		{0, []float64{1, 0, 0}},
		{1, []float64{0, 1, 0}},
		{2, []float64{0, 0, 1}},
	}

	f := NewLU(size[0], size[1])

	w := make([]float64, f.NumWeight())
	Set(w, 1)

	GradDescent(f, w, trainSet, 100)

	infer := make([]float64, f.NumOut())
	Infer(infer, f, w, trainSet[0].X)
	test.Approxv(t, infer, []float64{1, 0, 0}, 1e-30)

	Infer(infer, f, w, trainSet[1].X)
	test.Approxv(t, infer, []float64{0, 1, 0}, 1e-30)

	Infer(infer, f, w, trainSet[2].X)
	test.Approxv(t, infer, []float64{0, 0, 1}, 1e-30)

	test.Eq(t, Accuracy(f, w, trainSet), 3)
}

// Test Gradient of Loss by comparing to numerical approximation.
func TestLU_GradLoss(t *testing.T) {
	sizes := []Dim{
		{2, 2},
		{2, 5},
		{5, 2},
	}

	for _, s := range sizes {
		t.Run(s.String(), func(t *testing.T) {
			f := NewLU(s[0], s[1])

			w := make([]float64, f.NumWeight())
			Randomize(w, 1)
			x := make([]float64, f.NumIn())
			Randomize(x, 1)
			xl := LabeledVec{1, x}

			have := make([]float64, f.NumOut())
			GradLoss(have, f, w, xl, T{})

			want := make([]float64, f.NumOut())
			NumGradLoss(want, f, w, xl)

			test.Approxv(t, have, want, 1e-5)
		})
	}
}

// Test raw gradient by comparing to numerical approximation
func TestLU_Grad(t *testing.T) {

	sizes := []Dim{
		{2, 2},
		{2, 5},
		{5, 2},
		{1, 1},
	}

	for _, s := range sizes {
		t.Run(s.String(), func(t *testing.T) {
			f := NewLU(s[0], s[1])

			w := make([]float64, f.NumWeight())
			Randomize(w, 1)
			x := make([]float64, f.NumIn())
			Randomize(x, 1)

			have := MakeT(gradSize(f))
			f.Grad(have, w, x)

			want := MakeT(gradSize(f))
			NumGrad(want, f, w, x)

			test.Approxv(t, have.List(), want.List(), 1e-5)
		})
	}

}

func TestLU_Eval(t *testing.T) {
	f := NewLU(2, 4)

	w := make([]float64, 2*4+2)

	// A = [1 2 3 4]
	//     [5 6 7 8]
	// B = [9 10]

	f.Ai(w, 0)[0] = 1
	f.Ai(w, 0)[1] = 2
	f.Ai(w, 0)[2] = 3
	f.Ai(w, 0)[3] = 4

	f.Ai(w, 1)[0] = 5
	f.Ai(w, 1)[1] = 6
	f.Ai(w, 1)[2] = 7
	f.Ai(w, 1)[3] = 8

	f.B(w)[0] = 9
	f.B(w)[1] = 10

	tests := []struct {
		x    []float64
		want []float64
	}{
		{[]float64{1, 0, 0, 0}, []float64{1 + 9, 5 + 10}},
		{[]float64{2, 0, 0, 0}, []float64{2 + 9, 10 + 10}},
		{[]float64{0, 0, 0, 1}, []float64{4 + 9, 8 + 10}},
		{[]float64{0, 1, 1, 0}, []float64{2 + 3 + 9, 6 + 7 + 10}},
		{[]float64{0, 0, 0, 0}, []float64{9, 10}},
	}

	for _, c := range tests {
		have := make([]float64, f.NumOut())
		f.Eval(have, w, c.x)
		test.Eqv(t, have, c.want)
	}
}
