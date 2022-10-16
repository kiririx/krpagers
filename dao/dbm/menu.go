package dbm

type Menu struct {
	ParentId uint64 `ami:"parent_id"`
	Text     string `ami:"text"`
	Url      string `ami:"url"`
	BaseModel
}
