package main

import (
	"fmt"
	"path"
)

type Method1 struct {
	w []Mat
}

func (m *Method1) Train(dir string) {
	dir = path.Join(dir, "training")
	m.w = make([]Mat, 10)

	for i := range m.w {
		m.w[i] = NewMat(28, 28)
		w := m.w[i]
		dir := path.Join(dir, fmt.Sprint(i))
		for _, f := range readdir(dir) {
			img := loadPNG(path.Join(dir, f))
			add(w.List, img.List)
		}
		stdnorm(w.List, w.List)
	}
}

func (m *Method1) Infer(img Mat) int {
	bestOverlap := 0.0
	bestI := 0
	stdnorm(img.List, img.List)
	for i, w := range m.w {
		if overlap := dot(img.List, w.List); overlap > bestOverlap {
			bestI = i
			bestOverlap = overlap
		}
	}
	return bestI
}

func stdnorm(dst, src []float32) {
	Normalize(dst, src)
	c := float32(-0.5 * float64(len(dst)))
	AddConst(dst, dst, c)
}
