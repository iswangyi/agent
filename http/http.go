package http

import (
	"agent/logger"
	"agent/settings"
	"github.com/valyala/fasthttp"
)

func Start() {
	addr := settings.Config().Http.Listen
	if addr != "" {
		conn := func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/v1/check/health":
				agentStatusRoutes(ctx)

			default:
				ctx.Error("Not Found", fasthttp.StatusNotFound)
			}
		}
		conn = fasthttp.CompressHandler(conn)
		fasthttp.ListenAndServe(addr, conn)
	} else {
		logger.StartupInfo("cfg.config 中的addr配置错误", addr)
	}

}
