package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kalleriakronos24/mygoapp2nd/constants"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
)

type AuthResponseData struct {
	message string
}

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Error: "user not authenticated", Data: &AuthResponseData{
			message: "test",
		}})
	}
}
