package exercicio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func AtualizarExercicio(c *gin.Context) {
	id := c.Param("id")
	var exercicio domain.Exercicio
	if err := repository.DB.First(&exercicio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Exercício não encontrado"})
		return
	}

	var input domain.Exercicio
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	exercicio.Nome = input.Nome
	exercicio.Descricao = input.Descricao
	exercicio.GrupoMuscular = input.GrupoMuscular

	if err := repository.DB.Save(&exercicio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar exercício"})
		return
	}

	c.JSON(http.StatusOK, exercicio)
}
