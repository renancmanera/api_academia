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

	// CORS para permitir acesso do frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	// Rota de teste
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Rota protegida de teste
	r.GET("/protegido", middleware.AutenticarJWT(), func(c *gin.Context) {
		c.JSON(200, gin.H{"mensagem": "Você acessou uma rota protegida!"})
	})

	// Rotas de autenticação
	r.POST("/cadastro", usuario.CadastrarUsuario)
	r.POST("/login", usuario.LoginUsuario)

	// CRUD de Usuário
	// Visualização (GET) – ambos
	r.GET("/usuarios", middleware.AutenticarJWT(), usuario.ListarUsuarios)
	r.GET("/usuarios/:id", middleware.AutenticarJWT(), usuario.BuscarUsuarioPorID)
	// Gerenciamento (POST/PUT/DELETE) - só instrutor
	r.PUT("/usuarios/:id", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), usuario.AtualizarUsuario)
	r.DELETE("/usuarios/:id", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), usuario.DeletarUsuario)

	// CRUD de Treino
	// Visualização (GET) – ambos
	r.GET("/treinos", middleware.AutenticarJWT(), treino.ListarTreinos)
	r.GET("/treinos/:id", middleware.AutenticarJWT(), treino.BuscarTreinoPorID)
	r.GET("/treinos/:id/exercicios", middleware.AutenticarJWT(), treino.ListarExerciciosDoTreino)
	// Gerenciamento (POST/PUT/DELETE) – só instrutor
	r.POST("/treinos", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), treino.CadastrarTreino)
	r.PUT("/treinos/:id", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), treino.AtualizarTreino)
	r.DELETE("/treinos/:id", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), treino.DeletarTreino)
	r.POST("/treinos/:id/exercicios", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), treino.AssociarExercicios)
	r.DELETE("/treinos/:id/exercicios/:exercicio_id", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), treino.RemoverExercicioDoTreino)

	// CRUD de Exercício
	// Visualização (GET) – ambos
	r.GET("/exercicios", middleware.AutenticarJWT(), exercicio.ListarExercicios)
	r.GET("/exercicios/:id", middleware.AutenticarJWT(), exercicio.BuscarExercicioPorID)
	// Gerenciamento (POST/PUT/DELETE) – só instrutor
	r.POST("/exercicios", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), exercicio.CadastrarExercicio)
	r.PUT("/exercicios/:id", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), exercicio.AtualizarExercicio)
	r.DELETE("/exercicios/:id", middleware.AutenticarJWT(), middleware.PermitirPapeis("instrutor"), exercicio.DeletarExercicio)

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
