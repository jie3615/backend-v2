package bounty

import (
	"math/big"
	"metaLand/data/model"
	"metaLand/data/model/startup"
	"time"
)

type Bounty struct {
	model.Base
	ChainId             uint64               `gorm:"chain_id" json:"chain_id"`
	TxHash              string               `gorm:"tx_hash" json:"tx_hash"`
	DepositContract     string               `gorm:"deposit_contract" json:"deposit_contract"`
	StartupId           uint64               `gorm:"startup_id" json:"startup_id"`
	ComerId             uint64               `gorm:"comer_id" json:"comer_id"`
	Title               string               `gorm:"title" json:"title"`
	ApplyCutoffDate     time.Time            `gorm:"apply_cutoff_date" json:"apply_cutoff_date"`
	DiscussionLink      string               `gorm:"discussion_link" json:"discussion_link"`
	DepositTokenSymbol  string               `gorm:"deposit_token_symbol" json:"deposit_token_symbol"`
	ApplicantDeposit    big.Int              `gorm:"applicant_deposit" json:"applicant_deposit"`
	FounderDeposit      big.Int              `gorm:"founder_deposit" json:"founder_deposit"`
	Description         string               `gorm:"description" json:"description"`
	PaymentMode         int                  `gorm:"payment_mode" json:"payment_mode"`
	Status              int                  `gorm:"status" json:"status"`
	TotalRewardToken    int                  `gorm:"total_reward_token" json:"total_reward_token"`
	BountyPaymentPeriod BountyPaymentPeriod  `gorm:"foreignKey:BountyId" json:"reward"`
	Startup             startup.Startup      `gorm:"foreignKey:StartupId" json:"startup"`
	BountyApplicants    []BountyApplicant    `gorm:"foreignKey:BountyId" json:"applicants"`
	BountyContacts      []BountyContact      `gorm:"foreignKey:BountyId" json:"contacts"`
	BountyDeposits      []BountyDeposit      `gorm:"foreignKey:BountyId" json:"deposit_records"`
	BountyPaymentTerms  []BountyPaymentTerms `gorm:"foreignKey:BountyId" json:"terms"`
}

func (Bounty) TableName() string {
	return "bounty"
}

type BountyPaymentPeriod struct {
	model.Base
	BountyId     uint64 `gorm:"bounty_id" json:"bounty_id"`
	PeriodType   int    `gorm:"period_type" json:"period_type"`
	PeriodAmount uint64 `gorm:"period_amount" json:"period_amount"`
	HoursPerDay  int    `gorm:"hours_per_day" json:"hours_per_day"`
	Token1Symbol string `gorm:"token1_symbol" json:"token1_symbol"`
	Token1Amount int    `gorm:"token1_amount" json:"token1_amount"`
	Token2Symbol string `gorm:"token2_symbol" json:"token2_symbol"`
	Token2Amount int    `gorm:"token2_amount" json:"token2_amount"`
	Target       string `gorm:"target" json:"target"`
}

func (BountyPaymentPeriod) TableName() string {
	return "bounty_payment_period"
}

type BountyApplicant struct {
	model.Base
	BountyId    uint64    `gorm:"column:bounty_id;index:idx_bounty" json:"bounty_id"` // 赏金任务ID
	ComerId     uint64    `gorm:"column:comer_id;index:idx_comer" json:"comer_id"`    // 申请人ID
	ApplyAt     time.Time `gorm:"column:apply_at" json:"apply_at"`                    // 申请时间
	RevokeAt    time.Time `gorm:"column:revoke_at" json:"revoke_at"`                  // 撤销时间
	ApproveAt   time.Time `gorm:"column:approve_at" json:"approve_at"`                // 批准时间
	QuitAt      time.Time `gorm:"column:quit_at" json:"quit_at"`                      // 退出时间
	SubmitAt    time.Time `gorm:"column:submit_at" json:"submit_at"`                  // 提交时间
	Status      int       `gorm:"column:status;index:idx_status" json:"status"`       // 申请状态
	Description string    `gorm:"column:description" json:"description"`              // 申请描述
}

func (BountyApplicant) TableName() string {
	return "bounty_applicant"
}

type BountyContact struct {
	model.Base
	BountyId       uint64 `gorm:"column:bounty_id;uniqueIndex:bounty_contact_uindex" json:"bounty_id"`             // 赏金任务ID
	ContactType    uint8  `gorm:"column:contact_type;uniqueIndex:bounty_contact_uindex" json:"contact_type"`       // 联系方式类型
	ContactAddress string `gorm:"column:contact_address;uniqueIndex:bounty_contact_uindex" json:"contact_address"` // 联系地址
}

func (BountyContact) TableName() string {
	return "bounty_contact"
}

type BountyDeposit struct {
	model.Base
	ChainId     uint64    `gorm:"column:chain_id;uniqueIndex:chain_tx_uindex" json:"chain_id"` // 链ID，复合唯一索引
	TxHash      string    `gorm:"column:tx_hash;uniqueIndex:chain_tx_uindex" json:"tx_hash"`   // 交易哈希，复合唯一索引
	Status      int8      `gorm:"column:status" json:"status"`                                 // 质押状态
	BountyId    uint64    `gorm:"column:bounty_id;index:idx_bounty" json:"bounty_id"`          // 关联的赏金任务ID
	ComerId     uint64    `gorm:"column:comer_id;index:idx_comer" json:"comer_id"`             // 用户ID
	Access      int       `gorm:"column:access" json:"access"`                                 // 访问权限
	TokenSymbol string    `gorm:"column:token_symbol" json:"token_symbol"`                     // 代币符号
	TokenAmount int       `gorm:"column:token_amount" json:"token_amount"`                     // 代币数量
	Timestamp   time.Time `gorm:"column:timestamp" json:"timestamp"`                           // 时间戳(指针类型允许NULL)
}

func (BountyDeposit) TableName() string {
	return "bounty_deposit"
}

type BountyPaymentTerms struct {
	model.Base
	BountyId     uint64 `gorm:"column:bounty_id;index:idx_bounty" json:"bounty_id"` // 关联的赏金任务ID
	PaymentMode  int8   `gorm:"column:payment_mode" json:"payment_mode"`            // 支付方式
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1_symbol"`          // 第一种代币符号
	Token1Amount int    `gorm:"column:token1_amount" json:"token1_amount"`          // 第一种代币数量
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2_symbol"`          // 第二种代币符号
	Token2Amount int    `gorm:"column:token2_amount" json:"token2_amount"`          // 第二种代币数量
	Terms        string `gorm:"column:terms" json:"terms"`                          // 支付条款详情
	SeqNum       int    `gorm:"column:seq_num" json:"seq_num"`                      // 排序序号
	Status       int    `gorm:"column:status" json:"status"`                        // 状态
}

// TableName the BountyPaymentTerms table name for gorm
func (BountyPaymentTerms) TableName() string {
	return "bounty_payment_terms"
}
