package controller

import (
	"bluebell_sly/dao/postgres"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	list, err := postgres.GetCommunityList()
	if err != nil {
		zap.L().Error("查询社区列表失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
}

func CommunityDetailHandler(c *gin.Context) {
	communityId := c.Param("id")
	community, err := postgres.GetCommunityById(communityId)
	if err != nil {
		zap.L().Error("查询社区详情失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, community)
}
