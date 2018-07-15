package vs

import (
	"fmt"
	"math"
	"path"
	"runtime"
)

func copyv(a, b V) {
	checkSize(len(a), len(b))
	copy(a, b)
}

func checkSize(a, b int) {
	if a != b {
		panic(fmt.Sprintf("%v: size mismatch: %v != %v", caller(2), a, b))
	}
}

func checkDim2(a, b Dim2) {
	if a != b {
		panic(fmt.Sprintf("%v: size mismatch: %v != %v", caller(2), a, b))
	}
}

// assureM checks p is a matrix of the specified size.
// TODO: remove
func assureM(p M, size Dim2) {
	checkDim2(p.Size, size)
}

// assureV checks p is a vector of the specified size.
// TODO: remove
func assureV(p V, length int) {
	checkSize(p.Len(), length)
}

func assert(test bool) {
	if !test {
		panic(fmt.Sprintf("%v: assertion failed", caller(1)))
	}
}

func approxEq(a, b, tol float64) bool {
	return math.Abs(a-b) < tol
}

func caller(skip int) string {
	_, file, line, _ := runtime.Caller(skip + 1)
	return fmt.Sprintf("%v:%v", path.Base(file), line)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
