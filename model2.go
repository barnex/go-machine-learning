package main

import (
	"fmt"
	"math"
	"path"
)

type Model2 struct {
	w        [10]Mat
	b        [10]float64
	training [10][]Mat
}

func (m *Model2) Train(dir string) {
	m.training = loadAllDigits(path.Join(dir, "training"))

	for d := range digits {
		m.w[d] = NewMat(28, 28)
	}

	for i := 1; i < 10; i++ {
		for d, w := range m.w {
			for i := range w.List {
				w.List[i] += 0.1 * m.Diff(d, i)
			}
		}

		for _, w := range m.w {
			w.Render(MinMax(w.List))
			fmt.Println()
		}
		fmt.Println(m.Loss())
	}

}

func (m *Model2) Diff(d, i int) float32 {

	var der float64

	for j := range m.training[d][i].List {

		var avg float32
		for d := range m.training {
			avg += m.training[d][j].List[i]
		}
		avg /= 10

		der += float64(m.training[d][j].List[i] - avg)
	}

	return float32(der / float64(len(m.training[0][0].List)))
}

func (m *Model2) Infer(img Mat) []float32 {
	prob := make([]float32, 10)
	for i, w := range m.w {
		prob[i] = Dot(img.List, w.List)
	}
	SoftMax(prob, prob)
	return prob
}

func (m *Model2) Loss() float64 {
	loss := 0.0
	for i := range m.training {
		var real [10]float32
		real[i] = 1
		for _, img := range m.training[i] {
			pred := m.Infer(img)
			loss += float64(XEntropy(real[:], pred))
		}
	}
	return loss
}

func findMax(x []float32) int {
	maxX := float32(math.Inf(-1))
	maxI := 0
	for i, x := range x {
		if x > maxX {
			maxI = i
			maxX = x
		}
	}
	return maxI
}

//func (m *Model2) load(dir string) {
//	log.Println("loading", dir)
//	m.training = loadAllDigits(path.Join(dir, "training"))
//	log.Println("done")
//	for i := range m.w {
//		m.w[i] = NewMat(28, 28)
//		w := m.w[i]
//		for _, img := range m.training[i] {
//			Add(w.List, img.List)
//		}
//		stdnorm(w.List, w.List)
//	}
//}
