package model

import (
	"encoding/json"
	"fmt"

	"github.com/QuantumNous/new-api/common"
	"github.com/pkg/errors"
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
	Id          int                `json:"id"`
	ModelName   string             `json:"model_name"`
	Config      map[string]float64 `json:"config"`
	CreatedTime int64              `json:"created_time"`
	UpdatedTime int64              `json:"updated_time"`
}

func WsVideoRatio2map(w WsVideoRatio) (WsVideoRatioMap, error) {
	res := WsVideoRatioMap{}
	res.Id = w.Id
	res.ModelName = w.ModelName
	res.CreatedTime = w.CreatedTime
	res.UpdatedTime = w.UpdatedTime
	err := json.Unmarshal([]byte(w.Config), &res.Config)
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
	// 有声倍率
	{Name: "audio_ratio", Label: "有声倍率", Type: "audio_ratio", Value: 1.0},
	// 样片倍率
	{Name: "draft_ratio", Label: "样片倍率", Type: "audio_ratio", Value: 1.0},
	// 离线推理模式倍率
	{Name: "service_tier_flex_ratio", Label: "离线推理模式倍率", Type: "service_tier_flex_ratio", Value: 1.0},
}

// ResolutionList 支持的分辨率列表
var ResolutionList = []string{
	"1080p",
	"720p",
	"480p",
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

func WsVideoRatioUpdateConfigById(id int, config map[string]float64) error {
	if id == 0 {
		return errors.New("id不能为空！")
	}
	if config == nil {
		return errors.New("倍率配置不能为空")
	}
	var err error = nil
	// 将config转换为JSON字符串
	configBytes, err := json.Marshal(config)
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

func WsVideoRatioCreate(modelName string, config map[string]float64) (*WsVideoRatio, error) {
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
	for _, resolution := range ResolutionList {
		if _, ok := config[resolution]; !ok {
			return nil, errors.New("请填写" + resolution + "的倍率")
		}
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
		fmt.Printf("ID: %d, ModelName: %s, Config: %s\n",
			item.Id, item.ModelName, item.Config)
	}
	return nil
}
