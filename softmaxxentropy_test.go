package vs

// Test raw derivatives by comparing to numerical ≈imation
//func TestSoftmaxXEntropy_Diff(t *testing.T) {
//	f := NewSoftmaxXEntropy(10)
//	f2 := f.FixLabel(3) // fix a random label
//
//	x := make([]float64, f.NumIn())
//	Randomize(x, 1)
//
//	var have M
//	f2.DiffX(&have, nil, x)
//
//	var want M
//	NumericDiffX(&want, f2, nil, x)
//
//	test.Approxv(t, have.List, want.List, 1e-5)
//}
