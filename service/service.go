package service

import "github.com/koyo/kaede-prices/repo"

// RegisterFetchers sets up the fetcher dependencies for the application.
// In a real application, this might involve dependency injection or a service registry.
func RegisterFetchers() GroceryFetcher {
	return NewMockGroceryFetcher()
}

func GetRewardService() *RewardService {
	return NewRewardService(repo.NewRewardRepo())
}

func GetAffiliateService() *AffiliateService {
	return NewAffiliateService()
}
