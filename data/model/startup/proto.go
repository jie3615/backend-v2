package startup

import (
	"database/sql"
	"metaLand/data/model"
	"metaLand/data/model/tag"
)

type Startup struct {
	model.Base
	ComerID              uint64 `gorm:"comer_id" json:"comerID"`
	Name                 string `gorm:"name" json:"name"`
	Mode                 uint8  `gorm:"mode" json:"mode"`
	Logo                 string `gorm:"logo" json:"logo"`
	Cover                string `gorm:"Cover" json:"cover"`
	Mission              string `gorm:"mission" json:"mission"`
	TokenContractAddress string `gorm:"token_contract_address" json:"tokenContractAddress"`
	Overview             string `gorm:"overview" json:"overview"`
	//ChainID              uint64    `gorm:"chain_id" json:"chainID"`
	TxHash        string    `gorm:"tx_hash" json:"blockChainAddress"`
	OnChain       bool      `gorm:"on_chain" json:"onChain"`
	KYC           string    `gorm:"kyc" json:"kyc"`
	ContractAudit string    `gorm:"contract_audit" json:"contractAudit"`
	HashTags      []tag.Tag `gorm:"many2many:tag_target_rel;foreignKey:ID;joinForeignKey:TargetID;" json:"hashTags"`
	Website       string    `gorm:"website" json:"website"`
	Discord       string    `gorm:"discord" json:"discord"`
	Twitter       string    `gorm:"twitter" json:"twitter"`
	Telegram      string    `gorm:"telegram" json:"telegram"`
	Docs          string    `gorm:"docs" json:"docs"`

	Email    string `gorm:"email" json:"email"`
	Facebook string `gorm:"facebook" json:"facebook"`
	Medium   string `gorm:"medium" json:"medium"`
	Linktree string `gorm:"linktree" json:"linktree"`

	LaunchNetwork int          `gorm:"launch_network" json:"launchNetwork"`
	TokenName     string       `gorm:"token_name" json:"tokenName"`
	TokenSymbol   string       `gorm:"token_symbol" json:"tokenSymbol"`
	TotalSupply   int64        `gorm:"total_supply" json:"totalSupply"`
	PresaleStart  sql.NullTime `gorm:"presale_start" json:"presaleStart"`
	PresaleEnd    sql.NullTime `gorm:"presale_end" json:"presaleEnd"`
	LaunchDate    sql.NullTime `gorm:"launch_date" json:"launchDate"`
	TabSequence   string       `gorm:"tab_sequence" json:"tabSequence"`
}

// TableName Startup table name for gorm
func (Startup) TableName() string {
	return "startup"
}
