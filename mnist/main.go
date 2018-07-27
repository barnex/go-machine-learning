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
	"github.com/barnex/vectorstream/ui"
)

const (
	numOut = 10
	numPix = 28
	numIn  = numPix * numPix
	//numBias = 10
)

var digits = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var (
	flagRate = flag.Float64("r", 0.05, "learning rate")
)

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

	lu := LU(10*2, numPix*numPix)
	dropout := Dropout(20, 2, 0.10)
	max := MaxPool1D(10, 2)
	net := NewNet(max, dropout, lu)
	w := lu.Weights(net.LParams(0)).List
	imgs := Reshape3(w, Dim3{numPix, numPix, 10 * 2})

	Randomize(net.Params(), .01, 1234)

	//	output := func() {
	//		min, max := MinMax(imgs.List)
	//		for _, m := range imgs.Elems() {
	//			img.RenderText(m, min, max)
	//		}
	//	}

	for i, m := range imgs.Elems() {
		ui.RegisterImg(fmt.Sprintf("lu%02d", i), m)
	}
	go ui.Serve(":2536")

	for i := 0; ; i++ {

		dropout.NextState()

		// decay
		Mul(w, 1-*flagRate, w)

		// step
		loss := GradStep(net, train.Get(), *flagRate)
		fmt.Println(loss)

		// regularize norm == 1
		for _, m := range imgs.Elems() {
			m := m.List
			Mul(m, 1/Norm(m), m)
		}

		//if i%100 == 0 {
		//	output()
		//	dropout.Disable()
		//	fmt.Println(i, loss, Accuracy(net, all[:1000]))
		//	dropout.NextState()
		//}
	}
}

func regularize(x V) {
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
