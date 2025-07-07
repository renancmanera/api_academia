package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

// Listar todos os usuários
func ListarUsuarios(c *gin.Context) {
	var usuarios []domain.Usuario
	if err := repository.DB.Find(&usuarios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar usuários"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// Buscar usuário por ID
func BuscarUsuarioPorID(c *gin.Context) {
	id := c.Param("id")
	var usuario domain.Usuario
	if err := repository.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

// Atualizar usuário
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

// Deletar usuário
func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DB.Delete(&domain.Usuario{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar usuário"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário deletado com sucesso"})
}
