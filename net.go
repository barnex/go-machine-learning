package vs

import "fmt"

type Net interface {

	// NumOut returns the number of outputs (categories).
	NumOut() int

	// NumWeight returns the number of fittable weights.
	NumWeight() int

	// NumIn retruns the number of inputs.
	NumIn() int

	// Eval feeds input x through the network,
	// and stores the raw, unnormalized output in dst.
	// 	len(dst) == NumOut()
	// 	len(w) == NumWeight()
	// 	len(x) == NumIn()
	Eval(dst, w, x []float64)
}

func netCheckSize(f Net, dst, w, x []float64) {
	checkSize(len(dst), f.NumOut())
	checkSize(len(w), f.NumWeight())
	checkSize(len(x), f.NumIn())
}

func netPrintDim(f Net) {
	fmt.Printf("numout: %v, numweight: %v, numin: %v/n", f.NumOut(), f.NumWeight(), f.NumIn())
}
