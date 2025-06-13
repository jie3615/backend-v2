package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"metaLand/data/utility"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type PageData struct {
	List  []interface{} `json:"list"`
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Total int           `json:"total"`
}

type IsExistResponse struct {
	IsExist bool `json:"is_exist"`
}

// Base contains common columns for all tables.
type Base struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	IsDeleted bool      `gorm:"column:is_deleted;default:false" json:"is_deleted"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {

	if base == nil {
		return nil
	}

	if base.CreatedAt.IsZero() {
		base.CreatedAt = time.Now()
	}

	if base.UpdatedAt.IsZero() {
		base.UpdatedAt = time.Now()
	}

	// 关键修改：检查 Sequence 是否为 nil
	if utility.Sequence != nil {
		base.ID = utility.Sequence.Next()
	}

	return nil

}

// RelationBase contains common columns for all tables.
type RelationBase struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (base *RelationBase) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = utility.Sequence.Next()
	return
}

// ListRequest list request
type ListRequest struct {
	Limit     int  `form:"limit" binding:"gt=0"`
	Offset    int  `form:"offset" binding:"gte=0"`
	IsDeleted bool `form:"isDeleted"`
}

type BusinessModule int

const (
	ModuleStartup BusinessModule = iota + 1
	ModuleBounty
	ModuleCrowdfunding
	ModuleGovernance
	ModuleOtherDapp
)

var DefaultModules = []BusinessModule{ModuleBounty, ModuleCrowdfunding, ModuleGovernance, ModuleOtherDapp}

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var err error
	// 指定时区
	d.Time, err = time.ParseInLocation(`"2006-01-02"`, string(b), time.Local)
	if err != nil {
		return err
	}
	return nil
}
