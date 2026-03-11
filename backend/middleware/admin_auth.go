package middleware

import (
	"net/http"
	"strings"

	"myporto-backend/services"
	"myporto-backend/utils"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware(adminService *services.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "unauthorized", "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			utils.Error(c, http.StatusUnauthorized, "unauthorized", "invalid authorization format")
			c.Abort()
			return
		}

		admin, err := adminService.ValidateSession(parts[1])
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, "unauthorized", err.Error())
			c.Abort()
			return
		}

		c.Set("admin_id", admin.ID)
		c.Set("admin_username", admin.Username)
		c.Next()
	}
}
