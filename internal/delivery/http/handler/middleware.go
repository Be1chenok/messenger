package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h Handler) loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		logger := h.logger.
			WithStr("method", c.Request.Method).
			WithStr("path", c.Request.URL.Path).
			WithStr("ip", c.ClientIP()).
			WithStr("user-agent", c.Request.UserAgent()).
			WithInt("status", status).
			WithStr("latency", latency.String())
			// TODO:
		switch {
		case status >= http.StatusInternalServerError:
			logger.Error(c.Errors.Last(), "server error")
		case status >= http.StatusBadRequest:
			errs := c.Errors.Errors()
			if len(errs) != 0 {
				logger = logger.WithStrs("errors", errs...)
			}
			logger.Info("client error")
		default:
			logger.Info("request handled")
		}
	}
}
