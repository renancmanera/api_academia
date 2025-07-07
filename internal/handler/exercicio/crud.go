package exercicio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

// Lisar exercícios
func ListarExercicios(c *gin.Context) {
	var exercicios []domain.Exercicio
	if err := repository.DB.Find(&exercicios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar exercícios"})
		return
	}
	c.JSON(http.StatusOK, exercicios)
}

// Buscar exercício por ID
func BuscarExercicioPorID(c *gin.Context) {
	id := c.Param("id")
	var exercicio domain.Exercicio
	if err := repository.DB.First(&exercicio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Exercício não encontrado"})
		return
	}
	c.JSON(http.StatusOK, exercicio)
}

// Atualizar exercício
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

// Deletar exercício
func DeletarExercicio(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DB.Delete(&domain.Exercicio{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar exercício"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Exercício deletado com sucesso"})
}
