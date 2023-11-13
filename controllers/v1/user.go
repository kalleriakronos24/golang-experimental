package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
	master "github.com/kalleriakronos24/mygoapp2nd/models/master"
	"github.com/kalleriakronos24/mygoapp2nd/services"
)

func GETUser(c *gin.Context) {
	var err error
	var userInfo dto.UserInfoAll
	if err = c.ShouldBindUri(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	var user master.User
	if user, err = services.Handler.RetrieveUser(userInfo.Username); err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Error: "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: dto.UserInfo{
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
	}})
}

func PUTUser(c *gin.Context) {
	var err error
	var user dto.UserUpdate
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	id, _ := c.Params.Get("id")
	uuid, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: "Parameter ID is empty"})
		return
	}
	if err = services.Handler.UpdateUser(uuid, user); err != nil {
		c.JSON(http.StatusNotModified, dto.Response{Error: "failed updating user"})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: user})
}
