package comeraccount

import "gorm.io/gorm"

func FindComerAccount(db *gorm.DB, comerAccountID uint64) (comerAccount *ComerAccount, err error) {
	err = db.Where("id = ? and is_deleted = 0", comerAccountID).First(&comerAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func FindComerAccountByComerID(db *gorm.DB, comerID uint64) (comerAccount *ComerAccount, err error) {
	err = db.Where("comer_id = ? and is_deleted = 0", comerID).First(&comerAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func UpdateComerAccount(db *gorm.DB, comerAccount *ComerAccount) (err error) {
	err = db.Model(&ComerAccount{}).Where("id = ?", comerAccount.ID).Updates(comerAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return
}

func ListComerAccounts(db *gorm.DB, comerID uint64) (comerAccounts []*ComerAccount, err error) {
	err = db.Where("comer_id = ? and is_deleted = 0", comerID).Find(&comerAccounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func DeleteComerAccount(db *gorm.DB, comerAccountID uint64) (err error) {
	err = db.Model(&ComerAccount{}).Where("id = ?", comerAccountID).Update("is_deleted", 1).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return
}
