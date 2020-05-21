package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/efishery/fetchin-service/module/fetch/helpers"

	"github.com/gin-gonic/gin"
)

type JwtClaimed struct {
	Name      string
	Phone     string
	Role      string
	Timestamp string
}

const key = "kikyn0secret"

//AuthenticationRequired use as middleware and get the type role
func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedJwt := JwtClaimed{}
		tokenString := helpers.ExtractToken(c)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil {
			helpers.HTTPResponseError(c, err, 401, "token is invalid, cannot parsed")
			c.Abort()
			return
		}

		//Claim the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helpers.HTTPResponseError(c, err, 401, "token is invalid, cannot parsed")
			c.Abort()
			return
		}

		//assigning to struct
		decodedJwt.Name = claims["name"].(string)
		decodedJwt.Role = claims["role"].(string)
		decodedJwt.Phone = claims["phone"].(string)
		decodedJwt.Timestamp = claims["timestamp"].(string)

		c.Set("decodedJwt", decodedJwt)

		if len(auths) != 0 {
			if auths[0] == "admin" {
				if decodedJwt.Role != "admin" {
					helpers.HTTPResponseError(c, err, 403, "accessed just for admin only")
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}
