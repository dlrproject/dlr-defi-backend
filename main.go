package main

import "github.com/gin-gonic/gin"

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})
	aa := &Article{
		Title: "1",
		Desc:  "2",
	}
	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(200, aa)
	})

	//创建流动性
	r.POST("/liquidity/create", func(ctx *gin.Context) {
		ctx.JSON(200, aa)
	})

	//添加流动性
	r.GET("/liquidity/add", func(ctx *gin.Context) {
		ctx.JSON(200, aa)
	})
	r.Run()

}
