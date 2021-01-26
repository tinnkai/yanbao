package account

import (
	"net/http"
	"yanbao/controllers"
	"yanbao/pkg/app"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (this *User) GetUserInfo(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}

	userInfo := controllers.GetAuthUserInfo(ctx)

	//appG response
	appG.Response(http.StatusOK, app.SUCCESS, "", userInfo, false)
	return
}
