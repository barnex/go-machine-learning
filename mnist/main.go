package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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

	lu1 := LU(16, numPix*numPix)
	lu2 := LU(10, 16)
	lu3 := LU(10, 10)

	net := NewNet(lu3, Re(lu2), Re(lu1))

	imgs := Reshape3(lu1.Weights(net.LParams(0)).List, Dim3{numPix, numPix, lu1.NumOut()})
	for i, m := range imgs.Elems() {
		ui.RegisterImg(fmt.Sprintf("relu1_%02d", i), m)
	}
	ui.RegisterImg("relu2", lu2.Weights(net.LParams(1)))
	ui.RegisterImg("lu3", lu3.Weights(net.LParams(2)))

	y := MakeV(net.NumOut())
	dy := MakeV(net.NumParam())

	Randomize(net.Params(), .01, 134)
	Map(net.Params(), net.Params(), math.Abs)

	var loss V
	ui.RegisterPlot("loss", &loss)
	var accuracy V
	ui.RegisterPlot("accuracy", &accuracy)

	p := net.Params()
	ui.RegisterPlot("params", &p)
	ui.RegisterPlot("grad", &dy)

	//for i, m := range imgs.Elems() {
	//	ui.RegisterImg(fmt.Sprintf("lu%02d", i), m)
	//}
	go ui.Serve(":2536")

	for i := 0; ; i++ {

		if i%10 == 0 {
			//	dropout.Disable()
			acc := Accuracy(net, all[:1000])
			accuracy = append(accuracy, acc)
		}

		//dropout.NextState()

		// decay
		//Mul(w, 1-*flagRate, w)

		// step
		l := GradStepBuf(net, train.Get(), *flagRate, y, dy)
		loss = append(loss, l)

		// regularize norm == 1
		//for _, m := range imgs.Elems() {
		//	m := m.List
		//	Mul(m, 1/Norm(m), m)
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
