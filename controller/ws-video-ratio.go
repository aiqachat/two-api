package controller

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/gin-gonic/gin"
)

// 视频倍率配置列表
func WsVideoRatioPageList(c *gin.Context) {
	pageInfo := common.GetPageQuery(c)
	modelName := c.Query("model_name")
	items, total, err := model.WsVideoRatioPageList(pageInfo, modelName)
	if err != nil {
		common.ApiError(c, err)
		return
	}

	pageInfo.SetTotal(int(total))
	pageInfo.SetItems(items)

	common.ApiSuccess(c, pageInfo)
	return
}

// 视频倍率支付分辨率列表
func WsVideoRatioResolutionList(c *gin.Context) {
	var items []map[string]any
	for _, item := range model.ResolutionList {
		items = append(items, gin.H{
			"key": item,
			"name": item,
		})
	}
	common.ApiSuccess(c, gin.H{
		"items": items,
	})
}

// 创建视频倍率配置
func WsVideoRatioCreate(c *gin.Context) {
	var params model.WsVideoRatioMap
	if err := common.UnmarshalBodyReusable(c, &params); err != nil {
		common.ApiError(c, err)
		return
	}
	_, err := model.WsVideoRatioCreate(
		params.ModelName, params.Config,
	)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, gin.H{})
}

func WsVideoRatioGetById(c *gin.Context) {
	var params model.WsVideoRatioMap
	if err := common.UnmarshalBodyReusable(c, &params); err != nil {
		common.ApiError(c, err)
		return
	}
	res, err := model.WsVideoRatioGetById(params.Id)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, res)
}

func WsVideoRatioDeleteById(c *gin.Context) {
	var params model.WsVideoRatioMap
	if err := common.UnmarshalBodyReusable(c, &params); err != nil {
		common.ApiError(c, err)
		return
	}
	if err := model.WsVideoRatioDeleteById(params.Id); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, gin.H{})
}
