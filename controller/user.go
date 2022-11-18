package controller

import (
	"bluebell_sly/dao/postgres"
	"bluebell_sly/models"
	"bluebell_sly/pkg/jwt"
	"errors"
	"fmt"
	"net/http"
	"strings"

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
	atoken, rToken, _ := jwt.GenToken(u.UserID)

	ResponseSuccess(c, gin.H{
		"accessToken":  atoken,
		"refreshToken": rToken,
		"userID":       u.UserID,
		"username":     u.UserName,
	})

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

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式错误")
		c.Abort()
		return
	}
	atoken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  atoken,
		"refresh_token": rToken,
	})

}
