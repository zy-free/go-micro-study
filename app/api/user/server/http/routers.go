package server

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine{

	router := gin.Default()
	router.Group("/x/api/member")
	//router.POST("/name/update", )
	router.GET("",member)
	return router
}
