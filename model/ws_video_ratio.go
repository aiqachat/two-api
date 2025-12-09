package model

import (
	"encoding/json"

	"github.com/QuantumNous/new-api/common"
	"github.com/pkg/errors"
)

type WsVideoRatio struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ModeName string `json:"mode_name" gorm:"type:varchar(255);not null;unique"`
	Config   string `json:"config" gorm:"type:text"`
	//Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedTime int64 `json:"created_time" gorm:"type:bigint;default:null"`
	UpdatedTime int64 `json:"updated_time" gorm:"type:bigint;default:null"`
}

// ResolutionList 支持的分辨率列表
var ResolutionList = []string{
	"1080p",
	"720p",
	"480p",
}

func WsVideoRatioCreate(modeName string, config map[string]float64) (*WsVideoRatio, error) {
	if modeName == "" {
		return nil, errors.New("模型名称不能为空")
	}
	if config == nil {
		return nil, errors.New("分辨率不能为空")
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
		ModeName:    modeName,
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
