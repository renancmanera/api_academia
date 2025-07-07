package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func ListarExerciciosDoTreino(c *gin.Context) {
	treinoID := c.Param("id")
	var treino domain.Treino
	if err := repository.DB.Preload("Exercicios").First(&treino, treinoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Treino não encontrado"})
		return
	}
	c.JSON(http.StatusOK, treino.Exercicios)
}
