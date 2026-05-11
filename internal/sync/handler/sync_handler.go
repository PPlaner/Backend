package handler

import (
	"net/http"

	"github.com/PPlaner/Backend/internal/dto"
	"github.com/PPlaner/Backend/internal/response"
	"github.com/PPlaner/Backend/internal/sync/service"
	"github.com/gin-gonic/gin"
)

type SyncHandler struct {
	syncService *service.SyncService
}

func NewSyncHandler(syncService *service.SyncService) *SyncHandler {
	return &SyncHandler{
		syncService: syncService,
	}
}

func (h *SyncHandler) Sync(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "user is not authorized")
		return
	}

	userID, ok := userIDValue.(int)
	if !ok {
		response.Error(c, http.StatusUnauthorized, "invalid user id")
		return
	}

	var req dto.SyncRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	result, err := h.syncService.Sync(c.Request.Context(), userID, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to sync data")
		return
	}

	response.Success(c, http.StatusOK, result)

}
