package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Options struct {
	ListenAddress string
	EnableCORS    bool
}

type server struct {
	Options

	name    string
	version VersionInfo
	router  http.Handler
}

func NewServer(c Options) *server {
	server := new(server)
	server.version.CommitHash = CommitHash
	server.version.Version = Version
	server.name = AppName

	zap.S().Debugf("Server options: %+v", c)
	server.Options = c

	engine := gin.Default()
	if server.EnableCORS {
		defer zap.S().Sync()
		zap.S().Info("CORS is enabled")
		engine.Use(CORSMiddleware())
	}
	engine.GET("/", server.handleRoot)

	server.router = engine
	return server
}

func (s *server) Serve() error {
	defer zap.S().Sync()
	zap.S().Infof("%s version: %s", s.name, s.version)
	return http.ListenAndServe(s.ListenAddress, s.router)
}

func (s *server) handleRoot(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("Hello, World! %s version: %s", s.name, s.version))
}
