package vs

/*
import (
	"log"
	"testing"
)

func TestModel1_Aliasing(t *testing.T) {
	m := NewModel1()

	m.params[0] = 123
	testEqf(t, m.w[0].Elem[0][0], 123)

	m.params[len(m.params)-1] = 456

}

func TestModel1(t *testing.T) {
	//if testing.Short() {
	t.Skip("short")
	//}
	model := NewModel1()

	//TrainDumb(model, trainingSet)
	TrainGradNum(model, trainingSet)

	correct := Test(model, testingSet)
	success := float64(correct) / float64(len(testingSet))
	log.Println("have:", success)
	want := 0.88
	testApprox(t, success, want, 0.03)
}

//func TestModel1_Grad(t *testing.T) {
//	m := NewModel1()
//	Randomize(m.Params(), 0.5)
//	Set(m.b, 0)
//
//	x := LabeledImg{Img: NewImg(m.w[0].Rows(), m.w[0].Cols()), Label: 1}
//	Randomize(x.List, 0.5)
//
//	have := make([]float64, len(m.Params()))
//	m.Grad(have, x)
//
//	want := make([]float64, len(m.Params()))
//	GradNumerical(want, m, x)
//
//}
*/
