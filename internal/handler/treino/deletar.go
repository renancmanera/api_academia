package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func DeletarTreino(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DB.Delete(&domain.Treino{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar treino"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Treino deletado com sucesso"})
}
