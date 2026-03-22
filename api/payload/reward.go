package payload

type RewardsListResponse struct {
	UserID      uint `json:"user_id"`
	TotalPoints int  `json:"total_points"`
}
