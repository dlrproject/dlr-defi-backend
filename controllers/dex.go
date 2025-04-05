package controllers

import (
	"dlr-defi-backend/modules"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 拿所有的代币信息
// @Schemes
// @Description 拿所有的代币信息
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} findAllTokens
// @Router /dex/liquidity/findAllTokens [get]
func FindAllTokens(ctx *gin.Context) {
	allTokens := []modules.Tokens{}
	modules.DB.Find(&allTokens)
	ctx.JSON(200, gin.H{
		"result": allTokens,
	})
}

// @Summary 拿代币对信息
// @Schemes
// @Description 拿代币对信息
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} getMatchAmount
// @Router /dex/liquidity/getMatchAmount [get]
// @Param tokenId0 query string true "tokenId0"
// @Param tokenId1 query string true "tokenId1"
func GetMatchAmount(ctx *gin.Context) {
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
	modules.DB.Where("(token0_id=? and token1_id=?) or (token0_id=? and token1_id=?)", tokenIdInt0, tokenIdInt1, tokenIdInt1, tokenIdInt0).Find(&pairs)

	ctx.JSON(200, gin.H{
		"result": pairs,
	})
}

// @Summary 通过代币id拿 与之相对的 另外一个代币的信息
// @Schemes
// @Description 通过代币id拿 与之相对的 另外一个代币的信息
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} getAnotherTokenInfoByTokenId
// @Router /dex/liquidity/getAnotherTokenInfoByTokenId [get]
// @Param tokenId query string true "tokenId"
func GetAnotherTokenInfoByTokenId(ctx *gin.Context) {
	tokenId := ctx.Query("tokenId")
	tokenIdInt, err := strconv.ParseUint(tokenId, 10, 64)
	if err != nil {
		return
	}
	pairs := []modules.Pairs{}
	modules.DB.Where("token0_id=? or token1_id=?", tokenIdInt, tokenIdInt).Find(&pairs)

	ids := []uint64{}
	for _, pair := range pairs {
		ids = append(ids, pair.Token1Id)
		ids = append(ids, pair.Token0Id)
	}

	allTokens := []modules.Tokens{}
	modules.DB.Where("id in(?)", ids).Find(&allTokens)
	allTokens = DedupSlice(allTokens)
	allTokensAfterRemove := []modules.Tokens{}
	for _, token := range allTokens {
		if token.Id != tokenIdInt {
			allTokensAfterRemove = append(allTokensAfterRemove, token)
		}
	}

	ctx.JSON(200, gin.H{
		"result": allTokensAfterRemove,
	})
}

func DedupSlice[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
