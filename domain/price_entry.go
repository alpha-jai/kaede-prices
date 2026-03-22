package domain

type PriceEntry struct {
	ID              int64   `json:"id" db:"id" dbfield:"id"`
	UUID            string  `json:"uuid" db:"uuid" dbfield:"uuid"`
	ItemID          int64   `json:"item_id" db:"item_id" dbfield:"item_id"`
	StoreID         int64   `json:"store_id" db:"store_id" dbfield:"store_id"`
	Price           float64 `json:"price" db:"price" dbfield:"price"`
	ReporterID      int64   `json:"reporter_id" db:"reporter_id" dbfield:"reporter_id"`
	ReliabilityFlag int     `json:"reliability_flag" db:"reliability_flag" dbfield:"reliability_flag"`
	Timestamp       int64   `json:"timestamp" db:"timestamp" dbfield:"timestamp"`
	CreatedAt       int64   `json:"created_at" db:"created_at" dbfield:"created_at"`
	UpdatedAt       int64   `json:"updated_at" db:"updated_at" dbfield:"updated_at"`
	Source          string  `json:"source" db:"source" dbfield:"source"`
	DeletedAt       *int64  `json:"deleted_at" db:"deleted_at" dbfield:"deleted_at"`
}
