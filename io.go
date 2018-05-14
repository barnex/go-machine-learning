package main

import (
	"image/png"
	"log"
	"os"
)

func loadPNG(fname string) Mat {
	//log.Println("load", fname)
	f, err := os.Open(fname)
	check(err)
	img, err := png.Decode(f)
	check(err)
	mat := NewMat(img.Bounds().Dy(), img.Bounds().Dx())
	for iy := range mat.Elem {
		for ix := range mat.Elem[iy] {
			_, _ = ix, iy
			r, _, _, _ := img.At(ix, iy).RGBA()
			mat.Elem[iy][ix] = float32(float64(r) / 0xffff)
		}
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
