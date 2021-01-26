package routers

import (
	v1_order "yanbao/controllers/v1/order"
	"yanbao/middleware/local_auth"

	"github.com/gin-gonic/gin"
)

// 订单相关路由
func OrderRouter(router *gin.Engine) {
	// Order group: v1
	v1 := router.Group("/order/v1")
	v1.Use(local_auth.CheckLogin())
	{
		var userOrder v1_order.UserOrderController
		v1.POST("/list", userOrder.List)
		v1.POST("/detail", userOrder.Detail)

		var order v1_order.OrderController
		v1.POST("/checkout", order.Checkout)
		v1.POST("/saveorder", order.SaveOrder)
	}
}
