package handler

import (
	"net/http"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCodeEmail(c *gin.Context) {
	var inp domain.AuthEmail
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidInput.Error())
		return
	}

	err := h.services.Auth.SendCodeEmail(c, inp)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) CheckCodeEmail(c *gin.Context) {
	var inp domain.AuthCode
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidInput.Error())
		return
	}

	err := h.services.Auth.CheckSecretCode(c, inp)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Users.CreateUserIfNotExist(c, inp.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Users.GetUserByEmail(c, inp.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Auth.CreateSession(user.ID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Auth user %s", res.ID)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          res.ID,
		"email":       inp.Email,
		"accessToken": res.AccessToken,
	})
}
