package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware para proteger rotas usando JWT e disponibilizar o usuario_id
func AutenticarJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Token não informado"})
			c.Abort()
			return
		}

		partes := strings.Split(authHeader, " ")
		if len(partes) != 2 || partes[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Formato do token inválido"})
			c.Abort()
			return
		}

		tokenString := partes[1]
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido ou expirado"})
			c.Abort()
			return
		}

		// Extração do usuario_id das claims do token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// O jwt retorna números como float64, por isso a conversão
			if id, existe := claims["usuario_id"]; existe {
				if idFloat, ok := id.(float64); ok {
					c.Set("usuario_id", uint(idFloat))
				}
			}
		}

		c.Next()
	}
}
