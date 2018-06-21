package vs

import (
	"math"
	"math/rand"
)

// Add adds a and b.
// 	dst[i] = a[i] + b[i]
func Add(dst *V, a, b V) {
	AssureV(dst, len(a))
	CheckSize(len(a), len(b))
	for i := range *dst {
		(*dst)[i] = a[i] + b[i]
	}
}

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

// Dot returns the dot product
// 	sum_i a[i]*b[i]
func Dot(a, b []float64) float64 {
	CheckSize(len(a), len(b))
	sum := 0.0
	for i, a := range a {
		sum += float64(a * b[i])
	}
	return sum
}

// Len2 returns the length squared of vector x.
func Len2(x []float64) float64 {
	return Dot(x, x)
}

// Len returns the length of vector x.
func Len(x []float64) float64 {
	return math.Sqrt(Dot(x, x))
}

// MAdd performs a multiply+add:
// 	dst[i] = a[i] + s*b[i]
func MAdd(dst, a []float64, s float64, b []float64) {
	CheckSize(len(dst), len(a))
	CheckSize(len(dst), len(b))
	for i := range dst {
		dst[i] = a[i] + s*b[i]
	}
}

// Map applies f to all elements of a list:
// 	dst[i] = f(src[i])
func Map(dst, src []float64, f func(float64) float64) {
	CheckSize(len(dst), len(src))
	for i := range src {
		dst[i] = f(src[i])
	}
}

// MinMax returns the minimum and maximum values.
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

// Mul multiplies by a constant:
// 	dst[i] = s * a[i]
func Mul(dst []float64, s float64, a []float64) {
	CheckSize(len(dst), len(a))
	for i, a := range a {
		dst[i] = s * a
	}
}

// Randomize populates dst with random numbers between -amplitude and +amplitude.
func Randomize(dst []float64, amplitude float64) {
	ampl2 := amplitude * 2
	for i := range dst {
		dst[i] = (rand.Float64() - 0.5) * ampl2
	}
}

// Set sets all elements to value v.
func Set(dst []float64, v float64) {
	for i := range dst {
		dst[i] = v
	}
}

func SoftMax(dst, src []float64) {
	Map(dst, src, math.Exp)
	Mul(dst, 1/Sum(dst), dst)
}

// Sum returns the sum of all elements.
func Sum(list []float64) float64 {
	var sum float64
	for _, v := range list {
		sum += float64(v)
	}
	return float64(sum)
}
