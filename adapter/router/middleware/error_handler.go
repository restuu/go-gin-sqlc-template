package middleware

import (
	"database/sql"
	"errors"
	"go-gin-sqlc-template/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler error handler middleware
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Written() {
			return
		}

		err := c.Errors.Last()
		if err == nil {
			return
		}

		switch {
		case errors.Is(err, sql.ErrNoRows),
			errors.Is(err, domain.ErrNoData):
			c.JSON(http.StatusOK, gin.H{"error": domain.ErrNoData.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": domain.ErrInternalServerError.Error()})
		}
	}
}
