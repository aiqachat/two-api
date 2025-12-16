package ratio_setting

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/config"
	"github.com/QuantumNous/new-api/types"
)

var (
	groupRatio = map[string]float64{
		"default": 1,
		"vip":     1,
		"svip":    1,
	}
	groupRatioMutex sync.RWMutex
)

var (
	GroupGroupRatio = map[string]map[string]float64{
		"vip": {
			"edit_this": 0.9,
		},
	}
	groupGroupRatioMutex sync.RWMutex
)

var (
	GroupModelRatio = map[string]map[string]float64{
		"vip": {
			"edit_this": 0.9,
		},
	}
	groupModelRatioMutex sync.RWMutex
)

var defaultGroupSpecialUsableGroup = map[string]map[string]string{
	"vip": {
		"append_1":   "vip_special_group_1",
		"-:remove_1": "vip_removed_group_1",
	},
}

type GroupRatioSetting struct {
	GroupRatio              map[string]float64                      `json:"group_ratio"`
	GroupGroupRatio         map[string]map[string]float64           `json:"group_group_ratio"`
	GroupSpecialUsableGroup *types.RWMap[string, map[string]string] `json:"group_special_usable_group"`
}

var groupRatioSetting GroupRatioSetting

func init() {
	groupSpecialUsableGroup := types.NewRWMap[string, map[string]string]()
	groupSpecialUsableGroup.AddAll(defaultGroupSpecialUsableGroup)

	groupRatioSetting = GroupRatioSetting{
		GroupSpecialUsableGroup: groupSpecialUsableGroup,
		GroupRatio:              groupRatio,
		GroupGroupRatio:         GroupGroupRatio,
	}

	config.GlobalConfig.Register("group_ratio_setting", &groupRatioSetting)
}

func GetGroupRatioSetting() *GroupRatioSetting {
	if groupRatioSetting.GroupSpecialUsableGroup == nil {
		groupRatioSetting.GroupSpecialUsableGroup = types.NewRWMap[string, map[string]string]()
		groupRatioSetting.GroupSpecialUsableGroup.AddAll(defaultGroupSpecialUsableGroup)
	}
	return &groupRatioSetting
}

func GetGroupRatioCopy() map[string]float64 {
	groupRatioMutex.RLock()
	defer groupRatioMutex.RUnlock()

	groupRatioCopy := make(map[string]float64)
	for k, v := range groupRatio {
		groupRatioCopy[k] = v
	}
	return groupRatioCopy
}

func ContainsGroupRatio(name string) bool {
	groupRatioMutex.RLock()
	defer groupRatioMutex.RUnlock()

	_, ok := groupRatio[name]
	return ok
}

func GroupRatio2JSONString() string {
	groupRatioMutex.RLock()
	defer groupRatioMutex.RUnlock()

	jsonBytes, err := json.Marshal(groupRatio)
	if err != nil {
		common.SysLog("error marshalling model ratio: " + err.Error())
	}
	return string(jsonBytes)
}

func UpdateGroupRatioByJSONString(jsonStr string) error {
	groupRatioMutex.Lock()
	defer groupRatioMutex.Unlock()

	groupRatio = make(map[string]float64)
	return json.Unmarshal([]byte(jsonStr), &groupRatio)
}

func GetGroupRatio(name string) float64 {
	groupRatioMutex.RLock()
	defer groupRatioMutex.RUnlock()

	ratio, ok := groupRatio[name]
	if !ok {
		common.SysLog("group ratio not found: " + name)
		return 1
	}
	return ratio
}

func GetGroupGroupRatio(userGroup, usingGroup string) (float64, bool) {
	groupGroupRatioMutex.RLock()
	defer groupGroupRatioMutex.RUnlock()

	gp, ok := GroupGroupRatio[userGroup]
	if !ok {
		return -1, false
	}
	ratio, ok := gp[usingGroup]
	if !ok {
		return -1, false
	}
	return ratio, true
}

func GetGroupModelRatio(userGroup, modelName string) (float64, bool) {
	groupModelRatioMutex.RLock()
	defer groupModelRatioMutex.RUnlock()

	gm, ok := GroupModelRatio[userGroup]
	if !ok {
		return -1, false
	}
	ratio, ok := gm[modelName]
	if !ok {
		return -1, false
	}
	return ratio, true
}

type GroupRatioResult struct {
	GroupRatio      float64 `json:"group_ratio"`
	GroupGroupRatio *float64 `json:"group_group_ratio"`
	GroupModelRatio *float64 `json:"group_model_ratio"`
	Result          float64 `json:"result"`
}

// 获取分组倍率统计
func GetGroupRatioResult(userGroup, usingGroup string, modelName string) GroupRatioResult {
	result := GroupRatioResult{
		GroupRatio:      1,
		GroupGroupRatio: nil,
		GroupModelRatio: nil,
		Result:          1,
	}
	// ============================== 获取分组倍率
	result.GroupRatio = GetGroupRatio(userGroup)
	result.Result = result.GroupRatio
	// ============================== 获取分组倍率
	// ============================== 获取用户分组特殊分组倍率
	userGroupRatio, hasUserGroupRatio := GetGroupGroupRatio(userGroup, usingGroup)
	if hasUserGroupRatio {
		result.GroupGroupRatio = &userGroupRatio
		result.Result = userGroupRatio
	}
	// ============================== 获取用户分组特殊分组倍率
	// ============================== 获取用户分组特殊模型倍率
	groupModelRatio, hasGroupModelRatio := GetGroupModelRatio(userGroup, modelName)
	if hasGroupModelRatio {
		result.GroupModelRatio = &groupModelRatio
		result.Result = groupModelRatio
	}
	// ============================== 获取用户分组特殊模型倍率
	return result
}

func GroupGroupRatio2JSONString() string {
	groupGroupRatioMutex.RLock()
	defer groupGroupRatioMutex.RUnlock()

	jsonBytes, err := json.Marshal(GroupGroupRatio)
	if err != nil {
		common.SysLog("error marshalling group-group ratio: " + err.Error())
	}
	return string(jsonBytes)
}

func GroupModelRatio2JSONString() string {
	groupModelRatioMutex.RLock()
	defer groupModelRatioMutex.RUnlock()

	jsonBytes, err := json.Marshal(GroupModelRatio)
	if err != nil {
		common.SysLog("error marshalling group-model ratio: " + err.Error())
	}
	return string(jsonBytes)
}

func UpdateGroupGroupRatioByJSONString(jsonStr string) error {
	groupGroupRatioMutex.Lock()
	defer groupGroupRatioMutex.Unlock()

	GroupGroupRatio = make(map[string]map[string]float64)
	return json.Unmarshal([]byte(jsonStr), &GroupGroupRatio)
}

func UpdateGroupModelRatioByJSONString(jsonStr string) error {
	groupModelRatioMutex.Lock()
	defer groupModelRatioMutex.Unlock()

	GroupModelRatio = make(map[string]map[string]float64)
	return json.Unmarshal([]byte(jsonStr), &GroupModelRatio)
}

func CheckGroupRatio(jsonStr string) error {
	checkGroupRatio := make(map[string]float64)
	err := json.Unmarshal([]byte(jsonStr), &checkGroupRatio)
	if err != nil {
		return err
	}
	for name, ratio := range checkGroupRatio {
		if ratio < 0 {
			return errors.New("group ratio must be not less than 0: " + name)
		}
	}
	return nil
}