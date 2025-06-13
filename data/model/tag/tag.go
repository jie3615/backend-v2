package tag

type TagListRequest struct {
	Ad   bool   `form:"ad"`
	Type string `form:"type"`
}
