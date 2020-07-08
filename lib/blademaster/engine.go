package blademaster

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"os"
)

func EngineNew() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	webOut, err := os.OpenFile("./log/"+"web.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("web log open fail:%v", err))
	}
	panicOut, err := os.OpenFile("./log/"+"panic.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("panic log open fail:%v", err))
	}
	router.Use(LoggerWithConfig(LoggerConfig{Output: webOut}))
	router.Use(RecoveryWithConfig(LoggerConfig{Output: panicOut}))
	pprof.Register(router, "/debug/pprof/")
	return router
}
