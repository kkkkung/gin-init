package e

import (
	"net/http"

	"gin-init/models"
	"github.com/gin-gonic/gin"
)

// NotFound creates a gin middleware for handling page not found.
func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &models.Error{
			Error:            http.StatusText(http.StatusNotFound),
			ErrorCode:        http.StatusNotFound,
			ErrorDescription: "page not found",
		})
	}
}
