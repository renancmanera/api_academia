package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

// Estrutura para receber o JSON de cadastro
type CadastroUsuarioInput struct {
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
	Papel string `json:"papel" binding:"required"` // "aluno" ou "instrutor"
}

func CadastrarUsuario(c *gin.Context) {
	var input CadastroUsuarioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	usuario := domain.Usuario{
		Nome:  input.Nome,
		Email: input.Email,
		Senha: input.Senha, // Em produção, faça hash da senha!
		Papel: input.Papel,
	}

	if err := repository.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensagem": "Usuário cadastrado com sucesso!"})
}
