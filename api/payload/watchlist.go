package payload

type WatchlistCreateRequest struct {
	ItemID         int64   `json:"item_id"`
	AlertThreshold float64 `json:"alert_threshold"`
}

type WatchlistListResponse struct {
	Watchlists []WatchlistResponse `json:"watchlists"`
}

type WatchlistResponse struct {
	ID             int64   `json:"id"`
	ItemID         int64   `json:"item_id"`
	AlertThreshold float64 `json:"alert_threshold"`
}
