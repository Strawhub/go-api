package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
    engine := gin.Default()
	// router.Static("URL", "静的ファイル格納場所")
engine.Static("/static/styles", "./static/styles")
engine.LoadHTMLGlob("static/index.tmpl")
engine.GET("/", func(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
	"title": "XXX",
})
})
engine.Run(":3000")
}
