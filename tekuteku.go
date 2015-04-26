package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	//"github.com/jinzhu/gorm"
	//_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"time"
)

type ExperimentData struct {
	ID           int
	OnomName     string
	OnomIndex    int
	HijiKakudo   int
	UdeFuriHaba  int
	HizaMageHaba int
	AshiFuriHaba int
	KoshiKakudo  int
	Sokudo       int
	PostedAt     string
}

type UserData struct {
	ID        int
	Name      string
	Age       int
	Gender    string
	StartedAt string
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
	var userData UserData

	userData.Name = req.FormValue("name")
	userData.Age, _ = strconv.Atoi(req.FormValue("age"))
	userData.Gender = req.FormValue("gender")
	userData.StartedAt = time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println(subData.DateTimeStr)
	fmt.Println(userData)
	var onom Onom
	onom.Index = 0
	onom.Name = onoms[0]
	ren.HTML(200, "tekuteku", onom)
}

func sinceSecondPost(ren render.Render, req *http.Request) {
	var experimentData ExperimentData
	experimentData.OnomName = req.FormValue("onom-name")
	experimentData.OnomIndex, _ = strconv.Atoi(req.FormValue("onom-index"))
	experimentData.HijiKakudo, _ = strconv.Atoi(req.FormValue("hiji-kakudo"))
	experimentData.UdeFuriHaba, _ = strconv.Atoi(req.FormValue("ude-furi"))
	experimentData.HizaMageHaba, _ = strconv.Atoi(req.FormValue("hiza-mage"))
	experimentData.AshiFuriHaba, _ = strconv.Atoi(req.FormValue("asi-furi"))
	experimentData.KoshiKakudo, _ = strconv.Atoi(req.FormValue("koshi-kakudo"))
	experimentData.Sokudo, _ = strconv.Atoi(req.FormValue("sokudo"))
	experimentData.PostedAt = time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(experimentData)

	var onom Onom
	presentIndex := experimentData.OnomIndex + 1 // インデックス番号を1進める
	if presentIndex < len(onoms) {
		onom.Index = presentIndex
		onom.Name = onoms[presentIndex]
		ren.HTML(200, "tekuteku", onom)
	} else {
		fmt.Println("おわり")
		ren.HTML(200, "end", nil)
	}
}
