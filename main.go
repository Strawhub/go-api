package main

import "github.com/gin-gonic/gin"
import "net/http"


func main() {
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	engine.Run(":3000")
}
