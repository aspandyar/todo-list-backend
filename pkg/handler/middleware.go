package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaded = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeaded)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("empty auth header").Error())
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("invalid auth header").Error())
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("user id is not found").Error())
		return 0, errors.New("user id is not found")
	}

	idToInt, ok := userId.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("user id is of invalid type").Error())
		return 0, errors.New("user id is of invalid type")
	}

	return idToInt, nil
}
