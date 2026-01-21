package relay

import (
	"fmt"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type VideoModelRatioInfo struct {
	ModelName  string  `json:"model_name"`
	Resolution string  `json:"resolution"`
	Duration   int64   `json:"duration"`
	GroupRatio float64 `json:"group_ratio"`
	// 每秒单价
	Price float64 `json:"price"`
	// 有声倍率
	GenerateAudioRatio *float64 `json:"generate_audio_ratio,omitempty"`
	// 样片倍率
	DraftRatio *float64 `json:"draft_ratio,omitempty"`
	// 离线推理模式倍率
	ServiceTierFlexRatio *float64 `json:"service_tier_flex_ratio,omitempty"`
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
		ModelName:  modelName,
		GroupRatio: groupRatio,
	}
	videoTaskInfo, err := adaptor.GetVideoInfo(c)
	if err != nil {
		return videoInfo, err
	}
	videoInfo.Duration = int64(videoTaskInfo.Duration)
	videoInfo.Resolution = videoTaskInfo.Resolution
	resolutionPrice, found := lo.Find(config, func(item model.WsVideoRatioConfigItem) bool {
		return item.Name == videoInfo.Resolution
	})
	if !found {
		return videoInfo, errors.New(fmt.Sprintf("视频计费不支持视频分辨率：%s", videoInfo.Resolution))
	}
	videoInfo.Price = resolutionPrice.Value
	if videoTaskInfo.GenerateAudio {
		generateAudioRatio, foundGenerateAudioRatio := lo.Find(config, func(item model.WsVideoRatioConfigItem) bool {
			return item.Name == "generate_audio_ratio"
		})
		if foundGenerateAudioRatio {
			videoInfo.GenerateAudioRatio = &generateAudioRatio.Value
		}
	}
	if videoTaskInfo.Draft {
		draftRatio, foundDraftRatio := lo.Find(config, func(item model.WsVideoRatioConfigItem) bool {
			return item.Name == "draft_ratio"
		})
		if foundDraftRatio {
			videoInfo.DraftRatio = &draftRatio.Value
		}
	}
	if videoTaskInfo.ServiceTierFlex {
		serviceTierFlexRatio, foundServiceTierFlexRatio := lo.Find(config, func(item model.WsVideoRatioConfigItem) bool {
			return item.Name == "service_tier_flex_ratio"
		})
		if foundServiceTierFlexRatio {
			videoInfo.ServiceTierFlexRatio = &serviceTierFlexRatio.Value
		}
	}
	// =========================================== 获取视频参数
	// =========================================== 计算总价格
	res := decimal.NewFromFloat(videoInfo.GroupRatio)
	res = res.Mul(decimal.NewFromFloat(videoInfo.Price))
	res = res.Mul(decimal.NewFromInt(videoInfo.Duration))
	if videoInfo.GenerateAudioRatio != nil {
		res = res.Mul(decimal.NewFromFloat(*videoInfo.GenerateAudioRatio))
	}
	if videoInfo.DraftRatio != nil {
		res = res.Mul(decimal.NewFromFloat(*videoInfo.DraftRatio))
	}
	if videoInfo.ServiceTierFlexRatio != nil {
		res = res.Mul(decimal.NewFromFloat(*videoInfo.ServiceTierFlexRatio))
	}
	videoInfo.PriceTotal, _ = res.Float64()
	// =========================================== 计算总价格
	// =========================================== 打印日志
	logStr := fmt.Sprintf("视频分辨率: %s; ", videoInfo.Resolution)
	logStr += fmt.Sprintf("视频秒数: %d; ", videoInfo.Duration)
	logStr += fmt.Sprintf("分辨率每秒价格: %.4f; ", videoInfo.Price)
	logStr += fmt.Sprintf("用户分组倍率: %.4f; ", videoInfo.GroupRatio)
	if videoInfo.GenerateAudioRatio != nil {
		logStr += fmt.Sprintf("生成声音倍率: %.4f; ", *videoInfo.GenerateAudioRatio)
	}
	if videoInfo.DraftRatio != nil {
		logStr += fmt.Sprintf("样片倍率: %.4f; ", *videoInfo.DraftRatio)
	}
	if videoInfo.ServiceTierFlexRatio != nil {
		logStr += fmt.Sprintf("离线推理模式倍率: %.4f; ", *videoInfo.ServiceTierFlexRatio)
	}
	logStr += fmt.Sprintf("结果倍率: %.4f; ", videoInfo.PriceTotal)
	println(logStr)
	// =========================================== 打印日志
	return videoInfo, nil
}
