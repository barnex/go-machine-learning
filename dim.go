package vs

import "fmt"

type Dim interface {
	Len() int
}

// Dim stores a tensor size.
type Dim2 [2]int

// Len returns the total number of elements.
func (d Dim2) Len() int       { return d[0] * d[1] }
func (d Dim2) String() string { return tupleStr(d[:]) }

type Dim3 [3]int

// Len returns the total number of elements.
func (d Dim3) Len() int       { return d[0] * d[1] * d[2] }
func (d Dim3) String() string { return tupleStr(d[:]) }

type Dim4 [4]int

// Len returns the total number of elements.
func (d Dim4) Len() int       { return d[0] * d[1] * d[2] * d[3] }
func (d Dim4) String() string { return tupleStr(d[:]) }

type Dim5 [5]int

// Len returns the total number of elements.
func (d Dim5) Len() int       { return d[0] * d[1] * d[2] * d[3] * d[4] }
func (d Dim5) String() string { return tupleStr(d[:]) }

func tupleStr(d []int) string {
	str := fmt.Sprint("(", d[0])
	for _, v := range d[1:] {
		str += fmt.Sprint(",", v)
	}
	str += ")"
	return str
}
