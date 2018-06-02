package vs

type Model interface {

	// NumLabels returns the model's number of labels,
	// e.g.: 10 ditits for MNIST.
	NumLabels() int

	// Infer stores the guessed probabilities per label in dst.
	Infer(dst []float64, img Img)

	//RawGrad(dst []float64, x LabeledImg)

	Params() []float64
}
