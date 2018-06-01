package vs

import (
	"math"
	"math/rand"
)

// ArgMax returns the index of the maximum value in list x.
func ArgMax(x []float64) int {
	maxX := x[0]
	maxI := 0
	for i, x := range x {
		if x > maxX {
			maxI = i
			maxX = x
		}
	}
	return maxI
}

func Normalize(dst, src []float64) {
	N := float64(len(dst))
	iavg := (1 / (Sum(src) / N))
	for i := range src {
		dst[i] = src[i] * iavg
	}
}

func Randomize(dst []float64, ampl float64) {
	ampl2 := ampl * 2
	for i := range dst {
		dst[i] = (rand.Float64() - 0.5) * ampl2
	}
}

func Sum(list []float64) float64 {
	var sum float64
	for _, v := range list {
		sum += float64(v)
	}
	return float64(sum)
}

//func Avg(list []float64) float64 {
//	return Sum(list) / float64(len(list))
//}

func MinMax(list []float64) (min float64, max float64) {
	min = list[0]
	max = list[0]
	for _, v := range list {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func DotNorm(a, b []float64) float64 {
	checkSize(len(a), len(b))
	sum := 0.0
	for i, a := range a {
		sum += float64(a * b[i])
	}
	return float64(sum / float64(len(a)))
}

//func Add(dst, src []float64) {
//	for i := range dst {
//		dst[i] += src[i]
//	}
//}
//
//func AddConst(dst, src []float64, cnst float64) {
//	for i := range dst {
//		dst[i] = src[i] + cnst
//	}
//}
//
//func XEntropy(real, pred []float64) float64 {
//	checkSize(real, pred)
//	var sum float64
//	for i := range pred {
//		sum += float64(real[i]) * math.Log(float64(pred[i]))
//	}
//	return float64(-sum)
//}

func SoftMax(dst, src []float64) {
	Map64(dst, src, math.Exp)
	Mul(dst, dst, 1/Sum(dst))
}

func Map64(dst, src []float64, f func(float64) float64) {
	checkSize(len(dst), len(src))
	for i := range src {
		dst[i] = float64(f(float64(src[i])))
	}
}

func Mul(dst, src []float64, c float64) {
	for i := range src {
		dst[i] = src[i] * c
	}
}
