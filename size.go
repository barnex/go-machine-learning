package vs

type Size [4]int

func (s *Size) Len() int {
	return s[0] * s[1] * s[2] * s[3]
}
