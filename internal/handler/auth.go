package handler

import (
	"errors"
	"net/http"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCodeEmail(c *gin.Context) {
	var inp domain.AuthEmail
	if err := c.BindJSON(&inp); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrInvalidInput)
		return
	}

	err := h.services.Auth.SendCodeEmail(c, inp)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrSendCodeEmail)
		return
	}

	h.log.Infof("Success send code to email: %s", inp.Email)
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) CheckCodeEmail(c *gin.Context) {
	var inp domain.AuthCode
	if err := c.BindJSON(&inp); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrInvalidInput)
		return
	}

	err := h.services.Auth.CheckSecretCode(c, inp)
	if err != nil {
		if errors.Is(err, domain.ErrSecretCodeInvalid) {
			c.JSON(http.StatusOK, domain.ErrSecretCodeInvalid.Error())
			return
		}
		if errors.Is(err, domain.ErrSecretCodeExpired) {
			c.JSON(http.StatusOK, domain.ErrSecretCodeExpired.Error())
			return
		}
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrCheckCodeEmail)
		return
	}

	err = h.services.Users.CreateUserIfNotExist(c, inp.Email)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrCreateUser)
		return
	}

	user, err := h.services.Users.GetUserByEmail(c, inp.Email)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrGetUser)
		return
	}

	res, err := h.services.Auth.CreateSession(user.ID)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrCreateSession)
		return
	}

	h.log.Infof("Success auth user: %s", inp.Email)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          res.ID,
		"email":       inp.Email,
		"accessToken": res.AccessToken,
	})
}
