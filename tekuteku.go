package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/go-sql-driver/mysql"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// 実験データ
type Experiment struct {
	ID           int    // gormによってオートインクリメントでDBに格納される
	OnomName     string `sql:"size:20"` // オノマトペ名
	OnomIndex    int    // オノマトペのインデックス番号
	HijiKakudo   int
	UdeFuriHaba  int
	HizaMageHaba int
	AshiFuriHaba int
	KoshiKakudo  int
	Sokudo       int
	Feeling      int
	CreatedAt    time.Time // gormによって自動で作成日時がDBに格納される
	UserID       int       `sql:"index"` // 被験者ID(ユーザID)
}

// 被験者データ
type User struct {
	ID        int    // 被験者ID(ユーザID)
	Name      string `sql:"size:50"` // 被験者の名前
	Age       int    // 被験者の年齢
	Gender    string `sql:"size:10"` // 被験者の性別
	StartedAt string `sql:"size:30"` // 実験開始日時
	Finished  bool   // 終了したかどうか
}

type Feeling struct {
	ID   int
	Text string `sql:"size:10"`
}

// オノマトペを格納する
type Onom struct {
	Index int    // インデックス番号
	Name  string // オノマトペ名
}

// HTMLのフォームに保持するデータ
type SubData struct {
	UserID              int      // 被験者ID(ユーザID)
	CurrentOnom         Onom     // 現在のオノマトペ
	ShuffledOnomStrings []string // シャッフルしたオノマトペの配列
	FeelStrings         []string // 6段階アンケート用の文字列の配列
}

var db gorm.DB // グローバル変数としてgorm.DBを宣言
// 実験に用いるオノマトペ
var onomStrings = []string{"てくてく", "すたすた", "のろのろ", "とぼとぼ"}
var feelStrings = []string{"全くそう思わない", "そう思わない", "あまりそう思わない", "ややそう思う", "そう思う", "とてもそう思う"}

// DBと接続、DBのMigrate(自動設定)
func init() {
	db, _ = gorm.Open("sqlite3", "ExperimentA.db")
	db.DB()
	db.AutoMigrate(&User{}, &Experiment{}, &Feeling{})
	setupFeelingTable()
}

func main() {
	m := martini.Classic()
	// render html template
	m.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{
			{
				"add": func(a, b int) int { return a + b },
			},
		},
	}))
	m.Get("/", top)                      // トップページ(被験者データの入力ページ)
	m.Post("/tutrial", tutrial)          // 動作確認ページ
	m.Post("/first", first)              // 最初のオノマトペの実験ページ
	m.Post("/since_second", sinceSecond) // 2番目以降のオノマトペの実験ページ、終了ページ
	m.Get("/show", show)
	m.Run()
}

func setupFeelingTable() {
	var feeling Feeling
	if db.First(&feeling).RecordNotFound() { // テーブルが空ならばデータを追加する
		for i := 0; i < len(feelStrings); i++ {
			feeling.ID = i + 1
			feeling.Text = feelStrings[i]
			db.NewRecord(feeling)
			db.Create(&feeling)
		}
	}
}

// トップページ(被験者データの入力ページ)
func top(ren render.Render, req *http.Request) {
	ren.HTML(200, "top", nil)
}

// 動作確認ページ
func tutrial(ren render.Render, req *http.Request) {
	// 被験者データをフォームから取得
	var user User
	user.Name = req.FormValue("name")
	user.Age, _ = strconv.Atoi(req.FormValue("age"))
	user.Gender = req.FormValue("gender")
	user.StartedAt = time.Now().Format("2006-01-02 15:04:05.999999999")
	// DBのusersテーブルに被験者データを格納
	db.NewRecord(user)
	db.Create(&user)
	var me User
	// ユーザID取得のために格納した被験者データをDBから再取得し、meに格納
	db.Where("name = ? and started_at = ?", user.Name, user.StartedAt).First(&me)
	var subData SubData
	subData.UserID = me.ID // subDataにユーザIDをセット

	// オノマトペをシャッフルする
	shuffledOnomStrings := make([]string, len(onomStrings))
	perm := rand.Perm(len(onomStrings))
	for i, v := range perm {
		shuffledOnomStrings[v] = onomStrings[i]
	}

	subData.ShuffledOnomStrings = shuffledOnomStrings // subDataにシャッフルしたオノマトペを格納
	ren.HTML(200, "tutrial", subData)
}

func first(ren render.Render, req *http.Request) {
	req.ParseForm()                                            // req.Formの取得のためにパース
	shuffledOnomStrings := req.Form["shuffled-onom-strings[]"] // シャッフルしたオノマトペを取得
	var subData SubData
	subData.UserID, _ = strconv.Atoi(req.FormValue("user-id")) // subDataにユーザIDを格納
	subData.ShuffledOnomStrings = shuffledOnomStrings          // subDataにシャッフルしたオノマトペを格納
	subData.FeelStrings = feelStrings
	var onom Onom
	onom.Index = 0                     // 最初のオノマトペなので、インデックス番号は0
	onom.Name = shuffledOnomStrings[0] // シャッフルしたオノマトペの最初のもの
	subData.CurrentOnom = onom         // subDataにオノマトペを格納
	ren.HTML(200, "tekuteku", subData)
}

func sinceSecond(ren render.Render, req *http.Request) {
	req.ParseForm()                                            // req.Formの取得のためにパース
	shuffledOnomStrings := req.Form["shuffled-onom-strings[]"] // シャッフルしたオノマトペを取得
	// 実験データをフォームから取得
	var experiment Experiment
	experiment.OnomName = req.FormValue("onom-name")
	experiment.OnomIndex, _ = strconv.Atoi(req.FormValue("onom-index"))
	experiment.HijiKakudo, _ = strconv.Atoi(req.FormValue("hiji-kakudo"))
	experiment.UdeFuriHaba, _ = strconv.Atoi(req.FormValue("ude-furi"))
	experiment.HizaMageHaba, _ = strconv.Atoi(req.FormValue("hiza-mage"))
	experiment.AshiFuriHaba, _ = strconv.Atoi(req.FormValue("asi-furi"))
	experiment.KoshiKakudo, _ = strconv.Atoi(req.FormValue("koshi-kakudo"))
	experiment.Sokudo, _ = strconv.Atoi(req.FormValue("sokudo"))
	experiment.Feeling, _ = strconv.Atoi(req.FormValue("feeling"))
	userID, _ := strconv.Atoi(req.FormValue("user-id"))
	experiment.UserID = userID

	// DBのexperimentsテーブルに実験データを格納
	db.NewRecord(experiment)
	db.Create(&experiment)

	var onom Onom
	currentIndex := experiment.OnomIndex + 1 // インデックス番号を1進める
	if currentIndex < len(onomStrings) {     // オノマトペの数だけ実験ページを表示する
		// 今回表示するオノマトペを格納
		onom.Index = currentIndex
		onom.Name = shuffledOnomStrings[currentIndex]

		var subData SubData
		subData.CurrentOnom = onom                        // subDataにオノマトペを格納
		subData.UserID = userID                           // subDataにユーザIDを格納
		subData.ShuffledOnomStrings = shuffledOnomStrings // subDataにシャッフルしたオノマトペを格納
		subData.FeelStrings = feelStrings

		ren.HTML(200, "tekuteku", subData)
	} else { // すべてのオノマトペが終わったら終了ページを表示する
		fmt.Println("おわり")
		var me User
		db.Where("id = ?", userID).First(&me) // ユーザIDで自分の被験者データを取得
		me.Finished = true                    // 最後まで実験をしたら、被験者データのFinishedをtrueにする
		db.Save(&me)                          // 再保存(Update)する
		ren.HTML(200, "end", nil)
	}
}

// テスト用
func show(ren render.Render) {
	var me User
	db.First(&me, 2) // 2番目の被験者を取得
	fmt.Println(me)
	var experiments []Experiment
	db.Model(&me).Related(&experiments) // 2番目の被験者の実験データ(複数)を取得
	fmt.Println(experiments)
	ren.JSON(200, experiments) // JSONでレンダリング
}
