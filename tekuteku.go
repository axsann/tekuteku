package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

type FormVal struct {
	OnomName     string
	OnomIndex    int
	HijiKakudo   int
	UdeFuriHaba  int
	HizaMageHaba int
	AshiFuriHaba int
	KoshiKakudo  int
	Sokudo       int
}

type Onom struct {
	Index int
	Name  string
}

var onoms = [...]string{"てくてく", "すたすた", "のろのろ"}

func main() {

	m := martini.Classic()

	// render html template
	m.Use(render.Renderer())
	m.Get("/", top)
	m.Post("/first_post", firstPost)
	m.Post("/since_second_post", sinceSecondPost)
	m.Run()

}

func top(ren render.Render, req *http.Request) {
	ren.HTML(200, "top", nil)
}

func firstPost(ren render.Render, req *http.Request) {
	var onom Onom
	onom.Index = 0
	onom.Name = onoms[0]
	ren.HTML(200, "tekuteku", onom)
}

func sinceSecondPost(ren render.Render, req *http.Request) {
	var formVal FormVal
	formVal.OnomName = req.FormValue("onom-name")
	formVal.OnomIndex, _ = strconv.Atoi(req.FormValue("onom-index"))
	formVal.HijiKakudo, _ = strconv.Atoi(req.FormValue("hiji-kakudo"))
	formVal.UdeFuriHaba, _ = strconv.Atoi(req.FormValue("ude-furi"))
	formVal.HizaMageHaba, _ = strconv.Atoi(req.FormValue("hiza-mage"))
	formVal.AshiFuriHaba, _ = strconv.Atoi(req.FormValue("asi-furi"))
	formVal.KoshiKakudo, _ = strconv.Atoi(req.FormValue("koshi-kakudo"))
	formVal.Sokudo, _ = strconv.Atoi(req.FormValue("sokudo"))
	fmt.Println(formVal)

	var onom Onom
	presentIndex := formVal.OnomIndex + 1 // インデックス番号を1進める
	if presentIndex < len(onoms) {
		onom.Index = presentIndex
		onom.Name = onoms[presentIndex]
		ren.HTML(200, "tekuteku", onom)
	} else {
		fmt.Println("おわり")
		ren.HTML(200, "end", nil)
	}
}
