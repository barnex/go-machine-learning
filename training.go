package vs

import (
	"log"
	"math/rand"
)

func GradDescent(n *Net, s TrainingSet, rate float64, nStep int) {

}

func GradStep(n *Net, set []LV, rate float64) (avgLoss float64) {
	y := MakeV(n.NumOut())
	dy := MakeV(n.NumParam())
	step := MakeV(n.NumParam())

	totalLoss := 0.0
	for _, lx := range set {
		totalLoss += n.Backprop(dy, y, lx.V, lx.Label)
		add(step, step, dy)
	}

	N := float64(len(set))
	madd(n.Params(), n.Params(), -rate/N, step)
	return totalLoss / N
}

func Accuracy(n *Net, set []LV) float64 {
	y := MakeV(n.NumOut())
	correct := 0
	for _, lx := range set {
		n.Eval(y, lx.V)
		if argmax(y) == lx.Label {
			correct++
		}
	}
	return float64(correct) / float64(len(set))
}

type TrainingSet struct {
	ByLabel [][]V
	pos     int
}

func (s *TrainingSet) Get() []LV {
	lv := make([]LV, len(s.ByLabel))
	for i := range lv {
		lv[i] = LV{Label: i, V: s.ByLabel[i][s.pos]}
	}
	s.pos++
	if s.pos == len(s.ByLabel[0]) {
		s.pos = 0
		log.Println("wrapping set")
	}
	return lv
}

func (s *TrainingSet) Shuffle() {
	for _, set := range s.ByLabel {
		rand.Shuffle(len(set), func(i, j int) { set[i], set[j] = set[j], set[i] })
	}
}
