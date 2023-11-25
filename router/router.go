package router

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/kalleriakronos24/golang-experimental/controllers/middlewares"
	v1 "github.com/kalleriakronos24/golang-experimental/controllers/v1"
	v1Master "github.com/kalleriakronos24/golang-experimental/controllers/v1/master"
	"github.com/kalleriakronos24/golang-experimental/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin).Use()
			auth.POST("/logout", v1.POSTRegister)
		}
		user := v1route.Group("/user")
		{
			user.GET("/:username", utils.AuthOnly, v1Master.GETUser)
			user.PUT("/:id", utils.AuthOnly, v1Master.PUTUser)
		}
		misc := v1route.Group("/misc")
		{
			misc.GET("/ping", v1.Pong)
			misc.POST("/upload", v1.UploadFileSingle)
			misc.POST("/upload-multiple", v1.UploadFileMultiple)
		}

		masterModule := v1route.Group("/module")
		{
			masterModule.GET("/", v1Master.GETOneMasterModule)
			masterModule.POST("/", v1Master.POSTMasterModule)
			masterModule.PUT("/:id", v1Master.PUTMasterModule)
			masterModule.DELETE("/:id", v1Master.DELETEMasterModule)
		}

		userModulePermission := v1route.Group("/user-module-permission")
		{
			userModulePermission.POST("/", v1.POSTUserModulePermission)
		}
	}
	return
}

func InitializeEnterpriseRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin).Use()
			auth.POST("/signup", v1.POSTRegister)
		}
		user := v1route.Group("/user")
		{
			user.GET("/:username", utils.AuthOnly, v1Master.GETUser)
			user.PUT("/:id", utils.AuthOnly, v1Master.PUTUser)
		}
		misc := v1route.Group("/misc")
		{
			misc.GET("/ping", v1.Pong)
			misc.POST("/upload", v1.UploadFileSingle)
			misc.POST("/upload-multiple", v1.UploadFileMultiple)
		}
	}
	return
}
