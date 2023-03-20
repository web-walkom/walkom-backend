package handler

import "github.com/gin-gonic/gin"

type errorMsg struct {
	Message string `json:"message"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, err, message error) {
	h.log.Error(err.Error())
	c.AbortWithStatusJSON(statusCode, errorMsg{message.Error()})
}
