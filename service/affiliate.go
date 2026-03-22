package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type AffiliateService struct {
}

func NewAffiliateService() *AffiliateService {
	return &AffiliateService{}
}

func (s *AffiliateService) GenerateReferralLink(userID string) string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	code := hex.EncodeToString(bytes)
	return fmt.Sprintf("https://kaede.prices/ref/%s/%s", userID, code)
}
