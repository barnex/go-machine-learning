package vs

import "math"

func Loss(f Net, w []float64, xl []LabeledVec) float64 {
	loss := 0.0

	infer := make([]float64, f.NumOut())
	for _, xl := range xl {
		Infer(infer, f, w, xl.X)
		for _, v := range infer {
			Assert(v > 0)
		}
		loss += -math.Log(infer[xl.Label])
	}
	return loss / float64(len(xl))
}
