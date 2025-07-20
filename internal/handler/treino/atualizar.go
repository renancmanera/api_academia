package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func AtualizarTreino(c *gin.Context) {
	id := c.Param("id")
	var treino domain.Treino
	if err := repository.DB.First(&treino, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Treino não encontrado"})
		return
	}

	var input domain.Treino
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	treino.Nome = input.Nome
	treino.Descricao = input.Descricao

	if err := repository.DB.Save(&treino).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar treino"})
		return
	}

	c.JSON(http.StatusOK, treino)
}
