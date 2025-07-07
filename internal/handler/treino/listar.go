package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func ListarTreinos(c *gin.Context) {
	usuarioID, existe := c.Get("usuario_id")
	if !existe {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuário não autenticado"})
		return
	}

	var treinos []domain.Treino
	if err := repository.DB.Where("usuario_id = ?", usuarioID.(uint)).Find(&treinos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar treinos"})
		return
	}

	c.JSON(http.StatusOK, treinos)
}
