package handler

import (
	"github.com/b0shka/walkom-backend/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetAllExcursions(c *gin.Context) {
	excursions, err := h.services.Excursions.GetAllExcursions(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrGetAllExcusions)
		return
	}
	
	h.log.Info("Success get all excusrions")
	c.JSON(http.StatusOK, excursions)
}

func (h *Handler) GetExcursionsById(c *gin.Context) {
	excursionId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrServer)
		return
	}

	h.log.Infof("Success get data excusrion: %s", c.Param("id"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"excursionId": excursionId,
	})
}
