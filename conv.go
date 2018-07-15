package vs

/*
TODO:
Dim2 constructor
Dim1 constructor
Dim3 type+constructor
NumIn, NumOut -> Dim3??
*/

// conv is a convolutional unit, performing a batch of cross-correlations:
// 	y[h][i][j] = Sum_kl x[i+k][j+l] * w[h][k][l]
type conv struct {
	outS  Dim3
	kernS Dim3
	inS   Dim2
}

func Conv(kerns Dim3, inS Dim2) *conv {
	return &conv{
		outS:  Dim3{inS[0] - kerns[0] + 1, inS[1] - kerns[1] + 1, kerns[2]},
		kernS: kerns,
		inS:   inS,
	}
}

// ConvMM performs a 2D cross-correlation:
// 	dst[i][j] = Sum_kl src[i+k][j+l] * kern[k][l]
func ConvMM(dst, kern, src M) {

	assureM(dst, Dim2{src.Size[0] - kern.Size[0] + 1, src.Size[1] - kern.Size[1] + 1})

	kMax := kern.Size[1]
	lMax := kern.Size[0]
	iMax := dst.Size[1]

	for i := 0; i < iMax; i++ {
		dsti := dst.Elem(i)
		for j := range dsti {
			sum := 0.0
			for k := 0; k < kMax; k++ {
				srcik := src.Elem(i + k)
				kernk := kern.Elem(k)
				for l := 0; l < lMax; l++ {
					sum += srcik[j+l] * kernk[l]
				}
			}
			dsti[j] = sum
		}
	}
}

// ConvTM performs multiple 2D cross-correlations:
// 	dst[h][i][j] = Sum_kl src[i+k][j+l] * kern[h][k][l]
func ConvTM(dst, kern T3, src M) {
	for i := 0; i < kern.Size[2]; i++ {
		ConvMM(dst.Elem(i), kern.Elem(i), src)
	}
}

// Eval implements Func.
func (f *conv) Eval(y V, θ, x V) {
	dst := Reshape3(y, f.outS)
	kern := f.Kern(θ)
	src := Reshape2(x, f.inS)
	ConvTM(dst, kern, src)
}

// DiffW implements Func.
func (f *conv) DiffW(dy M, y, θ, x V) {
	/*
		y_ij = Σ_kl x_[i+k][j+l] * w_kl

		∂ y_ij / ∂ w_mn = Σ_kl x_[i+k][j+l] * δ_[kl][mn]
		                =      x_[i+m][j+n]
	*/

	assureM(dy, Dim2{f.NumParam(), f.NumOut()})
	set(dy.List, 0)

	for out := 0; out < dy.NumElem(); out++ {

		grad := Reshape3(dy.Elem(out), f.kernS)
		_ = grad

		//for h := 0; h < dy.NumElem(); h++ {
		//	//dy := dy.Elem(h)
		//	//_ = dy
		//}

	}

	//for h := 0; h < dy.NumRows(); h++ { // i = output

	//	for i := 0; i < dyh.NumRows(); i++ {
	//		dyi := dy.Elem(i)

	//		for j := range dyi {

	//			for m := 0; m < f.kernS[0]; m++ {
	//				for n := 0; n < f.kernS[1]; n++ {

	//					dyi[j] = x.Elem(i + m)[j+n]

	//				}
	//			}
	//		}
	//	}
	//}

	//kern := MakeM(Dim2{f.kernS[0], f.kernS[1]})

}

// DiffX implements Func.
func (f *conv) DiffX(dy M, y, θ, x V) {
	assureM(dy, Dim2{f.NumIn(), f.NumOut()})
}

func (f *conv) NumOut() int {
	return f.outS.Len()
}

func (f *conv) NumIn() int {
	return f.inS.Len()
}

func (f *conv) NumParam() int {
	return f.kernS.Len()
}

// Kern slices the convolution kernel from weights w.
func (f *conv) Kern(w V) T3 {
	return Reshape3(w, f.kernS)
}
