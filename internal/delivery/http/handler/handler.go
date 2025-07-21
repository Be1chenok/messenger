package handler

import (
	"net/http"

	"github.com/Be1chenok/messenger/config"
	"github.com/Be1chenok/messenger/pkg/logger"
	"github.com/gin-gonic/gin"
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

func (h Handler) health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h Handler) Init(r *gin.Engine) {
	router := r.Group("/api/v1")
	router.Use(h.loggingMiddleware())
	router.GET("/health", h.health)
}
