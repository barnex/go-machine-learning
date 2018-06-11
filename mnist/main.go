package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"path"

	. "github.com/barnex/vectorstream"
	"github.com/barnex/vectorstream/img"
)

const (
	numOut  = 10
	numIn   = 28 * 28
	numBias = 10
)

var (
	flagNTrain = flag.Int("ntrain", -1, "limit size of training set")
	flagNStep  = flag.Int("nstep", 1, "number of gradient descent steps")
)

var digits = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("need 1 argument")
	}

	dir := flag.Arg(0)
	train := loadSet(path.Join(dir, "training"), *flagNTrain)

	mod := NewModel(NewLU(numOut, numIn))
	Randomize(mod.W, 1)

	test := loadSet(path.Join(dir, "testing"), -1)

	for {
		mod.GradDescent(train, *flagNStep)
		log.Printf("loss: %5f, accuracy: %v/%v", mod.Loss(train), mod.Accuracy(test), len(test))
	}

}

func loadSet(dir string, N int) []LabeledVec {
	var set []LabeledVec
	for i := range digits {
		imgs := img.LoadN(path.Join(dir, fmt.Sprint(i)), N)
		for _, x := range imgs {
			set = append(set, LabeledVec{i, x.List})
		}
	}
	rand.Shuffle(len(set), func(i, j int) { set[i], set[j] = set[j], set[i] })
	return set
}
