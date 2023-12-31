package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kalleriakronos24/golang-experimental/config"
	"github.com/kalleriakronos24/golang-experimental/migrations"
	"github.com/kalleriakronos24/golang-experimental/pkg/sockets"
	"github.com/kalleriakronos24/golang-experimental/router"
	"github.com/kalleriakronos24/golang-experimental/services"
	"github.com/spf13/viper"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func runServer() {
	// initialize db and migrations
	if err := services.InitializeServices(); err != nil {
		log.Fatalln(err)
	}
	migrations.Migrate()

	// serve all routes and routes configuration
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

func main() {
	if viper.GetBool("SOCKET_ENABLED") == true {
		server := socketio.NewServer(nil)
		var serv = sockets.RunSocketConnection(server)
		http.Handle("/socket.io/", serv)
	}
	runServer()
}
