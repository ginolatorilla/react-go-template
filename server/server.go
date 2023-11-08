package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	version       string
	router        http.Handler
	listenAddress string
}

func NewServer(listenAddr string) *server {
	server := new(server)
	server.version = version
	server.listenAddress = listenAddr

	engine := gin.Default()
	server.router = engine

	engine.GET("/", server.handleRoot)

	return server
}

func (s *server) Serve() {
	http.ListenAndServe(s.listenAddress, s.router)
}

func (s *server) handleRoot(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("Hello, World! %s version: %s", app, s.version))
}
