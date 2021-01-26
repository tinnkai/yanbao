package main

import (
	"fmt"
	"net/http"
	"yanbao/models"
	"yanbao/pkg/gredis"
	"yanbao/pkg/logging"
	"yanbao/pkg/setting"
	"yanbao/routers"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func init() {
	// 读取配置文件
	setting.Setup()
	// mysql model
	models.Setup()
	// 日志
	logging.Setup()
	// redis
	gredis.Setup()
}

func main() {

	// 设置运行模式
	gin.SetMode(setting.ServerSetting.RunMode)
	// 设置读取超时时间
	readTimeout := setting.ServerSetting.ReadTimeout
	// 设置写入超时时间
	writeTimeout := setting.ServerSetting.WriteTimeout
	// 端口
	httpPort := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	// 请求头的最大字节数
	maxHeaderBytes := 1 << 20

	// 初始化 gin
	router := gin.Default()

	// pprof debug run
	if setting.ServerSetting.RunMode == "debug" {
		pprof.Register(router)
	}

	// 初始化路由
	routers.InitRouter(router)

	// 服务相关配置
	server := &http.Server{
		Addr:           httpPort,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	server.ListenAndServe()
}
