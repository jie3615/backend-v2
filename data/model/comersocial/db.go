package comersocial

import (
	"gorm.io/gorm"
)

func FindComerSocial(db *gorm.DB, comerId uint64) (*ComerSocial, error) {
	var social ComerSocial
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).First(&social).Error
	if err != nil {
		return nil, err
	}
	return &social, nil
}

func ListComerSocials(db *gorm.DB, comerId uint64) ([]ComerSocial, error) {
	var socials []ComerSocial
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).Find(&socials).Error
	return socials, err
}

func CreateComerSocial(db *gorm.DB, social *ComerSocial) (uint64, error) {
	err := db.Create(social).Error
	if err != nil {
		return 0, err
	}
	return social.ID, nil
}

func UpdateComerSocial(db *gorm.DB, social *ComerSocial) error {
	return db.Save(social).Error
}

func DeleteComerSocial(db *gorm.DB, id uint64) error {
	return db.Model(&ComerSocial{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
