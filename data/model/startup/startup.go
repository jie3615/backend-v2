package startup

import "time"

type ListStartupRequest struct {
	Limit     int    `form:"limit" binding:"gt=0"`
	Offset    int    `form:"offset" binding:"gte=0"`
	IsDeleted bool   `form:"isDeleted"`
	Keyword   string `form:"keyword"`
	Mode      uint8  `form:"mode"`
}

type ListStartupsResponse struct {
	List  []Startup `json:"list"`
	Total int64     `json:"total"`
}

type CreateStartupsRequest struct {
	ComerID              uint64     `json:"comer_id"`
	Name                 string     `json:"name"`
	Mode                 uint8      `json:"mode"`
	Logo                 string     `json:"logo"`
	Cover                string     `json:"cover"`
	Mission              string     `json:"mission"`
	TokenContractAddress string     `json:"token_contract_address"`
	Overview             string     `json:"overview"`
	TxHash               string     `json:"tx_hash"`
	OnChain              bool       `json:"on_chain"`
	KYC                  string     `json:"kyc"`
	ContractAudit        string     `json:"contract_audit"`
	Website              string     `json:"website"`
	Discord              string     `json:"discord"`
	Twitter              string     `json:"twitter"`
	Telegram             string     `json:"telegram"`
	Docs                 string     `json:"docs"`
	Email                string     `json:"email"`
	Facebook             string     `json:"facebook"`
	Medium               string     `json:"medium"`
	Linktree             string     `json:"linktree"`
	LaunchNetwork        int        `json:"launch_network"`
	TokenName            string     `json:"token_name"`
	TokenSymbol          string     `json:"token_symbol"`
	TotalSupply          int64      `json:"total_supply"`
	PresaleStart         *time.Time `json:"presale_start"`
	PresaleEnd           *time.Time `json:"presale_end"`
	LaunchDate           *time.Time `json:"launch_date"`
	TabSequence          string     `json:"tab_sequence"`
	IsDeleted            bool       `json:"is_deleted"`
}

type CheckStartupsRequest struct {
	IsDeleted            bool   `form:"isDeleted"`
	Name                 string `form:"name"`
	TokenContractAddress string `form:"tokenContractAddress"`
}
