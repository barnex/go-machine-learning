package vs

import (
	"bytes"
	"fmt"
	"io"
)

// V is a 1-dimensional vector
type V []float64

func MakeV(length int) V {
	return make(V, length)
}

func (v V) Len() int { return len(v) }

func (v V) PrintTo(w io.Writer) {
	fmt.Fprint(w, v[0])
	for _, v := range v[1:] {
		fmt.Fprint(w, " ", v)
	}
}

func (v V) String() string {
	return printToString(v)
}

func (v V) Dot(b V) float64 {
	return dot(v, b)
}

func printToString(p interface{ PrintTo(io.Writer) }) string {
	var b bytes.Buffer
	p.PrintTo(&b)
	return b.String()
}
