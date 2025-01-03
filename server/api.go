package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// setUpAPIHandler installs the API routes to the provided Gin engine.
func (s *server) setUpAPIHandler(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	api.GET("/hello", s.handleHello)
}

// handleHello is an example handler that returns a simple message.
func (s *server) handleHello(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("Hello, World! %s version: %s", s.name, s.version))
}
