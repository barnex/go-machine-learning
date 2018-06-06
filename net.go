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
	// and stores the raw, unnormalized output in y.
	// 	len(y) == NumOut()
	// 	len(w) == NumWeight()
	// 	len(x) == NumIn()
	Eval(y, w, x []float64)
}

func NetCheckSize(f Net, y, w, x []float64) {
	CheckSize(len(y), f.NumOut())
	CheckSize(len(w), f.NumWeight())
	CheckSize(len(x), f.NumIn())
}

func netPrintDim(f Net) {
	fmt.Printf("numout: %v, numweight: %v, numin: %v/n", f.NumOut(), f.NumWeight(), f.NumIn())
}
