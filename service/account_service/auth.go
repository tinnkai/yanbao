package account_service

import (
	"yanbao/models"
	"yanbao/pkg/app"
	"yanbao/pkg/errors"
)

// 验证登录
func CheckLogin(userId int64) (models.AuthUserInfo, error) {
	authUserInfo := models.AuthUserInfo{}
	userInfo, err := models.UserRepository.GetLoginUserInfoById(userId, "id,username,password,phone,`group`")
	if err != nil {
		return authUserInfo, err
	}

	// 用户id是否一致
	if userInfo.Id != userId || userInfo.Id < 1 {
		return authUserInfo, errors.Newf(app.ERROR_LOGIN_FAIL, "", "")
	}

	authUserInfo = models.AuthUserInfo{
		UserId:   userInfo.Id,
		Username: userInfo.Username,
		Phone:    userInfo.Phone,
	}

	return authUserInfo, nil
}
