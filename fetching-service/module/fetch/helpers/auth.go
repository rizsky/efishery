package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

//ExtractToken get token extraction
func ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header["Authorization"]
	if len(bearToken) == 0 {
		HTTPResponseError(c, nil, 401, "authorization header is null")
	}

	strArr := strings.Split(bearToken[0], " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
