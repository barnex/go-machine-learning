package main

import (
	"flag"
	"fmt"
	"log"
	"path"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalf("need 1 directory")
	}
	dir := path.Join(flag.Arg(0), "testing")

	var (
		hit, miss int
	)

	var m Method1
	m.Train(flag.Arg(0))

	var count [10][10]int
	for digit := 0; digit < 10; digit++ {

		dir := path.Join(dir, fmt.Sprint(digit))

		for _, name := range readdir(dir) {
			img := loadPNG(path.Join(dir, name))
			infer := m.Infer(img)
			count[digit][infer]++
			if infer == digit {
				hit++
			} else {
				miss++
			}
		}
	}

	for i := range count {
		for j := range count[i] {
			fmt.Printf("%v -> %v: %v \n", i, j, count[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("hit: %v, miss: %v, error: %.3f%%\n", hit, miss, 100*(float64(miss)/float64(hit+miss)))
}
