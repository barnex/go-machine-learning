package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
	"sync"
)

// loadAllDigits loads all images under subdirectories 0/ 1/ ... 9/ under dir.
func loadAllDigits(dir string) [10][]Mat {
	var perDigit [10][]Mat
	var wg sync.WaitGroup
	wg.Add(len(perDigit))
	for i := range perDigit {
		i := i
		go func() {
			defer wg.Done()
			perDigit[i] = loadAll(path.Join(dir, fmt.Sprint(i)))
		}()
	}
	wg.Wait()
	return perDigit
}

// loadAll loads all images in dir.
func loadAll(dir string) []Mat {
	ls := readdir(dir)
	img := make([]Mat, len(ls))
	for i, f := range ls {
		img[i] = loadPNG(path.Join(dir, f))
	}
	return img
}

// loadPNG loads one image.
func loadPNG(fname string) Mat {
	//log.Println("load", fname)
	f, err := os.Open(fname)
	check(err)
	defer f.Close()
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
