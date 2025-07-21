package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Be1chenok/messenger/config"
	"github.com/Be1chenok/messenger/pkg/logger"
	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
	conf       *config.HTTPServer
	logger     logger.Logger
}

func New(conf *config.HTTPServer, logger logger.Logger, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%d", conf.Port),
			ReadTimeout:    conf.ReadTimeout,
			WriteTimeout:   conf.WriteTimeout,
			MaxHeaderBytes: conf.MaxHeaderBytes,
			Handler:        handler,
		},
		conf:   conf,
		logger: logger.WithStr("component", "server"),
	}
}

func InitHandler(fs ...func(*mux.Router)) http.Handler {
	r := mux.NewRouter()

	for _, f := range fs {
		f(r)
	}

	return r
}

func (s Server) Run() {
	serverErr := make(chan error, 1)

	go func() {
		s.logger.Infof("server listening on port: %d", s.conf.Port)

		if err := s.httpServer.ListenAndServe(); err != nil {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	select {
	case <-quit:
		s.logger.Info("shutdown server ...")

		ctx, cancel := context.WithTimeout(context.Background(), s.conf.ServerShutdownTimeout)
		defer cancel()

		if err := s.httpServer.Shutdown(ctx); err != nil {
			s.logger.Error(err, "shutdown")
		}

		s.logger.Info("server stopped")
	case err := <-serverErr:
		s.logger.Error(err, "server error")
	}
}
