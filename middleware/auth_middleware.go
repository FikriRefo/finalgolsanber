package middleware

import (
	"socmed/errorhandle"
	"socmed/helper" // Tambahkan ini untuk mengimpor helper

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandle.HandleError(c, &errorhandle.UnauthorizedError{Message: "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString) // Perbaikan: panggil helper.ValidateToken
		if err != nil {
			errorhandle.HandleError(c, &errorhandle.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
