package model

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/pkg/errors"
)

type WsVideoRatio struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	ModeName    string  `json:"mode_name" gorm:"type:varchar(255);not null;default:''"`
	Resolution  string  `json:"resolution" gorm:"type:varchar(255);not null;default:''"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedTime int64   `json:"created_time" gorm:"type:bigint;default:null"`
	UpdatedTime int64   `json:"updated_time" gorm:"type:bigint;default:null"`
}

func WsVideoRatioCreate(modeName string, resolution string, price float64) (*WsVideoRatio, error) {
	if modeName == "" {
		return nil, errors.New("模型名称不能为空")
	}
	if resolution == "" {
		return nil, errors.New("分辨率不能为空")
	}
	if price == 0 {
		return nil, errors.New("价格不能为空")
	}
	wsVideoRatio := &WsVideoRatio{
		ModeName:    modeName,
		Resolution:  resolution,
		Price:       price,
		UpdatedTime: common.GetTimestamp(),
		CreatedTime: common.GetTimestamp(),
	}
	err := DB.Create(wsVideoRatio).Error
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
