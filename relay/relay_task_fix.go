package relay

import (
	"fmt"

	"github.com/QuantumNous/new-api/common"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
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

// HandleVideoModelRatio 处理视频模型价格比例
func HandleVideoModelRatio(
	c *gin.Context,
	info *relaycommon.RelayInfo,
	ratio float64) (float64, error) {
	bodyMap, err := GetCurrentRequestBodyMap(c)
	if err != nil {
		return 0, err
	}
	fmt.Println("------------------")
	fmt.Println(bodyMap)

	return 0, errors.New("测试")
}
