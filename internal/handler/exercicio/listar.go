package exercicio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func ListarExercicios(c *gin.Context) {
	var exercicios []domain.Exercicio
	if err := repository.DB.Find(&exercicios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar exercícios"})
		return
	}
	c.JSON(http.StatusOK, exercicios)
}

func BuscarExercicioPorID(c *gin.Context) {
	id := c.Param("id")
	var exercicio domain.Exercicio
	if err := repository.DB.First(&exercicio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Exercício não encontrado"})
		return
	}
	c.JSON(http.StatusOK, exercicio)
}
