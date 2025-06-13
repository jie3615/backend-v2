package comereducation

import "metaLand/data/model"

type ComerEducation struct {
	model.Base
	ComerId     uint64 `gorm:"column:comer_id" json:"comer_id"`
	School      string `gorm:"column:school" json:"school"`
	Degree      string `gorm:"column:degree" json:"degree"`
	Major       string `gorm:"column:major" json:"major"`
	StartDate   string `gorm:"column:start_date" json:"start_date"`
	EndDate     string `gorm:"column:end_date" json:"end_date"`
	Description string `gorm:"column:description" json:"description"`
	Level       int    `gorm:"column:level" json:"level"`
}

func (ComerEducation) TableName() string {
	return "comer_education"
}
