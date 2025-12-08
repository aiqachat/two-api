package wsController

import (
	"net/http"

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
