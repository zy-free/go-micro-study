package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zy-free/micro-study/api/member/service/loginlog"
	"github.com/zy-free/micro-study/api/member/service/member"
	bm "github.com/zy-free/micro-study/lib/blademaster"
)

var (
	memberSvc   *member.Service
	loginlogSvc *loginlog.Service
)

func init() {
	memberSvc = member.New()
	loginlogSvc = loginlog.New()
	// 可实现多个svc
}

func InitRouter() *gin.Engine {
	router := bm.EngineNew()
	g := router.Group("/x/api")
	g.Use()
	{
		member := g.Group("/members")
		{
			member.POST("", addMember) // 创建
			member.GET("", getMember)  //查询单个
			//g.GET("/members", queryMember)     //查询列表
			member.PUT("", setMember)       // 修改全部或新增
			member.PATCH("", updateMember)  // 部分字段修改
			member.DELETE("", deleteMember) // 删除
		}
		loginlog := g.Group("/loginlogs")
		{
			loginlog.POST("", addMember) // 创建
		}
	}

	return router
}
