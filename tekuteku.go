package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
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
	Name      string `sql:"size:50"`
	Age       int
	Gender    string `sql:"size:10"`
	StartedAt string `sql:"size:30"`
	Finished  bool
}

type Onom struct {
	Index int
	Name  string
}

type SubData struct {
	UserID        int
	CurrentOnom   Onom
	ShuffledOnoms []string
}

var db gorm.DB
var onoms = []string{"てくてく", "すたすた", "のろのろ"}

func init() {
	db, _ = gorm.Open("sqlite3", "ExperimentA.db")
	db.DB()
	db.AutoMigrate(&User{}, &Experiment{})
}

func main() {
	m := martini.Classic()

	// render html template
	m.Use(render.Renderer())
	m.Get("/", top)
	m.Post("/tutrial", tutrial)
	m.Post("/first", first)
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
	user.StartedAt = time.Now().Format("2006-01-02 15:04:05.999999999")

	db.NewRecord(user)
	db.Create(&user)
	var me User
	db.Where("name = ? and started_at = ?", user.Name, user.StartedAt).First(&me)
	var subData SubData
	subData.UserID = me.ID

	//shuffledOnoms := make([]string, len(onoms))
	//perm := rand.Perm(len(src))
	//for i, v := range perm {
	//  	dest[v] = src[i]
	//}

	subData.ShuffledOnoms = []string{"とんとん", "ぷかぷか", "きらきら"}
	ren.HTML(200, "tutrial", subData)
}

func first(ren render.Render, req *http.Request) {
	var subData SubData
	subData.UserID, _ = strconv.Atoi(req.FormValue("user-id"))
	var onom Onom
	onom.Index = 0
	onom.Name = onoms[0]
	subData.CurrentOnom = onom
	ren.HTML(200, "tekuteku", subData)
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
	experiment.UserID, _ = strconv.Atoi(req.FormValue("user-id"))

	db.NewRecord(experiment)
	db.Create(&experiment)

	fmt.Println(experiment)
	fmt.Println(req.FormValue("user-id"))
	var onom Onom
	currentIndex := experiment.OnomIndex + 1 // インデックス番号を1進める
	if currentIndex < len(onoms) {
		onom.Index = currentIndex
		onom.Name = onoms[currentIndex]

		var subData SubData
		subData.CurrentOnom = onom
		subData.UserID, _ = strconv.Atoi(req.FormValue("user-id"))

		ren.HTML(200, "tekuteku", subData)
	} else {
		fmt.Println("おわり")
		ren.HTML(200, "end", nil)
	}
}
