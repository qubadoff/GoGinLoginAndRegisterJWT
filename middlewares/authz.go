package middlewares

import (
	"LoginAndRegisterApiJWT/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authz() gin.HandlerFunc {
	return func(context *gin.Context) {
		clientToken := context.Request.Header.Get("Authorization")

		if clientToken == "" {
			context.JSON(403, "No Authorization header provided !")
			context.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			context.JSON(400, "Incorrect Format of Authorization Token")
			context.Abort()
			return
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)

		if err != nil {
			context.JSON(401, err.Error())
			context.Abort()
			return
		}

		context.Set("email", claims.Email)
		context.Next()
	}
}
