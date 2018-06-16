package vs

// M is a matrix.
type M struct {
	List V
	Size Dim2
}

func MakeM(size Dim2) M {
	return Reshape(make([]float64, size.Len()), size)
}

func Reshape(list []float64, size Dim2) M {
	CheckSize(len(list), size.Len())
	return M{list, size}
}

// Len returns the total number of elements.
func (t *M) Len() int {
	return len(t.List)
}

func (t *M) Row(i int) V {
	return t.List[i*t.Size[0] : (i+1)*t.Size[0]]
}

func MulMV(y *V, A M, x V) {
	AssureV(y, A.Size[1])
	for i := range *y {
		(*y)[i] = A.Row(i).Dot(x)
	}
}
