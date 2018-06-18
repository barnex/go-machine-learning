package vs

// V is a 1-dimensional vector
type V []float64

func MakeV(length int) V {
	return make(V, length)
}

// AssureV makes sure p points to a vector of the specified length.
// If p == nil then a vector is allocated,
// otherwise the size of the existing vector is checked.
func AssureV(p *V, length int) {
	if *p == nil {
		*p = make(V, length)
	}
	checkSize(p.Len(), length)
}

func (v V) Len() int { return len(v) }

func (v V) Dot(b V) float64 {
	return Dot(v, b)
}
