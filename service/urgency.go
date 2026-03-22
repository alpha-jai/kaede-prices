package service

import (
	"github.com/koyo/kaede-prices/domain"
)

type UrgencyService struct{}

func NewUrgencyService() *UrgencyService {
	return &UrgencyService{}
}

// CalculateUrgencyScore determines the urgency based on price thresholds.
// Logic: Higher urgency if price is below alert threshold significantly.
func (s *UrgencyService) CalculateUrgencyScore(currentPrice float64, watchlist *domain.Watchlist) float64 {
	if watchlist.AlertThreshold <= 0 {
		return 0
	}
	
	// Example algorithm: (Threshold - CurrentPrice) / Threshold
	score := (watchlist.AlertThreshold - currentPrice) / watchlist.AlertThreshold
	
	if score < 0 {
		return 0
	}
	if score > 1 {
		return 1
	}
	return score
}
