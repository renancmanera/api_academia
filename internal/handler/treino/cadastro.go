package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

type CadastroTreinoInput struct {
	Nome      string `json:"nome" binding:"required"`
	Descricao string `json:"descricao"`
}

func CadastrarTreino(c *gin.Context) {
	var input CadastroTreinoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	usuarioID, existe := c.Get("usuario_id")
	if !existe {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuário não autenticado"})
		return
	}

	treino := domain.Treino{
		Nome:      input.Nome,
		Descricao: input.Descricao,
		UsuarioID: usuarioID.(uint),
	}

	if err := repository.DB.Create(&treino).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensagem": "Treino cadastrado com sucesso!"})
}
