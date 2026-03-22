package service

import (
	"github.com/koyo/kaede-prices/domain"
)

type GroceryFetcher interface {
	FetchPrices(itemID int64) ([]domain.PriceEntry, error)
}

type MockGroceryFetcher struct{}

func NewMockGroceryFetcher() *MockGroceryFetcher {
	return &MockGroceryFetcher{}
}

func (m *MockGroceryFetcher) FetchPrices(itemID int64) ([]domain.PriceEntry, error) {
	// Mock implementation
	return []domain.PriceEntry{
		{
			ItemID:    itemID,
			Price:     10.99,
			Source:    "API",
			Timestamp: 1742616000, // Current rough timestamp
		},
	}, nil
}
