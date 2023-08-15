package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autorizationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	authHeader := c.GetHeader(autorizationHeader)
	if authHeader == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Autorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
