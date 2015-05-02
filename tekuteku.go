package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

type Experiment struct {
	ID           int
	OnomName     string `sql:"size:20"`
	OnomIndex    int
	HijiKakudo   int
	UdeFuriHaba  int
	HizaMageHaba int
	AshiFuriHaba int
	KoshiKakudo  int
	Sokudo       int
	CreatedAt    time.Time
	UserID       int `sql:"index"`
}

type User struct {
	ID        int
	Name      string `sql:"size:30"`
	Age       int
	Gender    string `sql:"size:10"`
	CreatedAt time.Time
}

type Onom struct {
	Index int
	Name  string
}

var onoms = [...]string{"てくてく", "すたすた", "のろのろ"}
var userID = 1111 // 初期化

func main() {
	m := martini.Classic()

	// render html template
	m.Use(render.Renderer())
	m.Get("/", top)
	m.Post("/tutrial", tutrial)
	m.Get("/first", first)
	m.Post("/since_second", sinceSecond)
	m.Run()

}

func top(ren render.Render, req *http.Request) {
	ren.HTML(200, "top", nil)
}

func tutrial(ren render.Render, req *http.Request) {
	var user User
	user.Name = req.FormValue("name")
	user.Age, _ = strconv.Atoi(req.FormValue("age"))
	user.Gender = req.FormValue("gender")
	user.CreatedAt = time.Now()
	fmt.Println(user)
	userID = 2222 // デバッグ用
	//var onom Onom
	//onom.Index = 0
	//onom.Name = onoms[0]
	ren.HTML(200, "tutrial", nil)
}

func first(ren render.Render, req *http.Request) {
	var onom Onom
	onom.Index = 0
	onom.Name = onoms[0]
	ren.HTML(200, "tekuteku", onom)
}

func sinceSecond(ren render.Render, req *http.Request) {
	var experiment Experiment
	experiment.OnomName = req.FormValue("onom-name")
	experiment.OnomIndex, _ = strconv.Atoi(req.FormValue("onom-index"))
	experiment.HijiKakudo, _ = strconv.Atoi(req.FormValue("hiji-kakudo"))
	experiment.UdeFuriHaba, _ = strconv.Atoi(req.FormValue("ude-furi"))
	experiment.HizaMageHaba, _ = strconv.Atoi(req.FormValue("hiza-mage"))
	experiment.AshiFuriHaba, _ = strconv.Atoi(req.FormValue("asi-furi"))
	experiment.KoshiKakudo, _ = strconv.Atoi(req.FormValue("koshi-kakudo"))
	experiment.Sokudo, _ = strconv.Atoi(req.FormValue("sokudo"))
	experiment.CreatedAt = time.Now()

	fmt.Println(experiment)

	var onom Onom
	presentIndex := experiment.OnomIndex + 1 // インデックス番号を1進める
	if presentIndex < len(onoms) {
		onom.Index = presentIndex
		onom.Name = onoms[presentIndex]
		ren.HTML(200, "tekuteku", onom)
	} else {
		fmt.Println("おわり")
		ren.HTML(200, "end", nil)
	}
}
