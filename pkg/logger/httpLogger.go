package httpLogger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func NewCustomResponseWriter(w gin.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{w, &bytes.Buffer{}}
}

func (crw *CustomResponseWriter) Write(b []byte) (int, error) {
	crw.Body.Write(b)
	return crw.ResponseWriter.Write(b)
}

func CustomLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.URL.Path == "/health" {
			c.Next()
			return
		}

		startTime := time.Now()

		crw := NewCustomResponseWriter(c.Writer)
		c.Writer = crw

		c.Next()

		duration := time.Since(startTime)
		statusCode := crw.Status()

		logger := log.Info()
		if statusCode >= http.StatusBadRequest {
			logger = log.Error()
		}

		logger.Str("protocol", "http").
			Str("client_ip", c.ClientIP()).
			//Int64("request_size", c.Request.ContentLength).
			//Int("response_size", crw.Size()).
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).
			Int("status_code", statusCode).
			//Str("response_body", crw.Body.String()).
			Dur("duration", duration).
			Msg("received a HTTP request")
	}
}
