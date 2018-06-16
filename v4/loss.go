package vs

import "math"

func Loss1(f Net, w []float64, xl LabeledVec, buf []float64) float64 {
	if buf == nil {
		buf = make([]float64, f.NumOut())
	}
	Infer(buf, f, w, xl.X)
	return -math.Log(buf[xl.Label])
}

func Loss(f Net, w []float64, xl []LabeledVec, buf []float64) float64 {
	if buf == nil {
		buf = make([]float64, f.NumOut())
	}
	loss := 0.0
	for _, xl := range xl {
		loss += Loss1(f, w, xl, buf)
	}
	return loss / float64(len(xl))
}
