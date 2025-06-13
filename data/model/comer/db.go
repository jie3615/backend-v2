package comer

import (
	"gorm.io/gorm"
)

func FindComer(db *gorm.DB, comerID uint64) (comer *Comer, err error) {
	err = db.Where("id = ?", comerID).First(&comer).Error
	return
}

func FindComerByAddress(db *gorm.DB, address string) (comer *Comer, err error) {
	err = db.Where("address = ? AND is_deleted = 0", address).First(&comer).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func InsertComer(db *gorm.DB, comer *Comer) (err error) {
	err = db.Create(comer).Error
	return
}
