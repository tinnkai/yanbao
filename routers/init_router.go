package routers

import (
	_ "yanbao/docs"
	middleip "yanbao/middleware/ip"
	middlelogger "yanbao/middleware/logger"
	middlerecover "yanbao/middleware/recover"
	"yanbao/pkg/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化路由
func InitRouter(router *gin.Engine) {
	// 独立中间件相关处理
	middleHandle(router)

	// 获取环境变量
	configorEnv := utils.GetConfigorEnv()

	// swagger dev/test run
	if configorEnv != "pro" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 普通方式路由
	NormalRouter(router)

	// 用户路由
	AccountRouter(router)

	// 订单相关路由
	OrderRouter(router)

	// 注解路由
	AnnotationRouter(router)
}

// 中间件相关处理
func middleHandle(router *gin.Engine) {

	// 恐慌性错误恢复中间件：一般由 panic 引起
	router.Use(middlerecover.ErrorRecover())

	// 日志中间件
	router.Use(middlelogger.Log())

	// IP白名单
	router.Use(middleip.IpWhiteListCheck())
}
