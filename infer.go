package vs

func Infer(dst []float64, f Net, w, x []float64) {
	netCheckSize(f, dst, w, x)
	f.Eval(dst, w, x)
	SoftMax(dst, dst)
}
