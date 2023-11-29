package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var HOST = "https://dc.dank.pw"
var PORT = "8080"

type Data struct {
	Host     string
	VideoUrl string
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("styles/index.gohtml"))

	data := Data{
		Host: HOST,
	}

	tmpl.Execute(w, data)
}

func Video(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if ps.ByName("video") == "/" || ps.ByName("video") == "" {
		Index(w, r, ps)
	} else {
		RenderVideo(w, r, ps)
	}
}

func RenderVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var videoUrl = strings.TrimPrefix(ps.ByName("video"), "/")

	tmpl := template.Must(template.ParseFiles("styles/video.gohtml"))

	data := Data{
		VideoUrl: videoUrl,
	}

	tmpl.Execute(w, data)
}

func main() {
	router := httprouter.New()
	router.GET("/*video", Video)

	fmt.Printf("Running on port %s\n", PORT)
	fmt.Printf("On host: %s\n", HOST)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
