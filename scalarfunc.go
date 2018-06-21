package vs

// ScalarFunc is a differentiable scaler-to-scalar function.
// Typically used as activation function.
type ScalarFunc struct {
	Eval func(x float64) float64
	Diff func(x float64) float64
}

var (
	Re = ScalarFunc{re, step}
)

func re(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func step(x float64) float64 {
	if x > 0 {
		return 1
	}
	return 0
}
