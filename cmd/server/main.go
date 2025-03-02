package main

import (
	"github.com/gin-gonic/gin"
	"github.com/project-Grok3_Golang/internal/database"
	"github.com/project-Grok3_Golang/internal/handlers"
	"github.com/project-Grok3_Golang/internal/middleware"
)

func main() {
	database.InitDB()
	handlers.InitAIClient()
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.POST("/chat_xi", handlers.Predict)
	r.Run(":8080")
}
