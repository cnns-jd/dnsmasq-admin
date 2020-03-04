package app

import (
	"github.com/gin-gonic/gin"
	"os"
)

var engine *gin.Engine

// 初始化配置
func init() {
	if engine == nil {
		engine = gin.Default()
	}
	gin.SetMode(gin.DebugMode)
}

// 运行web服务
func Run() {
	engine.Run(os.Getenv("LISTEN_ADDRESS"))
}

// 获取gin web引擎
func Engine() *gin.Engine {
	return engine
}