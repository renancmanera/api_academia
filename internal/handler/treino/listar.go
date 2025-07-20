package treino

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/domain"
	"github.com/renancmanera/api_academia/internal/repository"
)

func ListarTreinos(c *gin.Context) {
	papel, _ := c.Get("papel")
	usuarioID, _ := c.Get("usuario_id")
	var treinos []domain.Treino

	if papel == "instrutor" {
		repository.DB.Find(&treinos) // Todos os treinos
	} else {
		repository.DB.Where("usuario_id = ?", usuarioID).Find(&treinos) // Só do aluno
	}
	c.JSON(http.StatusOK, treinos)
}
