package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/hetao29/qrcode/coding"

	"github.com/SKatiyar/qr/web/resize"
)

var err error
var (
	//addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
	bindaddr = flag.String("b", "0.0.0.0:1023", "listen port")
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
	c, err := coding.Encode(r.FormValue("t"), coding.M)
	if err != nil {
		return
	}
	size, _ := strconv.Atoi(r.FormValue("s"))
	if size < 120 || size > 500 {
		size = 120
	}

	bgcolor, err := ParseHexColor(r.FormValue("bgcolor"))
	if err != nil {
		bgcolor = color.RGBA{0, 0, 0, 0}
	}
	fgcolor, err := ParseHexColor(r.FormValue("fgcolor"))
	if err != nil {
		fgcolor = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	}
	//https://pkg.go.dev/golang.org/x/image/colornames#pkg-variables

	pngdat := c.Image(fgcolor, bgcolor)
	//to do resize
	newImage := resize.Resample(pngdat, image.Rect(0, 0, c.Size, c.Size), size, size)
	png.Encode(w, newImage)
}
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 6:
		_, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(s, "%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
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
