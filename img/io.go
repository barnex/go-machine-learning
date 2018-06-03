package img

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
)

func (m Img) Print() {
	for _, row := range m.Elem {
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

func (m Img) Render(min, max float64) {
	// see https://en.wikipedia.org/wiki/ANSI_escape_code
	const black = 232
	const white = 255
	for _, row := range m.Elem {
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

// LoadImg loads a grayscale PNG image from file.
func Load(fname string) Img {
	f := mustOpen(fname)
	defer f.Close()
	img, err := png.Decode(f)
	check(err)
	mat := New(img.Bounds().Dy(), img.Bounds().Dx())
	for iy := range mat.Elem {
		for ix := range mat.Elem[iy] {
			_, _ = ix, iy
			r, _, _, _ := img.At(ix, iy).RGBA()
			mat.Elem[iy][ix] = float64(float64(r) / 0xffff)
		}
	}
	return mat
}

func mustOpen(fname string) *os.File {
	f, err := os.Open(fname)
	check(err)
	return f
}

func LoadN(dir string, N int) []Img {
	ls := readdir(dir, N)
	img := make([]Img, len(ls))
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
		log.Fatal(err)
	}
}
