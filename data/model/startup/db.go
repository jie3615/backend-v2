package startup

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"metaLand/data/model"
	"time"
)

// ListStartups  list startups
func ListStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false")
	if comerID != 0 {
		db = db.Where("comer_id = ?", comerID)
	}
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).
		//Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").
		//Preload("Members.ComerProfile").Preload("Follows").
		Find(startups).Error
	return
}

func StartupOnChain(db *gorm.DB, txHash string, chainID uint64, comerID uint64) (err error) {
	return db.Where(Startup{
		TxHash: txHash,
		//ChainID: chainID,
		ComerID: comerID,
	}).Where("on_chain", false).Updates(Startup{
		OnChain: true,
	}).Error
}

// 创建项目 方法

func CreateStartups(db *gorm.DB, input *CreateStartupsRequest) (suc bool, err error) {
	// 创建Startup实例
	startup := &Startup{
		Base:                 model.Base{}, // 关键：确保 Base 非 nil
		ComerID:              input.ComerID,
		Name:                 input.Name,
		Mode:                 input.Mode,
		Logo:                 input.Logo,
		Cover:                input.Cover,
		Mission:              input.Mission,
		TokenContractAddress: input.TokenContractAddress,
		Overview:             input.Overview,
		TxHash:               input.TxHash,
		OnChain:              input.OnChain,
		KYC:                  input.KYC,
		ContractAudit:        input.ContractAudit,
		Website:              input.Website,
		Discord:              input.Discord,
		Twitter:              input.Twitter,
		Telegram:             input.Telegram,
		Docs:                 input.Docs,
		Email:                input.Email,
		Facebook:             input.Facebook,
		Medium:               input.Medium,
		Linktree:             input.Linktree,
		LaunchNetwork:        input.LaunchNetwork,
		TokenName:            input.TokenName,
		TokenSymbol:          input.TokenSymbol,
		TotalSupply:          input.TotalSupply,
		PresaleStart:         TimePtrToNullTime(input.PresaleStart),
		PresaleEnd:           TimePtrToNullTime(input.PresaleEnd),
		LaunchDate:           TimePtrToNullTime(input.LaunchDate),
		TabSequence:          input.TabSequence,
	}

	// 执行创建操作
	result := db.Create(startup)

	// 检查是否有错误发生
	if result.Error != nil {
		return false, result.Error
	}

	// 检查是否成功插入记录
	if result.RowsAffected == 0 {
		return false, errors.New("failed to insert startup record")
	}

	return true, nil
}

// 辅助函数：将 string 转换为 sql.NullTime

func TimePtrToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}

// 实现 CheckExists 方法

func CheckExists(db *gorm.DB, input *CheckStartupsRequest) (bool, error) {
	var whereClause string
	var args []interface{}

	// 添加删除状态条件
	if !input.IsDeleted {
		whereClause += "is_deleted = 0 AND "
	}

	// 添加唯一字段条件
	if input.Name != "" {
		whereClause += "name = ? AND "
		args = append(args, input.Name)
	}
	if input.TokenContractAddress != "" {
		whereClause += "token_contract_address = ? AND "
		args = append(args, input.TokenContractAddress)
	}

	// 去除末尾的 "AND "
	if whereClause != "" {
		whereClause = whereClause[:len(whereClause)-5]
	} else {
		return false, nil // 无查询条件，直接返回不存在
	}

	// 执行查询
	var count int64
	query := db.Model(&Startup{}).Where(whereClause, args...)

	// 检查是否存在记录
	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetStartupInfo

func GetStartupInfo(db *gorm.DB, startupId *uint64) (resp *Startup, err error) {

	// 检查 startupId 是否为空指针
	if startupId == nil {
		return nil, fmt.Errorf("startupId is required")
	}

	// 初始化响应对象
	resp = &Startup{}

	// 查询数据库
	result := db.Where("id = ? AND is_deleted = ?", *startupId, false).First(resp)

	// 处理查询结果
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("startup with ID %d not found", *startupId)
		}
		return nil, fmt.Errorf("database error: %w", result.Error)
	}

	return resp, nil

}

// UpdateStartup

func UpdateStartup(db *gorm.DB, data *Startup) (bool, error) {

	// 执行更新操作
	result := db.Updates(data)

	// 检查是否有错误发生
	if result.Error != nil {
		return false, result.Error
	}

	// 检查是否成功更新记录
	if result.RowsAffected == 0 {
		return false, errors.New("failed to insert startup record")
	}

	return true, nil

}
