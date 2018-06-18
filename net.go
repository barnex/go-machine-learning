package vs

import "fmt"

type Net interface {

	// NumOut returns the number of outputs (categories).
	NumOut() int

	// NumParam returns the number of fittable weights.
	NumParam() int

	// NumIn retruns the number of inputs.
	NumIn() int

	// Eval feeds input x through the network,
	// and stores the raw, unnormalized output in y.
	// 	len(y) == NumOut()
	// 	len(w) == NumWeight()
	// 	len(x) == NumIn()
	Eval(y *V, w, x V)

	//Grad(gy T, w, x []float64)
}

func NetCheckSize(f Net, y, w, x []float64) {
	if !(len(y) == f.NumOut() &&
		len(w) == f.NumParam() &&
		len(x) == f.NumIn()) {
		panic(fmt.Sprintf("size mismatch: have y:%v, w:%v, x:%v, want: y:%v, w:%v, x:%v", len(y), len(w), len(x), f.NumOut(), f.NumParam(), f.NumIn()))
	}
}

func netPrintDim(f Net) {
	fmt.Printf("numout: %v, numparam: %v, numin: %v/n", f.NumOut(), f.NumParam(), f.NumIn())
}

type BackProp func(dst, w, x []float64)
