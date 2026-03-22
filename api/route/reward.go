package route

import (
	"kaede-prices/api/handler"
	"kaede-prices/repo"
	"net/http"
)

func RegisterRewardRoutes(mux *http.ServeMux) {
	r := repo.NewRewardRepository()
	h := handler.NewRewardHandler(r)
	
	mux.HandleFunc("GET /api/v1/rewards", h.GetRewards)
	mux.HandleFunc("POST /api/v1/rewards/claim", h.ClaimReward)
}
