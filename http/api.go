package http

import (
	"agent/logger"
	"fmt"
	"github.com/valyala/fasthttp"
)

func agentStatusRoutes(ctx *fasthttp.RequestCtx) {
	_, err := fmt.Fprintf(ctx, "已启动")
	if err != nil {
		logger.Info("health check failed")
	}
}
