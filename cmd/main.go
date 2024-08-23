package main

import (
	"github.com/Strawhub/go-api/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/questions", func(c *gin.Context) {
		// 値を取得する
		questions := model.GetAll()
		// JSONで返す
		c.JSON(200, questions)
	})
	r.GET("/:tag/:num", func(c *gin.Context) {
		tag := c.Param("tag")
		num := c.Param("num")
		question := model.GetBy(tag, num)
		c.JSON(200, question)
	})

	r.Run()
}
