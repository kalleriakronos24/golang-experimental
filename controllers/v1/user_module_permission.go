package v1

import (
	"fmt"
	"github.com/kalleriakronos24/golang-experimental/constants"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/golang-experimental/dto"
	masterModels "github.com/kalleriakronos24/golang-experimental/models/master"
	"github.com/kalleriakronos24/golang-experimental/services"
)

func POSTUserModulePermission(c *gin.Context) {
	var err error
	var p dto.CreateUserModulePermission
	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	fmt.Printf("%+v\n", p)

}

func GETUserModulePermission(c *gin.Context) {
	var err error
	var UserInfo dto.RetrieveUserInfo
	if err = c.ShouldBindUri(&UserInfo); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	var user masterModels.User
	if user, err = services.Handler.RetrieveUser(UserInfo.Username); err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Error: "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: dto.RetrieveUserInfo{
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
	}})
}

func PUTUserModulePermission(c *gin.Context) {
	var err error
	var user dto.UserUpdate
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	id, _ := c.Params.Get("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: "Parameter ID is empty"})
		return
	}
	if err = services.Handler.UpdateUser(userId, user); err != nil {
		c.JSON(http.StatusNotModified, dto.Response{Error: "failed updating user"})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: user})
}
