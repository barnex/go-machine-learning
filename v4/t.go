package vs

type T struct {
	list []float64
	size Dim
}

func MakeT(rows, cols int) T {
	return Reshape(make([]float64, cols*rows), rows, cols)
}

func Reshape(list []float64, nx, ny int) T {
	CheckSize(len(list), nx*ny)
	return T{list, Dim{nx, ny}}
}

// Len returns the total number of elements
func (t *T) Len() int {
	return len(t.list)
}

func (t *T) Size(dim int) int {
	return t.size[dim]
}

// List returns all elements in a contiguous list.
func (t *T) List() []float64 {
	return t.list
}

func (t *T) Row(i int) []float64 {
	return t.list[i*t.size[0] : (i+1)*t.size[0]]
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
