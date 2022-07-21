package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var HOST = "https://dc.noury.ee"
var PORT = "8080"

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var videoUrl = "https://cdn.discordapp.com/attachments/381520882608373761/989666371178754068/denkcats_1639474686233272.mp4"

	var indexHtml = fmt.Sprintf(`
<html>
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <title>Discord Video Embedder</title>
        <meta charset="UTF-8" />
        <meta name="description" content="" />
    </head>
    <body style="background-color:#181a1b;">
    <h1 style="color:#d8d4cf">Discord Video Embedder</h1>
    <p style="color:#d8d4cf">Lets you watch a discord video in the browser instead of downloading it by default. <br>
    Simply paste the link to a discord video at the end of the url (after the /, if there is none add one).
    </p>
    <p style="color:#d8d4cf">Original: <a style="color:#3391ff" href="%s/%s">%s/%s</a></p>
    <p style="color:#d8d4cf"><sup>this is beautiful web design shut up</sup></p>
    </body>
</html>
`, HOST, videoUrl, HOST, videoUrl)

	fmt.Fprintf(w, "%s", indexHtml)
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

	var videoHtml = fmt.Sprintf(`
<html>
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <title>Discord Video Embedder</title>
        <meta charset="UTF-8" />
        <meta name="description" content="" />
    </head>
    <body style="background-color:#181a1b;">
    <p>
	<video style="height:80%%;width:auto"
        controls
        autoplay
        src="%s"
    />
    </p>
    <p style="color:#d8d4cf">Original: <a style="color:#3391ff" href="%s">%s</a></p>
    </body>
</html>
    `, videoUrl, videoUrl, videoUrl)

	// fmt.Fprint(w, videoHtml)
	fmt.Fprintf(w, "%s", videoHtml)
}

func main() {
	router := httprouter.New()
	router.GET("/*video", Video)

	fmt.Printf("Running on port %s\n", PORT)
	fmt.Printf("On host: %s\n", HOST)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
