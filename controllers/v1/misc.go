package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{Data: "pong"})
}
