package controller

import (
	"bluebell_sly/dao/postgres"
	"bluebell_sly/models"
	"errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("invaild params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if err := postgres.Login(&u); err != nil {
		zap.L().Error("login failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, nil)

}

func SignUpHandler(c *gin.Context) {
	var fo models.RegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err := postgres.Register(&models.User{
		UserName: fo.UserName,
		Password: fo.Password,
	})
	if errors.Is(err, postgres.ErrorUserExit) {
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		zap.L().Error("postgres regiseter() failed ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
