package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zy-free/micro-study/api/member/service"
)

var svc *service.Service

func init() {
	svc = service.New()
}

func InitRouter() *gin.Engine {

	router := gin.Default()
	g := router.Group("/x/api/member")
	g.POST("", createMember)
	g.GET("", getMember)
	//g.POST("/name/update", setName)

	return router
}
