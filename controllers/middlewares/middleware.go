package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/mygoapp2nd/config"
	"github.com/kalleriakronos24/mygoapp2nd/constants"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
)

func AuthMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}
	claims, err := parseToken(token, config.AppConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Error: err.Error()})
		return
	}
	c.Set(constants.IsAuthenticatedKey, true)
	c.Set(constants.UserIDKey, claims.ID)
	c.Next()
}

func parseToken(tokenString, secret string) (claims dto.JWTClaims, err error) {
	if token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil || !token.Valid {
		return dto.JWTClaims{}, fmt.Errorf("invalid token. %s", err)
	}
	return
}
