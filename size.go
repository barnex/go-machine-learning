package vs

type Size [4]int

func (s *Size) Len() int {
	return s[0] * s[1] * s[2] * s[3]
}

func (s *Size) bytes() int64 {
	return int64(s.Len()) * sizeOfFloat32
}

const sizeOfFloat32 = 4
