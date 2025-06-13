package comersocial

import "metaLand/data/model"

type ComerSocial struct {
	model.Base
	ComerId    uint64 `gorm:"column:comer_id" json:"comer_id"`
	Platform   string `gorm:"column:platform" json:"platform"`
	Username   string `gorm:"column:username" json:"username"`
	Url        string `gorm:"column:url" json:"url"`
	IsVerified bool   `gorm:"column:is_verified" json:"is_verified"`
}

func (ComerSocial) TableName() string {
	return "comer_social"
}
