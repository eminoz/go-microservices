package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/eminoz/go-microservices/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func IsAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//verifytoken

		authToken := ctx.Request.Header.Get("Authorization")
		if authToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"e": "please log in"})
			ctx.Abort()
		}
		//get secretKey from config file
		var mySigningKey = config.GetConfig().AppSecret
		//parse token with secret key
		token, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "not auth",
				})
			}
			return []byte(mySigningKey), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"e": err.Error()})
			ctx.Abort()
		}
		//We can look inside of token
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["_id"]
		fmt.Println(userId)
		//if valid token move
		if token.Valid == true {
			ctx.Next()

		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

	}
}
