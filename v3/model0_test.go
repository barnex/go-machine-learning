package vs

var (
	// TODO: load all
	trainingSet []LabeledImg = LoadLabeledSet("mnist_png/training", 20)
	testingSet  []LabeledImg = LoadLabeledSet("mnist_png/testing", 20)
)

//func TestModel0(t *testing.T) {
//	model := NewModel0()
//
//	Train(model, trainingSet)
//
//	correct := Test(model, testingSet)
//	success := float64(correct) / float64(len(testingSet))
//	log.Println("have:", success)
//	want := 0.1
//	if !approx(success, want, 0.03) {
//		t.Errorf("success: have: %v, want: %v", success, want)
//	}
//}
