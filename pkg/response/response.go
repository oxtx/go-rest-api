package response

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func JSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

func Fail(c *gin.Context, status int, errMsg, detail string) {
	c.JSON(status, ErrorResponse{Error: errMsg, Message: detail})
}
