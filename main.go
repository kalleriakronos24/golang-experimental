package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/mygoapp2nd/config"
	"github.com/kalleriakronos24/mygoapp2nd/router"
	"github.com/kalleriakronos24/mygoapp2nd/services"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	if err := services.InitializeServices(); err != nil {
		log.Fatalln(err)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
