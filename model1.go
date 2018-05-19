package main

import (
	"fmt"
	"path"
)

type Model1 struct {
	w []Mat
}

func (m *Model1) Train(dir string) {
	dir = path.Join(dir, "training")
	m.w = make([]Mat, 10)

	for i := range m.w {
		m.w[i] = NewMat(28, 28)
		w := m.w[i]
		dir := path.Join(dir, fmt.Sprint(i))
		for _, f := range readdir(dir) {
			img := loadPNG(path.Join(dir, f))
			Add(w.List, img.List)
		}
		stdnorm(w.List, w.List)
	}
}

func (m *Model1) Infer(img Mat) []float32 {
	stdnorm(img.List, img.List)
	infer := make([]float32, 10)
	for i, w := range m.w {
		infer[i] = Dot(img.List, w.List)
	}
	return infer
}

func stdnorm(dst, src []float32) {
	Normalize(dst, src)
	AddConst(dst, dst, -0.5*Avg(dst))
}
