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
		err := fasthttp.ListenAndServe(addr, conn)
		if err != nil {
			logger.Info("fastHttp Listen error", err)
		}
	} else {
		logger.StartupInfo("cfg.config 中的addr配置错误:", addr)
	}

}
