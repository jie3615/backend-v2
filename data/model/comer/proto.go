package comer

import "metaLand/data/model"

type Comer struct {
	model.Base
	Address string `gorm:"address" db:"address"` // comer could save some useful info on block chain with this address
}

// TableName Startup table name for gorm
func (Comer) TableName() string {
	return "comer"
}
