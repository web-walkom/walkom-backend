package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllExcursions(c *gin.Context) {
	excursions, err := h.services.Excursions.GetAllExcursions(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Successfully received a list of all excursions")
	c.JSON(http.StatusOK, excursions)
}

func (h *Handler) GetExcursionsById(c *gin.Context) {

}
