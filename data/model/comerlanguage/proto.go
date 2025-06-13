package comerlanguage

import "metaLand/data/model"

type ComerLanguage struct {
	model.Base
	ComerId  uint64 `gorm:"column:comer_id" json:"comer_id"`
	Language string `gorm:"column:language" json:"language"`
	Code     string `gorm:"column:code" json:"code"`
	Level    int    `gorm:"column:level" json:"level"`
	IsNative bool   `gorm:"column:is_native" json:"is_native"`
}

func (ComerLanguage) TableName() string {
	return "comer_language"
}
