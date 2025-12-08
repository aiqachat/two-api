package relay

import (
	"fmt"
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
	if result["resolution"] != "" {
		result["rs"] = result["resolution"]
	}
	if result["duration"] != "" {
		result["dur"] = result["duration"]
	}
	if result["rs"] == "" {
		return result, errors.New("视频时长(秒)不能为空")
	}
	if result["dur"] == "" {
		return result, errors.New("视频分辨率不能为空")
	}
	return result, nil
}

// HandleVideoModelRatio 处理视频模型价格比例
func HandleVideoModelRatio(
	c *gin.Context,
	info *relaycommon.RelayInfo,
	ratio float64) (float64, error) {
	videoInfo, err := getVideoInfo(c)
	if err != nil {
		return 0, err
	}
	fmt.Println("------------------1")
	fmt.Println(videoInfo)
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

	return 0, errors.New("测试")
}
