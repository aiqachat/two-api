package relay

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/QuantumNous/new-api/common"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/setting/ratio_setting"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

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

func getVideoInfo(c *gin.Context) (map[string]string, error) {
	var result map[string]string
	bodyMap, err := GetCurrentRequestBodyMap(c)
	if err != nil {
		return result, err
	}
	str := bodyMap["prompt"]
	if str == nil {
		return result, errors.New("未获取到prompt内容")
	}
	result = parseVideoParamsSplit(str.(string))
	if result["rs"] != "" {
		result["resolution"] = result["rs"]
	}
	if result["dur"] != "" {
		result["duration"] = result["dur"]
	}
	if _, ok := result["resolution"]; !ok {
		return result, errors.New("视频分辨率不能为空")
	}
	if _, ok := result["duration"]; !ok {
		return result, errors.New("视频时长(秒)不能为空")
	}
	return result, nil
}

// 视频模型价格倍率
var videoModelRatioMap = map[string]float64{
	"480p":  0.05,
	"720p":  0.10,
	"1080p": 0.15,
}

// HandleVideoModelRatio 处理视频模型价格比例
func HandleVideoModelRatio(
	c *gin.Context,
	info *relaycommon.RelayInfo,
	ratio float64) (float64, error) {
	// =========================================== 获取视频参数
	videoInfo, err := getVideoInfo(c)
	if err != nil {
		return 0, err
	}
	resolution := videoInfo["resolution"]
	duration, err := strconv.Atoi(videoInfo["duration"])
	if err != nil {
		return 0, err
	}
	resolutionRatio, ok := videoModelRatioMap[resolution]
	if !ok {
		return 0, errors.New(fmt.Sprintf("不支持的视频分辨率：%s", resolution))
	}
	// =========================================== 获取视频参数
	modelName := info.OriginModelName
	if modelName == "" {
		return 0, errors.New("未获取到模型名称")
	}
	// =========================================== 获取用户分组倍率
	userGroupRatio, hasUserGroupRatio := ratio_setting.GetGroupGroupRatio(info.UserGroup, info.UsingGroup)
	if !hasUserGroupRatio {
		userGroupRatio = ratio_setting.GetGroupRatio(info.UsingGroup)
	}
	fmt.Println(userGroupRatio)
	// =========================================== 获取用户分组倍率
	resultRatio := resolutionRatio * float64(duration) * userGroupRatio
	println(fmt.Sprintf("视频分辨率: %s, 视频秒数: %d, 分辨率每秒价格: %.4f, 结果倍率: %.4f", resolution, duration, resolutionRatio, resultRatio))
	return resultRatio, nil
}
