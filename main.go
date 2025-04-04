package main

import (
	"dlr-defi-backend/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func main() {
	dsn := "root:123456@tcp(192.168.31.134:3306)/dlr-backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		print("error")
	}
	test := modules.Test{
		Test: "hi",
	}
	db.Create(&test)

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
