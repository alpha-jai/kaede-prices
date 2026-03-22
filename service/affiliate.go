package service

import (
	"fmt"
)

type AffiliateService struct{}

func NewAffiliateService() *AffiliateService {
	return &AffiliateService{}
}

func (s *AffiliateService) GenerateLink(userID uint) string {
	return fmt.Sprintf("https://kaede-prices.com/ref/%d", userID)
}
