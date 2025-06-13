package comerlanguage

import (
	"gorm.io/gorm"
)

func FindComerLanguage(db *gorm.DB, comerId uint64) (*ComerLanguage, error) {
	var language ComerLanguage
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).First(&language).Error
	if err != nil {
		return nil, err
	}
	return &language, nil
}

func ListComerLanguages(db *gorm.DB, comerId uint64) ([]ComerLanguage, error) {
	var languages []ComerLanguage
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).Find(&languages).Error
	return languages, err
}

func CreateComerLanguage(db *gorm.DB, language *ComerLanguage) (uint64, error) {
	err := db.Create(language).Error
	if err != nil {
		return 0, err
	}
	return language.ID, nil
}

func UpdateComerLanguage(db *gorm.DB, language *ComerLanguage) error {
	return db.Model(&ComerLanguage{}).Where("id = ?", language.ID).Updates(language).Error
}

func DeleteComerLanguage(db *gorm.DB, id uint64) error {
	return db.Model(&ComerLanguage{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

func DeleteComerLanguagesByComerId(db *gorm.DB, comerId uint64) error {
	return db.Model(&ComerLanguage{}).Where("comer_id = ?", comerId).Update("is_deleted", 1).Error
}
