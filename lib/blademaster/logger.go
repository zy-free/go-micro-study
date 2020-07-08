package blademaster

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerWithConfig(config LoggerConfig) gin.HandlerFunc {
	out := config.Output
	return func(c *gin.Context) {
		now := time.Now()

		ip := c.ClientIP()
		req := c.Request
		method := req.Method
		path := req.URL.Path
		agent := req.UserAgent()

		c.Next()

		latency := time.Since(now)
		// 记录 我们的访问状态
		status := c.Writer.Status()

		// 此处还可以增加一些慢接口日志
		fmt.Fprintf(out, "%s - [%s] \"%s %s %d %s \"%s\" \"\n",
			ip,
			now,
			method,
			path,
			status,
			latency,
			agent)

	}
}
