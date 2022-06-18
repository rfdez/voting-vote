package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/voting-vote/internal/platform/server/http/handler/health"
	"github.com/rfdez/voting-vote/internal/platform/server/http/middleware/logging"
	"github.com/rfdez/voting-vote/internal/platform/server/http/middleware/recovery"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration
}

func NewServer(ctx context.Context, host string, port uint, shutdownTimeout time.Duration) (context.Context, Server) {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		shutdownTimeout: shutdownTimeout,
	}

	srv.registerRoutes()

	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	// Register middleware
	s.engine.Use(recovery.Middleware(), logging.Middleware())

	// Register routes
	s.engine.GET("/ping", health.PingHandler())
}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("HTTP server listening on %s", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server stopped: %v", err)
		}
	}()

	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	log.Printf("HTTP server shutting down...")

	return srv.Shutdown(ctxShutdown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
