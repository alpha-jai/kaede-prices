package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koyo/kaede-prices/service"
)

type AffiliateHandler struct {
	svc *service.AffiliateService
}

func NewAffiliateHandler(svc *service.AffiliateService) *AffiliateHandler {
	return &AffiliateHandler{svc: svc}
}

func (h *AffiliateHandler) GetLinks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"link": h.svc.GenerateLink(1)})
}

func (h *AffiliateHandler) TrackClick(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "tracked"})
}
