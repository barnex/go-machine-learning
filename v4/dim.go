package vs

import "fmt"

type Dim [2]int

func (d Dim) String() string {
	return fmt.Sprintf("(%v,%v)", d[0], d[1])
}
