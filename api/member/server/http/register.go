package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zy-free/micro-study/api/member/service"
	bm "github.com/zy-free/micro-study/lib/blademaster"
)

var svc *service.Service

func init() {
	svc = service.New()
}

func InitRouter() *gin.Engine {
	router := bm.EngineNew()
	g := router.Group("/x/api")
	g.Use()
	{
		g.POST("/members", addMember) // 创建
		g.GET("/members", getMember)  //查询单个
		//g.GET("/members", queryMember)     //查询列表
		g.PUT("/members", setMember)       // 修改全部或新增
		g.PATCH("/members", updateMember)  // 部分字段修改
		g.DELETE("/members", deleteMember) // 删除
	}

	return router
}
