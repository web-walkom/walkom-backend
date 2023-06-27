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
	c.JSON(http.StatusOK, domain.ResultSendCode{Status: true})
}

func (h *Handler) CheckCodeEmail(c *gin.Context) {
	var inp domain.AuthCode
	if err := c.BindJSON(&inp); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err, domain.ErrInvalidInput)
		return
	}

	err := h.services.Auth.CheckSecretCode(c, inp)
	if err != nil {
		if errors.Is(err, domain.ErrSecretCodeInvalid) || errors.Is(err, domain.ErrSecretCodeExpired) {
			c.JSON(http.StatusOK, domain.ResultCheckCode{
				Status: false,
				Error:  err.Error(),
			})
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

	c.JSON(http.StatusOK, domain.ResultCheckCode{
		Status:      true,
		ID:          res.ID,
		Email:       inp.Email,
		AccessToken: res.AccessToken,
	})
}
