package controller

import (
	"net/http"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/gin-gonic/gin"
)

// WsVideoRatioList 视频倍率配置列表
func WsVideoRatioList(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"missing": 333,
		},
	})
}

// WsVideoRatioCreate 创建视频倍率配置
func WsVideoRatioCreate(c *gin.Context) {
	_, err := model.WsVideoRatioCreate(
		c.GetString("modeName"), c.GetString("resolution"), c.GetFloat64("price"),
	)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, gin.H{})
}
