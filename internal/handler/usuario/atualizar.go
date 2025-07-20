package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func AtualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuario domain.Usuario
	if err := repository.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}

	var input domain.Usuario
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	usuario.Nome = input.Nome
	usuario.Email = input.Email
	usuario.Papel = input.Papel

	if err := repository.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar usuário"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
