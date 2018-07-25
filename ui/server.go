package ui

import (
	. "github.com/barnex/vectorstream"
	"github.com/barnex/vectorstream/img"
	"image/png"
	"log"
	"net/http"
)

var imgs = make(map[string]M)

func RegisterImg(name string, m M) {
	imgs[name] = m
}

func Serve(addr string) {
	http.Handle("/render/", http.StripPrefix("/render/", http.HandlerFunc(handleRender)))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleRender(w http.ResponseWriter, r *http.Request) {
	m := imgs[r.URL.Path]

	min, max := MinMax(m.List)
	err := png.Encode(w, img.FromMatrix(m, min, max))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
