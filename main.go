package main

import (
	"html/template"
	"log"
	"net/http"
	// "github.com/gin-gonic/gin"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.tmpl")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "トップページです！")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// func main() {
//     // ルートへのリクエストを"static"ディレクトリ内のtmplファイルで処理する
//     http.Handle("/", http.FileServer("static/*.tmpl"))

//     // localhost:8080でサーバー処理開始
//     http.ListenAndServe(":8080", nil)
// }

// func gtplHandler(w http.ResponseWriter, r *http.Request) {
// 	// テンプレートをパースする
//     engine := template.Must(template.ParseFiles("static/*.tmpl"))
// 	// engine.Static("URL", "静的ファイル格納場所")
// 	engine.Static("/static", "./static")

// 	if err := engine.ExecuteTemplate(w, "index.tmpl"); err != nil {

//     }
// }

// func main() {
//     // ルートへのリクエストを"gtplHandler"関数で処理する
//     http.HandleFunc("/", gtplHandler)

//     // localhost:8080でサーバー処理開始
//     http.ListenAndServe(":3000", nil)
// }

// func main() {
// 	engine := gin.Default()

// 	engine.SetFuncMap(template.FuncMap{
// 		"upper": strings.ToUpper,
// 	})
// 	// engine.Static("URL", "静的ファイル格納場所")
// 	engine.Static("/static", "./static")
// 	engine.LoadHTMLGlob("static/*.tmpl") // 相対パスで事前にテンプレートをロード

// 	engine.GET("/", func(context *gin.Context) {
// 		context.HTML(http.StatusOK, "index.tmpl", gin.H{
// 			"title": "XXX",
// 		})
// 	})
// 	engine.GET("/data", func(context *gin.Context) {
// 		context.HTML(http.StatusOK, "db.tmpl", gin.H{
// 			"title": "追加",
// 		})
// 	})
// 	engine.GET("/data/result", func(context *gin.Context) {
// 		context.HTML(http.StatusOK, "result.tmpl", gin.H{
// 			"title": "追加",
// 		})
// 	})
// 	engine.Run(":3000")
// }
