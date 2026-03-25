package timestamp

import (
	"time"
)

type Entity struct {
	CreatedAt time.Time  `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time  `db:"updatedAt" json:"updatedAt"`
	DeletedAt *time.Time `db:"deletedAt" json:"deletedAt"`
}
