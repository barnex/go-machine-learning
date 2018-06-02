package vs

import (
	"fmt"
	"log"
	"math/rand"
)

func TrainGradNum(m *Model1, trainingSet []LabeledImg) {

	params := m.Params()
	Randomize(params, 0.001)
	grad := make([]float64, len(params))
	relRate := 1. / 16.

	for {
		GradNumerical(grad, m, trainingSet)
		fmt.Println(grad)

		lG := Len(grad)
		lP := Len(params)
		rate := relRate * lG / lP
		MAdd(params, params, -rate, grad)

		for _, w := range m.w {
			w.Render(MinMax(w.List))
		}
		log.Printf("lG:%v, lP:%v, rate:%v", lG, lP, rate)
		log.Println("have:", Test(m, trainingSet), "/", len(trainingSet), "loss:", Loss(m, trainingSet))
	}
}

func TrainDumb(m *Model1, trainingSet []LabeledImg) {

	params := m.Params()
	Randomize(params, 0.001)

	l := Loss(m, trainingSet)

	count := 0
	for {
		fmt.Println(l)
		i := rand.Intn(len(params))
		delta := (rand.Float64() - 0.5) * 0.1
		params[i] += delta

		if l2 := Loss(m, trainingSet); l2 < l {
			l = l2
		} else {
			params[i] -= delta
		}

		if count%1000 == 0 {
			log.Println("have:", Test(m, trainingSet), "/", len(trainingSet))
			for _, w := range m.w {
				w.Render(MinMax(w.List))
			}
		}
		count++
	}
}
