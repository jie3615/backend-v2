package comerskill

import (
	"gorm.io/gorm"
)

func FindComerSkill(db *gorm.DB, comerId uint64) (*ComerSkill, error) {
	var skill ComerSkill
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).First(&skill).Error
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func ListComerSkills(db *gorm.DB, comerId uint64) ([]ComerSkill, error) {
	var skills []ComerSkill
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).Find(&skills).Error
	return skills, err
}

func CreateComerSkill(db *gorm.DB, skill *ComerSkill) (uint64, error) {
	err := db.Create(skill).Error
	if err != nil {
		return 0, err
	}
	return skill.ID, nil
}

func UpdateComerSkill(db *gorm.DB, skill *ComerSkill) error {
	return db.Save(skill).Error
}

func DeleteComerSkill(db *gorm.DB, id uint64) error {
	return db.Model(&ComerSkill{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
