package model

import (
	"encoding/json"

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

	if err = query.Unscoped().Order("updated_time desc").Limit(pageInfo.GetPageSize()).Offset(pageInfo.GetStartIdx()).Find(&items).Error; err != nil {
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

func WsVideoRatioGetByModeName(modelName string) (*WsVideoRatio, error) {
	if modelName == "" {
		return nil, errors.New("模型名称不能为空")
	}
	wsVideoRatio := WsVideoRatio{ModelName: modelName}
	var err error = nil
	err = DB.First(&wsVideoRatio, "model_name = ?", modelName).Error
	if err != nil {
		return nil, err
	}
	return &wsVideoRatio, err
}

func WsVideoRatioGetById(id int) (*WsVideoRatio, error) {
	if id == 0 {
		return nil, errors.New("id 为空！")
	}
	wsVideoRatio := WsVideoRatio{Id: id}
	var err error = nil
	err = DB.First(&wsVideoRatio, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &wsVideoRatio, err
}

func WsVideoRatioCreate(modelName string, config map[string]float64) (*WsVideoRatio, error) {
	if modelName == "" {
		return nil, errors.New("模型名称不能为空")
	}
	if config == nil {
		return nil, errors.New("分辨率不能为空")
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
//func (wsVideoRatio *WsVideoRatio) Delete() error {
//	err := DB.Delete(wsVideoRatio).Error
//	return err
//}
