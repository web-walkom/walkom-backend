package handler

import (
	"net/http"

	"github.com/b0shka/walkom-backend/internal/domain"

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

func (h *Handler) UpdateUser(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrServer)
		return
	}

	var inp domain.UpdateUser
	if err := c.BindJSON(&inp); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrInvalidInput)
		return
	}
	inp.ID = userId

	if err = h.services.Users.UpdateUser(c, inp); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrGetUserById)
		return
	}

	h.log.Infof("Success update data user: %s", userId)
	c.JSON(http.StatusOK, domain.ResultUpdateUser{Status: true})
}
