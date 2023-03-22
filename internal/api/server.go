package api

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

type HTTPServer struct {
	*gin.Engine
	db  *store.Store
	log logger.Logger
}

func NewHTTPServer(db *store.Store, log logger.Logger) (*HTTPServer, error) {
	// initial HTTP setups
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Use(cors.Default())

	return &HTTPServer{Engine: engine, db: db, log: log.WithPrefix("HTTPServer")}, nil
}

func (s *HTTPServer) Start(ctx context.Context) {
	go s.startHandler()
	sysErr := make(chan os.Signal, 1)
	signal.Notify(sysErr,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	select {
	case _ = <-ctx.Done():
		s.log.Infof("terminating due to ctx.Done")
		return
	case sig := <-sysErr:
		s.log.Infof("terminating got `[%v]` signal", sig)
		return
	}
}
