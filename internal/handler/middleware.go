package handler

import (
	"net/http"
	"strings"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	_, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", domain.ErrHeaderAuthorizedIsEmpty
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", domain.ErrInvalidHeaderAuthorized
	}

	if len(headerParts[1]) == 0 {
		return "", domain.ErrTokenIsEmpty
	}

	return h.services.Auth.ParseToken(headerParts[1])
}
