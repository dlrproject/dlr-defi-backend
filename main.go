package main

import (
	"dlr-defi-backend/modules"

	"strconv"

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

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})

	r.GET("/dex/liquidity/findAllTokens", func(ctx *gin.Context) {
		allTokens := []modules.Tokens{}
		db.Find(&allTokens)
		ctx.JSON(200, gin.H{
			"result": allTokens,
		})
	})

	r.GET("/dex/liquidity/getMatchAmount", func(ctx *gin.Context) {
		tokenId0 := ctx.Query("tokenId0")
		tokenIdInt0, err0 := strconv.ParseUint(tokenId0, 10, 64)
		if err0 != nil {
			return
		}
		tokenId1 := ctx.Query("tokenId1")
		tokenIdInt1, err1 := strconv.ParseUint(tokenId1, 10, 64)
		if err1 != nil {
			return
		}

		pairs := []modules.Pairs{}
		db.Where("token0_id=? and token1_id=?", tokenIdInt0, tokenIdInt1).Find(&pairs)

		ctx.JSON(200, gin.H{
			"result": pairs,
		})
	})

	r.GET("/dex/liquidity/getAnotherTokenInfoByTokenId", func(ctx *gin.Context) {
		tokenId := ctx.Query("tokenId")
		tokenIdInt, err := strconv.ParseUint(tokenId, 10, 64)
		if err != nil {
			return
		}
		pairs := []modules.Pairs{}
		db.Where("token0_id=?", tokenIdInt).Find(&pairs)

		ids := []uint64{}
		for _, pair := range pairs {
			ids = append(ids, pair.Token1Id)
		}

		allTokens := []modules.Tokens{}
		db.Where("id in(?)", ids).Find(&allTokens)
		ctx.JSON(200, gin.H{
			"result": allTokens,
		})
	})

	r.Run("0.0.0.0:8080")

}
