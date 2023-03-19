package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetAllExcursions(c *gin.Context) {
	excursions, err := h.services.Excursions.GetAllExcursions(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, excursions)
}

func (h *Handler) GetExcursionsById(c *gin.Context) {
	excursionId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"excursionId": excursionId,
	})
}
