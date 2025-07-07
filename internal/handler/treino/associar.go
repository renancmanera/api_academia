package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

type AssociarExerciciosInput struct {
	ExerciciosIDs []uint `json:"exercicios_ids" binding:"required"`
}

func AssociarExercicios(c *gin.Context) {
	treinoID := c.Param("id")
	var input AssociarExerciciosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var treino domain.Treino
	if err := repository.DB.Preload("Exercicios").First(&treino, treinoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Treino não encontrado"})
		return
	}

	var exercicios []domain.Exercicio
	if err := repository.DB.Where("id IN ?", input.ExerciciosIDs).Find(&exercicios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar exercícios"})
		return
	}

	if err := repository.DB.Model(&treino).Association("Exercicios").Replace(&exercicios); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao associar exercícios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Exercícios associados ao treino com sucesso!"})
}
