package monitor

import (
	"regexp"

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
}
