package exercicio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

type CadastroExercicioInput struct {
	Nome          string `json:"nome" binding:"required"`
	Descricao     string `json:"descricao"`
	GrupoMuscular string `json:"grupo_muscular"`
}

func CadastrarExercicio(c *gin.Context) {
	var input CadastroExercicioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	exercicio := domain.Exercicio{
		Nome:          input.Nome,
		Descricao:     input.Descricao,
		GrupoMuscular: input.GrupoMuscular,
	}

	if err := repository.DB.Create(&exercicio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensagem": "Exercício cadastrado com sucesso!"})
}
