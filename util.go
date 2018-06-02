package vs

import (
	"fmt"
	"math"
	"path"
	"runtime"
)

var digits = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func checkSize(a, b int) {
	if a != b {
		panic(fmt.Sprintf("%v: size mismatch: %v != %v", caller(1), a, b))
	}
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
