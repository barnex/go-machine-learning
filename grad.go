package vs

//func GradW(dst []float64, f Net, w, x []float64, label int) float64 {
//	const delta = 1. / (1024 * 1024)
//	for i := range dst {
//		backup := w[i]
//
//		w[i] = backup - delta
//		v1 := f.Eval(w, x)
//
//		w[i] = backup + delta
//		v2 := f.Eval(w, x)
//
//		dst[i] = (v2 - v1) / (2 * delta)
//		w[i] = backup
//	}
//	panic("TODO")
//}
