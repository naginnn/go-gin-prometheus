package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
	"time"
)

func WithMetric(r gin.IRoutes) {
	gaugeMetric := &ginmetrics.Metric{
		Type:        ginmetrics.Gauge,
		Name:        "syka_gauge_metric",
		Description: "an example of syka_gauge_metric",
		Labels:      []string{"label1"},
	}
	// Add metric to global monitor object
	_ = ginmetrics.GetMonitor().AddMetric(gaugeMetric)

	m := ginmetrics.GetMonitor()
	upTime := time.Now().Unix()
	go func() {
		for {
			_ = ginmetrics.GetMonitor().GetMetric("syka_gauge_metric").SetGaugeValue([]string{"uptime_secs"}, float64(time.Now().Unix()-upTime))
			time.Sleep(1)
		}
	}()
	m.SetMetricPath("/metrics")
	m.Use(r)
}

func WithMetricServer(r gin.IRoutes, port string) {

	gaugeMetric := &ginmetrics.Metric{
		Type:        ginmetrics.Gauge,
		Name:        "syka_gauge_metric",
		Description: "an example of syka_gauge_metric",
		Labels:      []string{"label1"},
	}
	// Add metric to global monitor object
	_ = ginmetrics.GetMonitor().AddMetric(gaugeMetric)
	metricRouter := gin.Default()
	m := ginmetrics.GetMonitor()
	upTime := time.Now().Unix()
	go func() {
		for {
			_ = ginmetrics.GetMonitor().GetMetric("syka_gauge_metric").SetGaugeValue([]string{"uptime_secs"}, float64(time.Now().Unix()-upTime))
			time.Sleep(1)
		}
	}()
	m.SetMetricPath("/metric")
	m.UseWithoutExposingEndpoint(r)
	m.Expose(metricRouter)

	go func() {
		_ = metricRouter.Run("0.0.0.0:" + port)
	}()
}

func main() {
	r := gin.New()
	//WithMetric(r)
	WithMetricServer(r, "57475")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run("0.0.0.0:2223")
}
