package vs

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
	madd(n.Params(), n.Params(), rate/N, step)
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
	byLabel [][]V
	pos     int
}

//func(s*TrainingSet) Get()
