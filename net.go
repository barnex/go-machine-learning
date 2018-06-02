package vs

type Net interface {
	// Eval stores in dst the output for input x and weights w.
	Eval(dst, w, x []float64)
}
