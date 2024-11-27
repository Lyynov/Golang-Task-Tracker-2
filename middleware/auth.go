package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		//ambil cookie
		tokenStr, err := ctx.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.Redirect(http.StatusSeeOther, "/login")
			return
		}
		//parsing jwt token untuk mendapatkan claims
		var claims model.Claims
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) { return model.JwtKey, nil })
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//menyimpan nilai email ke context
		ctx.Set("email", claims.Email)
		ctx.Next()

	})
}
