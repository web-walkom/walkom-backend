package handler

import (
	"net/http"

	"github.com/b0shka/walkom-backend/internal/domain"

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

	excursion, err := h.services.Excursions.GetExcursionById(c, excursionId)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrGetExcusion)
		return
	}

	h.log.Infof("Success get data excusrion by id: %s", c.Param("id"))
	c.JSON(http.StatusOK, excursion)
}
