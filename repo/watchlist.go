package repo

import (
	"context"
	"github.com/koyo/kaede-prices/domain"
)

type WatchlistRepository interface {
	Create(ctx context.Context, watchlist *domain.Watchlist) error
	ListByUserID(ctx context.Context, userID int64) ([]*domain.Watchlist, error)
	Delete(ctx context.Context, id int64) error
}
