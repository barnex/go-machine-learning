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
	checkSize(len(list), size.Len())
	return M{list, size}
}

// assureM makes sure p points to a matrix of the specified size.
// If p == nil then a matrix is allocated,
// otherwise the size of the existing matrix is checked.
func assureM(p M, size Dim2) {
	checkDim2(p.Size, size)
}

// Len returns the total number of elements.
func (t M) Len() int {
	return len(t.List)
}

func (m M) IsNil() bool {
	return m.Size == Dim2{} && m.List == nil
}

func (t M) Row(i int) V {
	return t.List[i*t.Size[0] : (i+1)*t.Size[0]]
}

func (t M) Rows() int {
	return t.Size[1]
}

func mulMV(y V, A M, x V) {
	assureV(y, A.Size[1])
	for i := range y {
		y[i] = A.Row(i).Dot(x)
	}
}

func mulVM(y V, x V, A M) {
	assureV(y, A.Size[0])
	assureV(x, A.Size[1])
	for i := 0; i < A.Size[0]; i++ {
		y[i] = 0
		for j := 0; j < A.Size[1]; j++ {
			y[i] += x[j] * A.Row(j)[i]
		}
	}
	/*
	   x: 2
	   A: 2x3 = {3, 2}
	   [x0 x1] [ a00 a01 a02 ]
	           [ a10 a11 a12 ]
	   y: 3
	*/
}
