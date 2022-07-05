package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
    engine := gin.Default()
	// router.Static("URL", "静的ファイル格納場所")
	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("static/index.tmpl")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "XXX",
		})
	})
	engine.GET("/hello", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "japan",
		})
	})
	engine.Run(":3000")
}