package vs

import (
	"fmt"
	"path"

	. "github.com/barnex/vectorstream"
	"github.com/barnex/vectorstream/img"
)

var _trainSet []LabeledVec

func trainSet() []LabeledVec {
	if _trainSet != nil {
		return _trainSet
	}
	for i := range digits {
		imgs := img.LoadN(path.Join("mnist_png/training", fmt.Sprint(i)), 10)
		for _, x := range imgs {
			_trainSet = append(_trainSet, LabeledVec{i, x.List})
		}
	}
	return _trainSet
}
