package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type serverConfig struct {
	listenAddress string
	enableCORS    bool
}

type server struct {
	serverConfig

	name    string
	version string
	router  http.Handler
}

func NewServer(c serverConfig, version, name string) *server {
	server := new(server)
	server.version = version
	server.name = name
	server.listenAddress = c.listenAddress
	server.enableCORS = c.enableCORS

	engine := gin.Default()
	server.router = engine

	if server.enableCORS {
		defer zap.S().Sync()
		zap.S().Info("CORS is enabled")
		engine.Use(CORSMiddleware())
	}

	engine.GET("/", server.handleRoot)

	return server
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *server) Serve() {
	defer zap.S().Sync()
	zap.S().Infof("%s version: %s", s.name, s.version)
	http.ListenAndServe(s.listenAddress, s.router)
}

func (s *server) handleRoot(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("Hello, World! %s version: %s", s.name, s.version))
}
