package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func loggerMiddleware(ctx *gin.Context) {
	// calc latency
	now := time.Now()
	ctx.Next()
	end := time.Since(now)

	msg := ctx.Errors.String()
	if msg == "" {
		msg = "request recieved"
	}

	path := getWholePath(ctx)

	log.Info().
		Str("method", ctx.Request.Method).
		Str("path", path).
		Int("status", ctx.Writer.Status()).
		Dur("latency", end).
		Str("client_ip", ctx.ClientIP()).
		Msg(msg)
}

func getWholePath(ctx *gin.Context) (path string) {

	path = ctx.FullPath()
	if path == "" {
		path = ctx.Request.URL.Path
	}
	if ctx.Request.URL.RawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, ctx.Request.URL.RawQuery)
	}
	return
}
