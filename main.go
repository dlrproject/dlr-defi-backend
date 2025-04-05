package main

import (
	"dlr-defi-backend/controllers"
	"net/http"

	docs "dlr-defi-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @Title main方法
func main() {
	//获取数据库对象DB,逻辑在dlr-defi-backend/modules，通过init方法

	//获取router
	r := gin.Default()
	//欢迎页面
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})

	r.GET("/dex/liquidity/findAllTokens", controllers.FindAllTokens)

	r.GET("/dex/liquidity/getMatchAmount", controllers.GetMatchAmount)

	r.GET("/dex/liquidity/getAnotherTokenInfoByTokenId", controllers.GetAnotherTokenInfoByTokenId)

	r.GET("/helloworld", Helloworld)

	//初始化swagger文档
	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.PersistAuthorization(true)))

	//启动项目
	r.Run("0.0.0.0:8080")

}
