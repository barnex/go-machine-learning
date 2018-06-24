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

// AssureM makes sure p points to a matrix of the specified size.
// If p == nil then a matrix is allocated,
// otherwise the size of the existing matrix is checked.
func AssureM(p M, size Dim2) {
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

func MulMV(y V, A M, x V) {
	AssureV(y, A.Size[1])
	for i := range y {
		y[i] = A.Row(i).Dot(x)
	}
}
