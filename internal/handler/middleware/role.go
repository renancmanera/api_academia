package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// funcao para permitir apenas usuários com papéis especificos, como aluno e instrutor
func PermitirPapeis(papeisPermitidos ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		papel, existe := c.Get("papel")
		if !existe {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Papel do usuário não encontrado"})
			c.Abort()
			return
		}
		for _, permitido := range papeisPermitidos {
			if papel == permitido {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"erro": "Permissão negada para este papel"})
		c.Abort()
	}
}
