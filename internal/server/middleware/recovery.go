package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxtx/go-rest-api/pkg/response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				response.Fail(c, http.StatusInternalServerError, "panic", "internal server error")
			}
		}()
		c.Next()
	}
}
