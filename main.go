package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// router.Static("URL", "静的ファイル格納場所")

	engine.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("static/*.tmpl") // 相対パスで事前にテンプレートをロード

	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "XXX",
		})
	})
	engine.GET("/data", func(context *gin.Context) {
		context.HTML(http.StatusOK, "db.tmpl", gin.H{
			"title": "追加",
		})
	})
	engine.Run(":3000")
}
