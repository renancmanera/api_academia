package exercicio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func DeletarExercicio(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DB.Delete(&domain.Exercicio{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar exercício"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Exercício deletado com sucesso"})
}
