package vs

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
	"sync"
)

// LoadImg loads a grayscale PNG image from file.
func LoadImg(fname string) Img {
	f, err := os.Open(fname)
	check(err)
	defer f.Close()
	img, err := png.Decode(f)
	check(err)
	mat := NewImg(img.Bounds().Dy(), img.Bounds().Dx())
	for iy := range mat.Elem {
		for ix := range mat.Elem[iy] {
			_, _ = ix, iy
			r, _, _, _ := img.At(ix, iy).RGBA()
			mat.Elem[iy][ix] = float64(float64(r) / 0xffff)
		}
	}
	return mat
}

func LoadLabeledSet(dir string) []LabeledImg {

	var perDigit [10][]Img
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

	var set []LabeledImg
	for i, imgs := range perDigit {
		for _, img := range imgs {
			set = append(set, LabeledImg{img, i})
		}
	}
	return set
}

// loadAll loads all images in dir.
func loadAll(dir string) []Img {
	ls := readdir(dir)
	img := make([]Img, len(ls))
	for i, f := range ls {
		img[i] = LoadImg(path.Join(dir, f))
	}
	return img
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
