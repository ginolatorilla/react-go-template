package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *server) setUpAPIHandler(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	api.GET("/hello", s.handleRoot)
}

func (s *server) handleRoot(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("Hello, World! %s version: %s", s.name, s.version))
}
