package domain

type AffiliateInfo struct {
	PartnerID    uint   `json:"partner_id"`
	DiscountCode string `json:"discount_code"`
	ReferralLink string `json:"referral_link"`
}
