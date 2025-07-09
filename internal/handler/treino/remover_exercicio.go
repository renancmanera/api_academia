package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

// Remove a associação de um exercício a um treino (não deleta o exercício do banco)
func RemoverExercicioDoTreino(c *gin.Context) {
	treinoID := c.Param("id")
	exercicioID := c.Param("exercicio_id")

	var treino domain.Treino
	if err := repository.DB.Preload("Exercicios").First(&treino, treinoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Treino não encontrado"})
		return
	}

	var exercicio domain.Exercicio
	if err := repository.DB.First(&exercicio, exercicioID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Exercício não encontrado"})
		return
	}

	if err := repository.DB.Model(&treino).Association("Exercicios").Delete(&exercicio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao remover associação"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Exercício removido do treino com sucesso!"})
}
