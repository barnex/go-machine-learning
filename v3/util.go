package vs

import "fmt"

var digits = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func checkSize(have, want int) {
	if have != want {
		panic(fmt.Sprintf("size: have %v, want: %v", have, want))
	}
}
