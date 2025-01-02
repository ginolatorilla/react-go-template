package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	u "github.com/ginolatorilla/go-template/pkg/utils"
	"github.com/ginolatorilla/react-go-template/ui"
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

	defer zap.S().Sync()
	zap.S().Debugf("Server options: %+v", c)
	server.Options = c

	server.setUpGin()
	return server
}

func (s *server) Serve() error {
	defer zap.S().Sync()
	zap.S().Infof("Starting %s version %s and listening at %s", s.name, s.version, s.ListenAddress)
	return http.ListenAndServe(s.ListenAddress, s.router)
}

func (s *server) setUpGin() {
	engine := gin.Default()
	s.setUpMiddleware(engine)
	s.setUpHandlers(engine)
	s.router = engine
}

func (s *server) setUpMiddleware(engine *gin.Engine) {
	if s.EnableCORS {
		defer zap.S().Sync()
		zap.S().Info("CORS is enabled")
		engine.Use(CORSMiddleware())
	}
}

func (s *server) setUpHandlers(engine *gin.Engine) {
	engine.GET("/", s.handleRoot)
	engine.StaticFS("/ui", http.FS(u.Must(fs.Sub(ui.Embedded, "dist"))))
}

func (s *server) handleRoot(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("Hello, World! %s version: %s", s.name, s.version))
}
