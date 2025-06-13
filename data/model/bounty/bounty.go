package bounty

type ListBountiesRequest struct {
	Limit     int    `form:"limit" binding:"gt=0"`
	Offset    int    `form:"offset" binding:"gte=0"`
	IsDeleted bool   `form:"isDeleted"`
	Keyword   string `form:"keyword"`
	Mode      uint8  `form:"mode"`
}

type ListBountiesResponse struct {
	List  []Bounty `json:"list"`
	Total int64    `json:"total"`
}
