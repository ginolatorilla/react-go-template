package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serverConfig struct {
	listenAddress string
}

type server struct {
	serverConfig

	version string
	router  http.Handler
}

func NewServer(c serverConfig) *server {
	server := new(server)
	server.version = version
	server.listenAddress = c.listenAddress

	engine := gin.Default()
	server.router = engine

	engine.GET("/", server.handleRoot)

	return server
}

func (s *server) Serve() {
	logger().Infof("%s version: %s", app, s.version)
	http.ListenAndServe(s.listenAddress, s.router)
}

func (s *server) handleRoot(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("Hello, World! %s version: %s", app, s.version))
}
