package img

import (
	"fmt"
	"github.com/barnex/vectorstream"
	"image/png"
	"log"
	"os"
	"path"
)

func Print(m vs.M) {
	for j := 0; j < m.NumElem(); j++ {
		row := m.Elem(j)
		for i, v := range row {
			if i != 0 {
				fmt.Print(" ")
			}
			if v == 0 {
				fmt.Print("     ")
			} else {
				fmt.Printf("%.2f", v)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Render(m vs.M, min, max float64) {
	// see https://en.wikipedia.org/wiki/ANSI_escape_code
	const black = 232
	const white = 255
	for j := 0; j < m.NumElem(); j++ {
		row := m.Elem(j)
		for _, v := range row {
			col := int(black + ((v-min)/(max-min))*(white-black))
			if col < black {
				col = 12 // blue, underflow
			}
			if col > white {
				col = 9 // red, overflow
			}
			fmt.Printf("\033[48;5;%dm% 3.2f", col, v)
		}
		fmt.Println("\033[m")
	}
	fmt.Println()
}

// Load loads a grayscale PNG image from file.
func Load(fname string) vs.M {
	f := mustOpen(fname)
	defer f.Close()
	img, err := png.Decode(f)
	check(err)
	b := img.Bounds()
	mat := vs.MakeM(vs.Dim2{b.Dx(), b.Dy()})
	for iy := 0; iy < mat.NumElem(); iy++ {
		row := mat.Elem(iy)
		for ix := range row {
			_, _ = ix, iy
			r, _, _, _ := img.At(ix, iy).RGBA()
			row[ix] = float64(float64(r) / 0xffff)
		}
	}
	return mat
}

func mustOpen(fname string) *os.File {
	f, err := os.Open(fname)
	check(err)
	return f
}

func LoadN(dir string, N int) []vs.M {
	ls := readdir(dir, N)
	img := make([]vs.M, len(ls))
	for i, f := range ls {
		img[i] = Load(path.Join(dir, f))
	}
	return img
}

func readdir(dir string, N int) []string {
	f, err := os.Open(dir)
	check(err)
	ls, err := f.Readdirnames(N)
	check(err)
	return ls
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
