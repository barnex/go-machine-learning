package main

import (
	"fmt"
	"math"
)

func Sum(list []float32) float32 {
	var sum float64
	for _, v := range list {
		sum += float64(v)
	}
	return float32(sum)
}

func Avg(list []float32) float32 {
	return Sum(list) / float32(len(list))
}

func MinMax(list []float32) (min float32, max float32) {
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

func Dot(a, b []float32) float32 {
	checkSize(a, b)
	sum := 0.0
	for i, a := range a {
		sum += float64(a * b[i])
	}
	return float32(sum / float64(len(a)))
}

func Normalize(dst, src []float32) {
	N := float32(len(dst))
	iavg := (1 / (Sum(src) / N))
	for i := range src {
		dst[i] = src[i] * iavg
	}
}

func Add(dst, src []float32) {
	for i := range dst {
		dst[i] += src[i]
	}
}

func AddConst(dst, src []float32, cnst float32) {
	for i := range dst {
		dst[i] = src[i] + cnst
	}
}

func XEntropy(real, pred []float32) float32 {
	checkSize(real, pred)
	var sum float64
	for i := range pred {
		sum += float64(real[i]) * math.Log(float64(pred[i]))
	}
	return float32(-sum)
}

func SoftMax(dst, src []float32) {
	Map64(dst, src, math.Exp)
	Mul(dst, dst, 1/Sum(dst))
}

func Map64(dst, src []float32, f func(float64) float64) {
	checkSize(dst, src)
	for i := range src {
		dst[i] = float32(f(float64(src[i])))
	}
}

func Mul(dst, src []float32, c float32) {
	for i := range src {
		dst[i] = src[i] * c
	}
}

func checkSize(a, b []float32) {
	if len(a) != len(b) {
		panic(fmt.Sprintf("size mismatch: %v != %v", len(a), len(b)))
	}
}

func assert(test bool) {
	if !test {
		panic("assertion failed")
	}
}
