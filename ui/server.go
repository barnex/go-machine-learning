package ui

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	. "github.com/barnex/vectorstream"
	"github.com/barnex/vectorstream/img"
)

var state = struct {
	Imgs   map[string]M
	Matrix map[string]M
	Plots  map[string]*V
	All    []string
}{
	Imgs:   make(map[string]M),
	Matrix: make(map[string]M),
	Plots:  make(map[string]*V),
}

func RegisterImg(name string, m M) {
	state.Imgs[name] = m
	state.All = append(state.All, name)
}

func RegisterMatrix(name string, m M) {
	state.Matrix[name] = m
	state.All = append(state.All, name)
}

func RegisterPlot(name string, data *V) {
	state.Plots[name] = data
	state.All = append(state.All, name)
}

func Serve(addr string) {
	handle("/img/", handleImg)
	handle("/matrix/", handleMatrix)
	handle("/plot/", handlePlot)
	handle("/", handleRoot)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handlePlot(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Path
	data, ok := state.Plots[name]
	if !ok {
		return fmt.Errorf("plot not found: %q", name)
	}

	fname := "plot-" + name + ".txt"
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)

	for _, v := range *data {
		fmt.Fprintln(buf, v)
	}
	buf.Flush()

	plotCmd := fmt.Sprintf(`set term png size 800,300; set key off; set ylabel %q; plot %q w li;`, name, fname)
	return gnuplot(w, plotCmd)
}

func handleMatrix(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("TODO")
}

func gnuplot(w io.Writer, plotCmd string) error {
	cmd := exec.Command("gnuplot", "-e", plotCmd)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout
	if err := cmd.Run(); err != nil {
		return errors.New(stdout.String())
	}
	_, err := w.Write(stdout.Bytes())
	return err
}

func handleImg(w http.ResponseWriter, r *http.Request) error {
	m, ok := state.Imgs[r.URL.Path]
	if !ok {
		return fmt.Errorf("image not found: %q", r.URL.Path)
	}

	min, max := MinMax(m.List)
	return png.Encode(w, img.FromMatrix(m, min, max))
}

func handleRoot(w http.ResponseWriter, r *http.Request) error {
	return rootTemplate.Execute(w, state)
}

func handle(pattern string, h func(http.ResponseWriter, *http.Request) error) {
	http.Handle(pattern, http.StripPrefix(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})))
}

var rootTemplate = template.Must(template.New("root").Parse(rootHTML))

const rootHTML = `
<html>

 <head>

  <script>
  
   function refresh(){
    {{range $k,$v := .Plots}}  
   	 document.getElementById("{{$k}}").src = "/plot/{{$k}}?cachebreak=" + Math.random();
    {{end}}  
    {{range $k,$v := .Matrix}}  
   	 document.getElementById("{{$k}}").src = "/matrix/{{$k}}?cachebreak=" + Math.random();
    {{end}}  
    {{range $k,$v := .Imgs}}  
     document.getElementById("{{$k}}").src = "/img/{{$k}}?cachebreak=" + Math.random();
    {{end}}  
   }
   window.setInterval(refresh, 1000)
   
  </script>

 </head>

 <body>

   <div>
    {{range $k,$v := .Plots}}  
      <img id="{{$k}}" src="/plot/{{$k}}"></img>
    {{end}}  
   </div>

   <div style="float:left">
    {{range $k,$v := .Imgs}}  
     <figure style="float:left; padding-left:2px;">
      <img id="{{$k}}" width={{index .Size 0}} height={{index .Size 1}} src="/img/{{$k}}"></img>
      <figcaption style="font-size:6pt">{{$k}}</figcaption>
     </figure>
   {{end}}  
   </div>


 </body>

</html>
`
