package relay

import (
	"fmt"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type VideoModelRatioInfo struct {
	ModelName  string  `json:"model_name"`
	Resolution string  `json:"resolution"`
	Duration   int64   `json:"duration"`
	GroupRatio float64 `json:"group_ratio"`
	// 每秒单价
	Price float64 `json:"price"`
	// 总价
	PriceTotal float64 `json:"price_total"`
}

// HandleVideoModelRatio 处理视频模型价格比例
func HandleVideoModelRatio(
	c *gin.Context,
	adaptor channel.TaskAdaptor,
	info *relaycommon.RelayInfo,
	groupRatio float64) (*VideoModelRatioInfo, error) {
	// =========================================== 获取视频配置
	modelName := info.OriginModelName
	if modelName == "" {
		return nil, errors.New("未获取到模型名称")
	}
	item, err := model.WsVideoRatioGetByModeName(modelName)
	if err != nil || item == nil {
		return nil, err
	}
	config := item.Config
	// =========================================== 获取视频配置
	// =========================================== 获取视频参数
	videoInfo := &VideoModelRatioInfo{
		ModelName: modelName,
		GroupRatio: groupRatio,
	}
	videoTaskInfo, err := adaptor.GetVideoInfo(c)
	if err != nil {
		return videoInfo, err
	}
	videoInfo.Duration = int64(videoTaskInfo.Duration)
	videoInfo.Resolution = videoTaskInfo.Resolution
	price, ok := config[videoInfo.Resolution]
	if !ok {
		return videoInfo, errors.New(fmt.Sprintf("视频计费不支持视频分辨率：%s", videoInfo.Resolution))
	}
	videoInfo.Price = price
	// =========================================== 获取视频参数
	videoInfo.PriceTotal = videoInfo.Price * float64(videoInfo.Duration) * videoInfo.GroupRatio
	res := decimal.NewFromFloat(videoInfo.GroupRatio)
	res = res.Mul(decimal.NewFromFloat(videoInfo.Price))
	res = res.Mul(decimal.NewFromInt(videoInfo.Duration))
	videoInfo.PriceTotal, _ = res.Float64()
	println(fmt.Sprintf(
		"视频分辨率: %s, 视频秒数: %d, 分辨率每秒价格: %.4f, 用户分组倍率: %.4f, 结果倍率: %.4f",
		videoInfo.Resolution, videoInfo.Duration, videoInfo.Price, videoInfo.GroupRatio, videoInfo.PriceTotal,
	))
	return videoInfo, nil
}