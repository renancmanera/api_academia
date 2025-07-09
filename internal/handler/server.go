package handler

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/renancmanera/api_academia/internal/handler/exercicio"
	"github.com/renancmanera/api_academia/internal/handler/middleware"
	"github.com/renancmanera/api_academia/internal/handler/treino"
	"github.com/renancmanera/api_academia/internal/handler/usuario"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

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

	// Rotas de autenticação
	r.POST("/cadastro", usuario.CadastrarUsuario)
	r.POST("/login", usuario.LoginUsuario)

	// CRUD de Usuário
	r.GET("/usuarios", middleware.AutenticarJWT(), usuario.ListarUsuarios)
	r.GET("/usuarios/:id", middleware.AutenticarJWT(), usuario.BuscarUsuarioPorID)
	r.PUT("/usuarios/:id", middleware.AutenticarJWT(), usuario.AtualizarUsuario)
	r.DELETE("/usuarios/:id", middleware.AutenticarJWT(), usuario.DeletarUsuario)

	// CRUD de Treino
	r.POST("/treinos", middleware.AutenticarJWT(), treino.CadastrarTreino)
	r.GET("/treinos", middleware.AutenticarJWT(), treino.ListarTreinos)
	r.GET("/treinos/:id", middleware.AutenticarJWT(), treino.BuscarTreinoPorID)
	r.PUT("/treinos/:id", middleware.AutenticarJWT(), treino.AtualizarTreino)
	r.DELETE("/treinos/:id", middleware.AutenticarJWT(), treino.DeletarTreino)

	// CRUD de Exercício
	r.POST("/exercicios", middleware.AutenticarJWT(), exercicio.CadastrarExercicio)
	r.GET("/exercicios", middleware.AutenticarJWT(), exercicio.ListarExercicios)
	r.GET("/exercicios/:id", middleware.AutenticarJWT(), exercicio.BuscarExercicioPorID)
	r.PUT("/exercicios/:id", middleware.AutenticarJWT(), exercicio.AtualizarExercicio)
	r.DELETE("/exercicios/:id", middleware.AutenticarJWT(), exercicio.DeletarExercicio)

	// Associação de exercícios ao treino
	r.POST("/treinos/:id/exercicios", middleware.AutenticarJWT(), treino.AssociarExercicios)
	r.GET("/treinos/:id/exercicios", middleware.AutenticarJWT(), treino.ListarExerciciosDoTreino)

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
