package main

func Sum(list []float32) float64 {
	var sum float64
	for _, v := range list {
		sum += float64(v)
	}
	return sum
}

func dot(a, b []float32) float64 {
	sum := 0.0
	for i, a := range a {
		sum += float64(a * b[i])
	}
	return sum
}

func Normalize(dst, src []float32) {
	N := float64(len(dst))
	iavg := float32(1 / (Sum(src) / N))
	for i := range src {
		dst[i] = src[i] * iavg
	}
}

func add(dst, src []float32) {
	for i := range dst {
		dst[i] += src[i]
	}
}

func AddConst(dst, src []float32, cnst float32) {
	for i := range dst {
		dst[i] = src[i] + cnst
	}
}
