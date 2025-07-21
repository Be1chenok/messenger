package handler

import (
	"net/http"

	"github.com/Be1chenok/messenger/config"
	"github.com/Be1chenok/messenger/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	conf   *config.Handler
	logger logger.Logger
}

func New(conf *config.Handler, logger logger.Logger) *Handler {
	return &Handler{
		conf:   conf,
		logger: logger,
	}
}

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Init(r *mux.Router) {
	router := r.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/health", h.Health)
}
