package comereducation

import (
	"gorm.io/gorm"
)

func FindComerEducation(db *gorm.DB, comerId uint64) (*ComerEducation, error) {
	var education ComerEducation
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).First(&education).Error
	if err != nil {
		return nil, err
	}
	return &education, nil
}

func ListComerEducations(db *gorm.DB, comerId uint64) ([]ComerEducation, error) {
	var educations []ComerEducation
	err := db.Where("comer_id = ? AND is_deleted = 0", comerId).Find(&educations).Error
	return educations, err
}

func CreateComerEducation(db *gorm.DB, education *ComerEducation) (uint64, error) {
	err := db.Create(education).Error
	if err != nil {
		return 0, err
	}
	return education.ID, nil
}

func UpdateComerEducation(db *gorm.DB, education *ComerEducation) error {
	return db.Save(education).Error
}

func DeleteComerEducation(db *gorm.DB, id uint64) error {
	return db.Model(&ComerEducation{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

func DeleteComerEducationsByComerId(db *gorm.DB, comerId uint64) error {
	return db.Model(&ComerEducation{}).Where("comer_id = ?", comerId).Update("is_deleted", 1).Error
}
