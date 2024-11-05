`Prometheus` is an open-source monitoring and alerting toolkit widely used for monitoring applications and infrastructure.
It collects metrics from configured targets at specified intervals, evaluates rule expressions, and can trigger alerts if certain conditions are met.

```go
package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/gin-gonic/gin"
)

var (
    // Define a new counter metric
    requestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests.",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    // Register the counter with Prometheus
    prometheus.MustRegister(requestCounter)
}

func main() {
    r := gin.Default()

    // Define a simple route
    r.GET("/", func(c *gin.Context) {
        requestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc() // Increment counter
        c.String(http.StatusOK, "Hello, World!")
    })

    // Serve the Prometheus metrics on /metrics endpoint
    r.GET("/metrics", promhttp.Handler())

    // Start the server
    r.Run(":8080")
}
```
