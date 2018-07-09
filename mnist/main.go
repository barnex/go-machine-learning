package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"

	. "github.com/barnex/vectorstream"
	"github.com/barnex/vectorstream/img"
)

const (
	numOut  = 10
	numIn   = 28 * 28
	numBias = 10
)

var digits = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	log.SetFlags(0)
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("need 1 argument")
	}

	dir := flag.Arg(0)
	log.Println("loading data...")
	start := time.Now()
	train, all := loadSet(path.Join(dir, "training"), 5000)
	log.Printf("loading data: %v examples, %v", len(all), time.Since(start))

	net := NewNet(LU(numOut, numIn))
	Randomize(net.Params(), .01)

	out := make([][]float64, net.NumParam())

	for i := range train.ByLabel[0] {
		loss := GradStep(net, train.Get(), 0.03)
		//fmt.Println(i, loss)
		if i%100 == 0 {
			fmt.Println(i, loss, Accuracy(net, all[:1000]))

			for j, w := range net.Params() {
				out[j] = append(out[j], w)
			}
		}
	}

	saveW(out)

	fmt.Println(Accuracy(net, all))

}

func saveW(w [][]float64) {
	f, err := os.Create("weights.txt")
	check(err)
	defer f.Close()

	for _, w := range w {
		for j, w := range w {
			fmt.Fprintln(f, j, w)
		}
		fmt.Fprintln(f)
	}
}

func loadSet(dir string, N int) (*TrainingSet, []LV) {
	set := &TrainingSet{ByLabel: make([][]V, 10)}
	var all []LV
	var wg sync.WaitGroup
	wg.Add(len(digits))

	for i := range digits {
		go func(i int) {
			defer wg.Done()
			imgs := img.LoadN(path.Join(dir, fmt.Sprint(i)), N)
			for j := range imgs {
				set.ByLabel[i] = append(set.ByLabel[i], imgs[j].List)
			}
		}(i)
	}
	wg.Wait()
	for i := range digits {
		for _, x := range set.ByLabel[i] {
			all = append(all, LV{Label: i, V: x})
		}
	}
	set.Shuffle()
	rand.Shuffle(len(all), func(i, j int) { all[i], all[j] = all[j], all[i] })
	return set, all
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
