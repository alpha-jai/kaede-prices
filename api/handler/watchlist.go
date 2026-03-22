package handler

import (
	"github.com/koyo/kaede-prices/api/payload"
	"net/http"
	"github.com/gin-gonic/gin"
)

type WatchlistHandler struct {
	// Add repository here later
}

func NewWatchlistHandler() *WatchlistHandler {
	return &WatchlistHandler{}
}

func (h *WatchlistHandler) Create(c *gin.Context) {
	var req payload.WatchlistCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to create...
	c.JSON(http.StatusCreated, gin.H{"message": "Watchlist created"})
}

func (h *WatchlistHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, payload.WatchlistListResponse{Watchlists: []payload.WatchlistResponse{}})
}

func (h *WatchlistHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
