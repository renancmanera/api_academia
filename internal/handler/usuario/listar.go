package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func ListarUsuarios(c *gin.Context) {
	var usuarios []domain.Usuario
	if err := repository.DB.Find(&usuarios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar usuários"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

func BuscarUsuarioPorID(c *gin.Context) {
	id := c.Param("id")
	var usuario domain.Usuario
	if err := repository.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}
