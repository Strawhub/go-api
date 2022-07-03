package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
    engine := gin.Default()
    engine.LoadHTMLGlob("static/index.tmpl")

    engine.GET("/", func(context *gin.Context) {
        context.HTML(http.StatusOK, "index.tmpl", gin.H{
                    "title": "Hello Gin!",
                })
    })

    engine.Run(":3000")
}
