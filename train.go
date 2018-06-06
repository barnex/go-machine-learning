package vs

func GradDescent(f, gll Net, w []float64, xl []LabeledVec, NStep int) {
	const relRate = 1. / 16.

	CheckSize(gll.NumIn(), f.NumIn())
	CheckSize(gll.NumWeight(), f.NumWeight())
	CheckSize(gll.NumOut(), f.NumWeight()) // output = grad_weight f

	buf := make([]float64, gll.NumOut())
	grad := make([]float64, gll.NumOut())

	for i := 0; i < NStep; i++ {
		avgGrad(grad, gll, w, xl, buf)
		lG := Len(grad)
		lW := Len(w)
		rate := relRate * lG / lW
		MAdd(w, w, -rate, grad)

		//for _, w := range m.w {
		//	w.Render(MinMax(w.List))
		//}
		//log.Printf("lG:%v, lP:%v, rate:%v", lG, lP, rate)
		//log.Println("have:", Test(m, trainingSet), "/", len(trainingSet), "loss:", Loss(m, trainingSet))
	}
}

func avgGrad(grad []float64, gll Net, w []float64, xl []LabeledVec, buf []float64) {
	Set(grad, 0)
	for _, xl := range xl {
		gll.Eval(buf, w, xl.X)
		Add(grad, grad, buf)
	}
	Mul(grad, 1/float64(len(xl)), grad)
}

//func TrainDumb(m *Model1, trainingSet []LabeledImg) {
//
//	params := m.Params()
//	Randomize(params, 0.001)
//
//	l := Loss(m, trainingSet)
//
//	count := 0
//	for {
//		fmt.Println(l)
//		i := rand.Intn(len(params))
//		delta := (rand.Float64() - 0.5) * 0.1
//		params[i] += delta
//
//		if l2 := Loss(m, trainingSet); l2 < l {
//			l = l2
//		} else {
//			params[i] -= delta
//		}
//
//		if count%1000 == 0 {
//			log.Println("have:", Test(m, trainingSet), "/", len(trainingSet))
//			for _, w := range m.w {
//				w.Render(MinMax(w.List))
//			}
//		}
//		count++
//	}
//}
