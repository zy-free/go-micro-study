package server

import "github.com/gin-gonic/gin"

func member(ctx *gin.Context){
	ctx.JSON(200,"member test")
}