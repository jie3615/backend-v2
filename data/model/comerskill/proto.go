package comerskill

import "metaLand/data/model"

type ComerSkill struct {
	model.Base
	ComerId     uint64 `gorm:"column:comer_id" json:"comer_id"`
	SkillName   string `gorm:"column:skill_name" json:"skill_name"`
	Level       int    `gorm:"column:level" json:"level"`
	Years       int    `gorm:"column:years" json:"years"`
	Description string `gorm:"column:description" json:"description"`
}

func (ComerSkill) TableName() string {
	return "comer_skill"
}
