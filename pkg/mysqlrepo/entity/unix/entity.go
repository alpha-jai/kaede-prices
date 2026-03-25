package unix

type Entity struct {
	CreatedAt int64 `db:"createdAt" json:"createdAt"`
	UpdatedAt int64 `db:"updatedAt" json:"updatedAt"`
	DeletedAt int64 `db:"deletedAt" json:"deletedAt"`
}
