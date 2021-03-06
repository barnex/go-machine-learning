package vs

import (
	"math"
	"math/rand"
)

// add adds a and b.
// 	dst[i] = a[i] + b[i]
func add(dst V, a, b V) {
	assureV(dst, len(a))
	checkSize(len(a), len(b))
	for i := range dst {
		dst[i] = a[i] + b[i]
	}
}

// argmax returns the index of the maximum value in list x.
func argmax(x []float64) int {
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

// Average returns the Average of elements.
func Average(x []float64) float64 {
	return Sum(x) / float64(len(x))
}

// dot returns the dot product
// 	sum_i a[i]*b[i]
func dot(a, b []float64) float64 {
	checkSize(len(a), len(b))
	sum := 0.0
	for i, a := range a {
		sum += float64(a * b[i])
	}
	return sum
}

// norm2 returns the length squared of vector x.
func norm2(x []float64) float64 {
	return dot(x, x)
}

// norm returns the length of vector x.
func Norm(x []float64) float64 {
	return math.Sqrt(dot(x, x))
}

// madd performs a multiply+add:
// 	dst[i] = a[i] + s*b[i]
func madd(dst, a []float64, s float64, b []float64) {
	checkSize(len(dst), len(a))
	checkSize(len(dst), len(b))
	for i := range dst {
		dst[i] = a[i] + s*b[i]
	}
}

// Map applies f to all elements of a list:
// 	dst[i] = f(src[i])
func Map(dst, src []float64, f func(float64) float64) {
	checkSize(len(dst), len(src))
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

func Decay(x V, r float64) {
	s := 1 - r
	for i := range x {
		x[i] *= s
	}
}

func maxv(list V) float64 {
	max := list[0]
	for _, v := range list[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Mul multiplies by a constant:
// 	dst[i] = s * a[i]
func Mul(dst []float64, s float64, a []float64) {
	checkSize(len(dst), len(a))
	for i, a := range a {
		dst[i] = s * a
	}
}

// Randomize populates dst with random numbers between -amplitude and +amplitude.
func Randomize(dst []float64, amplitude float64, seed int64) {
	rand := rand.New(rand.NewSource(seed))
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

func softmax(dst, src []float64) {
	Map(dst, src, math.Exp)
	Mul(dst, 1/Sum(dst), dst)
}

// Sum returns the Sum of all elements.
func Sum(list []float64) float64 {
	var sum float64
	for _, v := range list {
		sum += float64(v)
	}
	return float64(sum)
}

// Cross-entropy of softmax
// TODO: move
func softXen(y V, c int) float64 {
	buf := MakeV(len(y)) // TODO: don't alloc
	softmax(buf, y)
	return -math.Log(buf[c])
}

// TODO: move
func gradSoftXen(grad, y V, c int) {
	copyv(grad, y)
	grad[c] -= 1
}

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
