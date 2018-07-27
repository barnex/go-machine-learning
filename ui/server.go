package ui

import (
	"html/template"
	"image/png"
	"log"
	"net/http"

	. "github.com/barnex/vectorstream"
	"github.com/barnex/vectorstream/img"
)

var state = struct {
	Imgs map[string]M
}{
	Imgs: make(map[string]M),
}

func RegisterImg(name string, m M) {
	state.Imgs[name] = m
}

func Serve(addr string) {
	http.Handle("/render/", http.StripPrefix("/render/", http.HandlerFunc(handleRender)))
	http.Handle("/", http.HandlerFunc(handleRoot))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleRender(w http.ResponseWriter, r *http.Request) {
	m := state.Imgs[r.URL.Path]

	min, max := MinMax(m.List)
	err := png.Encode(w, img.FromMatrix(m, min, max))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, state)
}

var rootTemplate = template.Must(template.New("root").Parse(rootHTML))

const rootHTML = `
<html>

 <head>

  <script>
  
   function refresh(){
  {{range $k,$v := .Imgs}}  
   	document.getElementById("{{$k}}").src = "/render/{{$k}}?cachebreak=" + Math.random();
  {{end}}  
   }
   window.setInterval(refresh, 500)
   
  </script>

 </head>

 <body>
   <div style="float:left">
  {{range $k,$v := .Imgs}}  
    <figure style="float:left; padding-left:2px;">
     <img id="{{$k}}" width={{index .Size 0}} height={{index .Size 1}} src="/render/{{$k}}"></img>
     <figcaption style="font-size:6pt">{{$k}}</figcaption>
    </figure>
  {{end}}  
   </div>
 </body>

</html>
`
