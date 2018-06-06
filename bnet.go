package vs

type BNet interface {
	OutDim() Dim

	// NumWeight returns the number of fittable weights.
	NumWeight() int

	// NumIn retruns the number of inputs.
	NumIn() int

	Eval(dst *T, w, x []float64)
}
