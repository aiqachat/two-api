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

type WsVideoRatioCreateParams struct {
	ModeName string             `json:"mode_name"`
	Config   map[string]float64 `json:"config"`
}

// 视频倍率支付分辨率列表
func WsVideoRatioResolutionList(c *gin.Context) {
	common.ApiSuccess(c, gin.H{
		"list": model.ResolutionList,
	})
}

// 创建视频倍率配置
func WsVideoRatioCreate(c *gin.Context) {
	var params WsVideoRatioCreateParams
	if err := common.UnmarshalBodyReusable(c, &params); err != nil {
		common.ApiError(c, err)
		return
	}
	_, err := model.WsVideoRatioCreate(
		params.ModeName, params.Config,
	)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, gin.H{})
}
