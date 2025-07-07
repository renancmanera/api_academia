package usuario

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

type LoginInput struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

func LoginUsuario(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var usuario domain.Usuario
	if err := repository.DB.Where("email = ?", input.Email).First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "E-mail ou senha inválidos"})
		return
	}

	// Comparação de senha (com hash, se implementado)
	if usuario.Senha != input.Senha {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "E-mail ou senha inválidos"})
		return
	}

	// Geração do token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario_id": usuario.ID,
		"nome":       usuario.Nome,
		"papel":      usuario.Papel,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Login realizado com sucesso!",
		"token":    tokenString,
	})
}
