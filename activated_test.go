package vs

//// Test raw derivatives by comparing to numerical ≈imation
//func TestReLU_DiffW(t *testing.T) {
//	for _, s := range stdSizes {
//		t.Run(s.String(), func(t *testing.T) {
//			testDiffW(t, NewReLU(s[0], s[1]))
//		})
//	}
//}
//
//// Test raw derivatives by comparing to numerical ≈imation
//func TestReLU_DiffX(t *testing.T) {
//	for _, s := range stdSizes {
//		t.Run(s.String(), func(t *testing.T) {
//			testDiffX(t, NewReLU(s[0], s[1]))
//		})
//	}
//}
//
//func TestReLU_Eval(t *testing.T) {
//	f := NewLU(2, 4)
//	θ := MakeV(2*4 + 2)
//
//	// A = [1 2 3 4]
//	//     [5 6 7 8]
//	// B = [9 10]
//
//	Copy(f.Weights(θ).List, V{1, 2, 3, 4, 5, 6, 7, 8})
//	Copy(f.Biases(θ), V{9, 10})
//
//	tests := []struct {
//		x    V
//		want V
//	}{
//		{V{1, 0, 0, 0}, V{1 + 9, 5 + 10}},
//		{V{2, 0, 0, 0}, V{2 + 9, 10 + 10}},
//		{V{0, 0, 0, 1}, V{4 + 9, 8 + 10}},
//		{V{0, 1, 1, 0}, V{2 + 3 + 9, 6 + 7 + 10}},
//		{V{0, 0, 0, 0}, V{9, 10}},
//	}
//
//	for _, c := range tests {
//		var have V
//		f.Eval(&have, θ, c.x)
//		test.Eqv(t, have, c.want)
//	}
//}
