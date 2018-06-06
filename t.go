package vs

type T struct {
	list []float64
	size Dim
}

func NewT(rows, cols int) *T {
	return Reshape(make([]float64, cols*rows), rows, cols)
}

func Reshape(list []float64, rows, cols int) *T {
	CheckSize(len(list), rows*cols)
	return &T{list, Dim{rows, cols}}
}

// Len returns the total number of elements
func (t *T) Len() int {
	return len(t.list)
}

//func (m Img) Print() {
//	for _, row := range m.Elem {
//		for i, v := range row {
//			if i != 0 {
//				fmt.Print(" ")
//			}
//			if v == 0 {
//				fmt.Print("     ")
//			} else {
//				fmt.Printf("%.2f", v)
//			}
//		}
//		fmt.Println()
//	}
//	fmt.Println()
//}
//
//func (m Img) Render(min, max float64) {
//	// see https://en.wikipedia.org/wiki/ANSI_escape_code
//	const black = 232
//	const white = 255
//	for _, row := range m.Elem {
//		for _, v := range row {
//			col := int(black + ((v-min)/(max-min))*(white-black))
//			if col < black {
//				col = 12 // blue, underflow
//			}
//			if col > white {
//				col = 9 // red, overflow
//			}
//			fmt.Printf("\033[48;5;%dm% 3.2f", col, v)
//		}
//		fmt.Println("\033[m")
//	}
//	fmt.Println()
//}
