package main

import "github.com/skatiyar/qr"
import "github.com/skatiyar/qr/web/resize"
import (
	"flag"
	"image"
	"image/png"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var err error
var (
	//addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
	bindaddr = flag.String("b", "0.0.0.0:80", "listen port")
)

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	c, err := qr.Encode(r.FormValue("t"), qr.M)
	if err != nil {
	}
	size, _ := strconv.Atoi(r.FormValue("s"))
	if size < 120 || size > 500 {
		size = 120
	}
	pngdat := c.Image()
	//to do resize
	newImage := resize.Resample(pngdat, image.Rect(0, 0, c.Size, c.Size), size, size)
	png.Encode(w, newImage)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", makeHandler(viewHandler))
	err := http.ListenAndServe(*bindaddr, nil)
	if err != nil {
		log.Println("Error:")
		log.Println(err)
		log.Fatal("Exit!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}
}
