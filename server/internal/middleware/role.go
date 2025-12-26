package middleware

import (
	"net/http"

	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func HostOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")

		var user models.User
		if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "User not found")
			c.Abort()
			return
		}

		if user.Role != models.RoleHost {
			utils.ErrorResponse(c, http.StatusForbidden, "Host access required")
			c.Abort()
			return
		}

		c.Next()
	}
}
