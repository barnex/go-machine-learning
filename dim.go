package vs

import "fmt"

// Dim stores a tensor size.
type Dim2 [2]int

// Len returns the total number of elements.
func (d Dim2) Len() int {
	return d[0] * d[1]
}

//func (d Dim2) String() string {
//	return fmt.Sprintf("(%v,%v)", d[0], d[1])
//}

func CheckDim2(a, b Dim2) {
	if a != b {
		panic(fmt.Sprintf("%v: size mismatch: %v != %v", caller(2), a, b))
	}
}
