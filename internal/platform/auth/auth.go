package auth

import (
	"strings"

	"shadow/internal/config"

	"shadow/internal/registrations"

	"shadow/internal/helpers/messages"

	"github.com/dgrijalva/jwt-go.git"
	"github.com/gin-gonic/gin"
)

type E = messages.Error
type S = messages.Success

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		secret := config.GetString("auth.secret")

		if len(strings.TrimSpace(secret)) == 0 {
			c.JSON(E{401, "ERR012", nil}.Gen())
			c.Abort()
			return
		}

		tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok &&
				token.Valid &&
				registrations.CheckUserTokenMatch(claims["Username"].(string), claims["Email"].(string), tokenString) {
				c.Set("Username", claims["Username"])
				c.Set("Email", claims["Email"])
			} else {
				c.JSON(E{401, "ERR013", nil}.Gen())
				c.Abort()
				return
			}
		} else if verify, ok := err.(*jwt.ValidationError); ok {
			if verify.Errors&jwt.ValidationErrorMalformed != 0 {
				c.JSON(E{401, "ERR014", nil}.Gen())
				c.Abort()
				return
			} else if verify.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.JSON(E{401, "ERR015", nil}.Gen())
				c.Abort()
				return
			} else {
				c.JSON(E{401, "ERR016", nil}.Gen())
				c.Abort()
				return
			}
			return
		} else {
			c.JSON(E{401, "ERR016", nil}.Gen())
			c.Abort()
			return
		}
		c.Next()
	}
}
