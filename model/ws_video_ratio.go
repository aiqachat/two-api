package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type WsVideoRatio struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ModelName string `json:"model_name" gorm:"type:varchar(255);not null;unique"`
	Config    string `json:"config" gorm:"type:text"`
	//Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedTime int64 `json:"created_time" gorm:"type:bigint;default:null"`
	UpdatedTime int64 `json:"updated_time" gorm:"type:bigint;default:null"`
}

type WsVideoRatioConfigItem struct {
	Name  string  `json:"name"`
	Label string  `json:"label"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type WsVideoRatioMap struct {
	Id          int                      `json:"id"`
	ModelName   string                   `json:"model_name"`
	Config      []WsVideoRatioConfigItem `json:"config"`
	CreatedTime int64                    `json:"created_time"`
	UpdatedTime int64                    `json:"updated_time"`
}

func WsVideoRatio2map(w WsVideoRatio) (WsVideoRatioMap, error) {
	res := WsVideoRatioMap{}
	res.Id = w.Id
	res.ModelName = w.ModelName
	res.CreatedTime = w.CreatedTime
	res.UpdatedTime = w.UpdatedTime
	err := json.Unmarshal([]byte(w.Config), &res.Config)
	for i := range res.Config {
		configItem := res.Config[i]
		foundItem, found := lo.Find(WsVideoRatioInitConfig, func(item WsVideoRatioConfigItem) bool {
			return item.Name == configItem.Name
		})
		if found {
			res.Config[i].Label = foundItem.Label
		}
	}
	if err != nil {
		return res, err
	}
	return res, nil
}

// WsVideoRatioInitConfig 视频倍率初始配置
var WsVideoRatioInitConfig = []WsVideoRatioConfigItem{
	{Name: "480p", Label: "480p分辨率每秒价格", Type: "resolution_price", Value: 1.0},
	{Name: "720p", Label: "720p分辨率每秒价格", Type: "resolution_price", Value: 1.0},
	{Name: "1080p", Label: "1080p分辨率每秒价格", Type: "resolution_price", Value: 1.0},
	// 生成声音倍率
	{Name: "generate_audio_ratio", Label: "生成声音倍率", Type: "generate_audio_ratio", Value: 1.0},
	// 样片倍率
	{Name: "draft_ratio", Label: "样片倍率", Type: "draft_ratio", Value: 1.0},
	// 离线推理模式倍率
	{Name: "service_tier_flex_ratio", Label: "离线推理模式倍率", Type: "service_tier_flex_ratio", Value: 1.0},
}

func WsVideoRatioPageList(
	pageInfo *common.PageInfo, modelName string) (
	list []WsVideoRatioMap, total int64, err error) {
	var items []WsVideoRatio
	list = []WsVideoRatioMap{}
	tx := DB.Begin()
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	query := tx.Model(&WsVideoRatio{})

	if modelName != "" {
		query = query.Where("model_name LIKE ?", "%"+modelName+"%")
	}

	if err = query.Count(&total).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	if err = query.Unscoped().Order("id asc").
		Limit(pageInfo.GetPageSize()).Offset(pageInfo.GetStartIdx()).Find(&items).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	if err = tx.Commit().Error; err != nil {
		return nil, 0, err
	}
	for i := range items {
		m, err := WsVideoRatio2map(items[i])
		if err == nil {
			list = append(list, m)
		}
	}
	return list, total, nil
}

func WsVideoRatioGetByModeName(modelName string) (*WsVideoRatioMap, error) {
	if modelName == "" {
		return nil, errors.New("模型名称不能为空")
	}
	wsVideoRatio := WsVideoRatio{ModelName: modelName}
	var err error = nil
	err = DB.First(&wsVideoRatio, "model_name = ?", modelName).Error
	if err != nil {
		return nil, err
	}
	m, err := WsVideoRatio2map(wsVideoRatio)
	return &m, err
}

func WsVideoRatioGetById(id int) (*WsVideoRatioMap, error) {
	if id == 0 {
		return nil, errors.New("id 为空！")
	}
	wsVideoRatio := WsVideoRatio{Id: id}
	var err error = nil
	err = DB.First(&wsVideoRatio, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	m, err := WsVideoRatio2map(wsVideoRatio)
	return &m, err
}

func WsVideoRatioUpdateConfigById(id int, config []WsVideoRatioConfigItem) error {
	if id == 0 {
		return errors.New("id不能为空！")
	}
	if config == nil {
		return errors.New("倍率配置不能为空")
	}
	var err error = nil
	// 将 config 转换为临时结构体切片，排除 Label 字段
	type TempConfigItem struct {
		Name  string  `json:"name"`
		Type  string  `json:"type"`
		Value float64 `json:"value"`
	}
	tempConfig := make([]TempConfigItem, len(config))
	for i, item := range config {
		tempConfig[i] = TempConfigItem{
			Name:  item.Name,
			Type:  item.Type,
			Value: item.Value,
		}
	}
	// 将config转换为JSON字符串
	configBytes, err := json.Marshal(tempConfig)
	if err != nil {
		return errors.Wrap(err, "无法序列化config为JSON")
	}
	err = DB.Model(&WsVideoRatio{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"config":       string(configBytes),
			"updated_time": common.GetTimestamp(),
		}).Error
	return err
}

func WsVideoRatioCreate(modelName string, config []WsVideoRatioConfigItem) (*WsVideoRatio, error) {
	if modelName == "" {
		return nil, errors.New("模型名称不能为空")
	}
	if config == nil {
		return nil, errors.New("倍率配置不能为空")
	}
	current, err := WsVideoRatioGetByModeName(modelName)
	if current != nil {
		return nil, errors.New("已存在模型'" + modelName + "'的视频配置")
	}
	// 将config转换为JSON字符串
	configBytes, err := json.Marshal(config)
	if err != nil {
		return nil, errors.Wrap(err, "无法序列化config为JSON")
	}
	wsVideoRatio := &WsVideoRatio{
		ModelName:   modelName,
		Config:      string(configBytes),
		UpdatedTime: common.GetTimestamp(),
		CreatedTime: common.GetTimestamp(),
	}
	err = DB.Create(wsVideoRatio).Error
	return wsVideoRatio, err
}

//func (wsVideoRatio *WsVideoRatio) Update() error {
//	wsVideoRatio.UpdatedTime = common.GetTimestamp()
//	err := DB.Model(wsVideoRatio).Updates(wsVideoRatio).Error
//	return err
//}
//

func WsVideoRatioDeleteById(id int) error {
	if id == 0 {
		return errors.New("配置ID不能为空")
	}
	err := DB.Delete(&WsVideoRatio{}, id).Error
	return err
}

func WsVideoRatioFixConfig() error {
	var allItems []WsVideoRatio

	// 查询所有数据记录
	if err := DB.Find(&allItems).Error; err != nil {
		return err
	}
	// 处理查询结果
	for _, item := range allItems {
		if !strings.HasPrefix(strings.TrimSpace(item.Config), "{") {
			continue
		}
		var oldConfig map[string]float64
		err := json.Unmarshal([]byte(item.Config), &oldConfig)
		if err != nil {
			fmt.Printf("解析配置失败: %v\n", err)
			continue
		}
		// 创建与 WsVideoRatioInitConfig 内容一样的副本
		configList := make([]WsVideoRatioConfigItem, len(WsVideoRatioInitConfig))
		copy(configList, WsVideoRatioInitConfig)
		for i := range configList {
			for oldName, oldValue := range oldConfig {
				if configList[i].Name == oldName {
					configList[i].Value = oldValue
				}
			}
		}
		err = WsVideoRatioUpdateConfigById(item.Id, configList)
		if err != nil {
			return err
		}
	}
	return nil
}
