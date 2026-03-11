package middleware

import (
	"net/http"
	"strings"

	"myporto-backend/config"
	"myporto-backend/utils"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := buildAllowedOrigins()
	allowedMethods := "GET, POST, PUT, DELETE, OPTIONS"
	allowedHeaders := "Origin, Content-Type, Accept, Authorization"

	return func(c *gin.Context) {
		origin := strings.TrimSpace(c.GetHeader("Origin"))
		if origin != "" {
			if !isOriginAllowed(origin, allowedOrigins) {
				if c.Request.Method == http.MethodOptions {
					c.AbortWithStatus(http.StatusForbidden)
					return
				}

				utils.Error(c, http.StatusForbidden, "origin not allowed", origin)
				c.Abort()
				return
			}

			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
		}

		c.Header("Access-Control-Allow-Methods", allowedMethods)
		c.Header("Access-Control-Allow-Headers", allowedHeaders)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func buildAllowedOrigins() map[string]struct{} {
	raw := config.GetEnv("CORS_ALLOWED_ORIGINS", "")
	origins := make(map[string]struct{})

	if raw == "" {
		// Safe development defaults. Production should always use explicit env configuration.
		origins["http://localhost:5173"] = struct{}{}
		origins["http://127.0.0.1:5173"] = struct{}{}
		origins["http://localhost:4173"] = struct{}{}
		origins["http://127.0.0.1:4173"] = struct{}{}
		return origins
	}

	for _, item := range strings.Split(raw, ",") {
		origin := strings.TrimSpace(item)
		if origin == "" {
			continue
		}
		origins[origin] = struct{}{}
	}

	return origins
}

func isOriginAllowed(origin string, allowedOrigins map[string]struct{}) bool {
	_, ok := allowedOrigins[origin]
	return ok
}
