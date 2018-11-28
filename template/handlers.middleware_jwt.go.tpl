package handlers

import (
	"{{.OutPackage}}/models"
	"github.com/gin-gonic/gin"
	"strings"
)

var jwtMiddleware = jwtCheck()

const tokenPrefix = "Bearer "

func jwtCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), tokenPrefix, "", 1)
		user, err := models.JwtParseUser(token)
		if err != nil {
			handleError(c, err)
			//c.Abort has been called
			return
		}
		//store the user Model in the context
		c.Set("user", user)
		c.Next()
		// after request
	}
}
