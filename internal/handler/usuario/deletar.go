package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DB.Delete(&domain.Usuario{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar usuário"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário deletado com sucesso"})
}
