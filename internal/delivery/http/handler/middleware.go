package handler

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()
		userAgent := c.Request.UserAgent()
		errs := getErrs(c)

		if errs != nil {
			if len(errs) > 1 {
				h.logger.Errorsf("errors", errs, "method: %s path: %s ip: %s user-agent: %s status: %d", method, path, ip, userAgent, status)
				return
			}
			h.logger.Errorf(errs[0], "method: %s path: %s ip: %s user-agent: %s status: %d", method, path, ip, userAgent, status)
			return
		}
		h.logger.Infof("method: %s path: %s ip: %s user-agent: %s status: %d", method, path, ip, userAgent, status)
	}
}

func getErrs(c *gin.Context) []error {
	ginErrs := c.Errors
	if len(ginErrs) == 0 {
		return nil
	}
	errorStrings := make([]error, len(ginErrs))
	for i, err := range ginErrs {
		errorStrings[i] = err
	}
	return errorStrings
}
