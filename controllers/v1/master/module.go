package v1

import (
	"github.com/kalleriakronos24/golang-experimental/constants"
	"github.com/kalleriakronos24/golang-experimental/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/kalleriakronos24/golang-experimental/dto"
	"github.com/kalleriakronos24/golang-experimental/services"
)

func GETOneMasterModule(c *gin.Context) {
	var p dto.RetrieveOneMasterModule
	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("param-query-error", err, ""))
		return
	}

	if _, err := services.Handler.RetrieveMasterModule(p); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("retrieve-failed", err, "module"))
		return
	}

	data, _ := services.Handler.RetrieveMasterModule(p)

	c.JSON(http.StatusOK, dto.Response{Data: dto.RetrieveOneMasterModule{
		ModuleName:  data.ModuleName,
		Description: data.Description,
	}})
}

func POSTMasterModule(c *gin.Context) {
	var err error
	var p dto.CreateMasterModule
	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingMasterModule(p.ModuleName, ""); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, ""))
		return
	}

	if _, err := services.Handler.CreateMasterModule(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "module"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTMasterModule(c *gin.Context) {
	var err error
	var p dto.UpdateMasterModule
	if err = c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	id, _ := c.Params.Get("id")
	moduleId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingMasterModule("", id); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, ""))
		return
	}

	if err = services.Handler.UpdateMasterModule(moduleId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "module"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: p})
}

func DELETEMasterModule(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	moduleId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingMasterModule("", id); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, ""))
		return
	}

	if err = services.Handler.DeleteMasterModule(moduleId); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("delete-failed", err, "module"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
