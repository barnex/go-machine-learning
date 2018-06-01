package vs

import (
	"fmt"
	"log"
	"math/rand"
)

func TrainDumb(m *Model1, training []LabeledImg) {

	params := m.Params()

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
