package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteRegistrar is implemented by any handler that can register its routes
// onto a Gin router. This keeps the server package decoupled from concrete
// handler types.
type RouteRegistrar interface {
	RegisterRoutes(router gin.IRouter)
}

// maxBytesMiddleware limits incoming request body size to 1 MB.
// This prevents memory exhaustion from oversized bodies on all API routes.
func maxBytesMiddleware(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		c.Next()
	}
}

func New(
	port int,
	registrars ...RouteRegistrar,
) *http.Server {
	router := gin.Default()

	router.Use(maxBytesMiddleware(1 << 20)) // 1 MB

	router.GET("/api/v1/", func(c *gin.Context) {
		c.String(http.StatusOK, "Qubit is running")
	})

	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "qubit-api",
		})
	})

	for _, r := range registrars {
		r.RegisterRoutes(router)
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
}
