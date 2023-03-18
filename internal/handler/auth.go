package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendEmailCode(c *gin.Context) {
	var inp domain.AuthEmail
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	sender := NewGmailSender(os.Getenv("EMAIL_SENDER_NAME"), os.Getenv("EMAIL_SENDER_ADDRESS"), os.Getenv("EMAIL_SENDER_PASSWORD"))

	subject := "Код подтверждения для входа в аккаунт"
	to := inp.Email
	secret_code := utils.RandomInt(100000, 999999)
	content := fmt.Sprintf(`
	<p>Вы отправили запрос на вход в аккаунт под адресом %s.</p>
	<p>Код подтверждения: %d</p>
	`, to, secret_code)

	err := sender.SendEmail(subject, content, to)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Auth.CreateVerifyEmail(c, to, secret_code)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) CheckEmailCode(c *gin.Context) {
	var inp domain.AuthCode
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
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

	res, err := h.services.Auth.CreateSession(c, user.ID)
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
