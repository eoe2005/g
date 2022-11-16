package gmonitor

import (
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(c *gin.Engine) {
	// Add Go module build info.
	localReg.MustRegister(collectors.NewBuildInfoCollector())
	localReg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))
	c.Any("/metrics", gin.WrapH(promhttp.HandlerFor(
		localReg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
			Registry:          localReg,
		},
	)))
	c.Use(func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		tags := map[string]string{
			"method": ctx.Request.Method,
			"path":   ctx.Request.URL.Path,
		}
		Counter("api_total", "api_total", 1, tags)
		Summary("api_cost", "api_cost", float64(time.Now().Sub(start)), tags)
	})
}
