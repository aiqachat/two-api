package relay

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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

func GetCurrentRequestBodyMap(c *gin.Context) (map[string]interface{}, error) {
	body, err := common.GetRequestBody(c)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, errors.New("未获取到body内容")
	}
	// 解析请求体为JSON
	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		return nil, errors.New("body内容解析失败")
	}
	return requestData, nil
}

func parseVideoParamsSplit(paramStr string) map[string]string {
	result := make(map[string]string)
	// 按空格分割
	parts := strings.Fields(paramStr)

	for i := 0; i < len(parts)-1; i++ {
		if strings.HasPrefix(parts[i], "--") && !strings.HasPrefix(parts[i+1], "--") {
			key := strings.TrimPrefix(parts[i], "--")
			value := parts[i+1]
			result[key] = value
			i++ // 跳过已处理的value
		}
	}
	return result
}

func loadVideoInfo(c *gin.Context, info *VideoModelRatioInfo) error {
	var result map[string]string
	bodyMap, err := GetCurrentRequestBodyMap(c)
	if err != nil {
		return err
	}
	str := bodyMap["prompt"]
	if str == nil {
		return errors.New("未获取到prompt内容")
	}
	result = parseVideoParamsSplit(str.(string))
	if result["rs"] != "" {
		result["resolution"] = result["rs"]
	}
	if result["dur"] != "" {
		result["duration"] = result["dur"]
	}
	if _, ok := result["resolution"]; !ok {
		return errors.New("视频分辨率不能为空")
	}
	if _, ok := result["duration"]; !ok {
		return errors.New("视频时长(秒)不能为空")
	}
	duration, err := strconv.Atoi(result["duration"])
	if err != nil {
		return err
	}
	info.Resolution = result["resolution"]
	info.Duration = int64(duration)
	return nil
}

// HandleVideoModelRatio 处理视频模型价格比例
func HandleVideoModelRatio(
	c *gin.Context,
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
	err = loadVideoInfo(c, videoInfo)
	if err != nil {
		return videoInfo, err
	}
	price, ok := config[videoInfo.Resolution]
	if !ok {
		return videoInfo, errors.New(fmt.Sprintf("不支持的视频分辨率：%s", videoInfo.Resolution))
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