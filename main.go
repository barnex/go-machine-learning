package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalf("need 1 directory")
	}
	dir := flag.Arg(0)

	for i := 0; i < 10; i++ {
		dir := path.Join(dir, fmt.Sprint(i))
		for _, name := range readdir(dir) {
			img := loadPNG(path.Join(dir, name))
			fmt.Println(img)
		}
	}

}

func predict(img [][]float32) int {
	return 0
}

func loadPNG(fname string) [][]float32 {
	log.Println("load", fname)
	f, err := os.Open(fname)
	check(err)
	img, err := png.Decode(f)
	check(err)
	mat := newMatrix(img.Bounds().Dy(), img.Bounds().Dx())
	for iy := range mat {
		for ix := range mat[iy] {
			_, _ = ix, iy
			r, _, _, _ := img.At(iy, ix).RGBA()
			mat[iy][ix] = float32(r) / 0xffff
		}
	}
	return mat
}

func newMatrix(rows, cols int) [][]float32 {
	list := make([]float32, cols*rows)
	mat := make([][]float32, rows)
	for iy := range mat {
		mat[iy] = list[iy*cols : (iy+1)*cols]
	}
	return mat
}

func readdir(dir string) []string {
	f, err := os.Open(dir)
	check(err)
	ls, err := f.Readdirnames(-1)
	check(err)
	return ls
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
