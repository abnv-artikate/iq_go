package handlers

import (
	"net/http"
	"strconv"

	"iq-go/internal/services"
	"iq-go/internal/utils"

	"github.com/gin-gonic/gin"
)

type ResultHandler struct {
	resultService *services.ResultService
}

func NewResultHandler(resultService *services.ResultService) *ResultHandler {
	return &ResultHandler{
		resultService: resultService,
	}
}

func (h *ResultHandler) GetResults(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	results, err := h.resultService.GetResultsByUserID(userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch results")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Results fetched successfully", results)
}

func (h *ResultHandler) GetResult(c *gin.Context) {
	resultIDStr := c.Param("id")
	resultID, err := strconv.ParseUint(resultIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid result ID")
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	result, err := h.resultService.GetResultByID(uint(resultID), userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Result not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Result fetched successfully", result)
}
