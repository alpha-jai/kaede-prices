package domain

type Item struct {
	ID            int64   `json:"id" db:"id" dbfield:"id"`
	UUID          string  `json:"uuid" db:"uuid" dbfield:"uuid"`
	Name          string  `json:"name" db:"name" dbfield:"name"`
	Category      *string `json:"category" db:"category" dbfield:"category"`
	Brand         *string `json:"brand" db:"brand" dbfield:"brand"`
	StandardUnit  *string `json:"standard_unit" db:"standard_unit" dbfield:"standard_unit"`
	CreatedAt     int64   `json:"created_at" db:"created_at" dbfield:"created_at"`
	UpdatedAt     int64   `json:"updated_at" db:"updated_at" dbfield:"updated_at"`
	DeletedAt     *int64  `json:"deleted_at" db:"deleted_at" dbfield:"deleted_at"`
}
