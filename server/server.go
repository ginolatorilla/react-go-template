// Package server defines the HTTP web server.
//
// # Copyright © 2024 Gino Latorilla
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package server

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	u "github.com/ginolatorilla/go-template/pkg/utils"
	"github.com/ginolatorilla/react-go-template/ui"
	"go.uber.org/zap"
)

// Options defines the server's configuration.
type Options struct {
	ListenAddress string
	EnableCORS    bool
}

// server is the HTTP web server.
type server struct {
	Options

	name    string
	version VersionInfo
	router  http.Handler
}

// NewServer creates a new web server with the given configuration options.
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

// Serve starts the web server, listening to the configured address.
func (s *server) Serve() error {
	defer zap.S().Sync()
	zap.S().Infof("Starting %s version %s and listening at %s", s.name, s.version, s.ListenAddress)
	return http.ListenAndServe(s.ListenAddress, s.router)
}

// setUpGin sets up the Gin engine.
func (s *server) setUpGin() {
	engine := gin.Default()
	s.setUpMiddleware(engine)
	s.setUpUIHandler(engine)
	s.setUpAPIHandler(engine)
	s.router = engine
}

// setUpMiddleware installs middleware to the provided Gin engine.
func (s *server) setUpMiddleware(engine *gin.Engine) {
	if s.EnableCORS {
		defer zap.S().Sync()
		zap.S().Info("CORS is enabled")
		engine.Use(CORSMiddleware())
	}
}

// setUpUIHandler sets up the UI handler, which will listen to the default route ("/").
func (s *server) setUpUIHandler(engine *gin.Engine) {
	engine.NoRoute(
		gin.WrapH(
			http.FileServer(
				http.FS(
					u.Must(fs.Sub(ui.Embedded, "dist")),
				),
			),
		),
	)

	api := engine.Group("/api/v1")
	api.GET("/hello", s.handleHello)
}
