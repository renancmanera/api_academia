package handler

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/handler/middleware"
	"github.com/renancmanera/api_academia/internal/handler/treino"
	"github.com/renancmanera/api_academia/internal/handler/usuario"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rota de teste
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Middleware de autenticação JWT
	r.GET("/protegido", middleware.AutenticarJWT(), func(c *gin.Context) {
		c.JSON(200, gin.H{"mensagem": "Você acessou uma rota protegida!"})
	})

	// Rota de cadastro de usuário
	r.POST("/cadastro", usuario.CadastrarUsuario)
	r.POST("/login", usuario.LoginUsuario)

	// Rotas de Treino
	r.POST("/treinos", middleware.AutenticarJWT(), treino.CadastrarTreino)

	return r
}

func RunServer() {
	r := SetupRouter()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	r.Run(addr) // Inicia o servidor na porta definida
}
