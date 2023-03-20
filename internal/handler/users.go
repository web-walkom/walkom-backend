package handler

import (
	"github.com/b0shka/walkom-backend/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetUserById(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrServer)
		return
	}

	user, err := h.services.Users.GetUserById(c, userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrGetUserById)
		return
	}

	h.log.Infof("Success get data user: %s", c.Param("id"))
	c.JSON(http.StatusOK, user)
}
