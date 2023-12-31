package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/golang-experimental/constants"
	"github.com/kalleriakronos24/golang-experimental/dto"
	"net/http"
)

type AuthResponseData struct {
	message string
}

type queryParamPayload struct {
	key   string
	value []string
}

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Error: "User not authenticated", Data: nil})
	}
}

func ApiRequestLogger(err error, msg string) {}
