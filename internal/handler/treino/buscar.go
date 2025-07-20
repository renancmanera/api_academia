package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func BuscarTreinoPorID(c *gin.Context) {
	id := c.Param("id")
	var treino domain.Treino
	if err := repository.DB.First(&treino, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Treino não encontrado"})
		return
	}
	c.JSON(http.StatusOK, treino)
}
