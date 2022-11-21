package controller

import "github.com/gin-gonic/gin"

func getCurrentUserId(c *gin.Context) (userId uint64, err error) {
	_userId, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
	}
	userId = _userId.(uint64)
	return
}
