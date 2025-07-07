package main

import (
	"github.com/gin-gonic/gin"
	"github.com/renancmanera/api_academia/internal/handler"
	"github.com/renancmanera/api_academia/internal/repository"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	repository.InitDB()
	handler.RunServer()
}
