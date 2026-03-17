package main

import (
	"todo-app/handlers"
	"todo-app/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialiser le stockage
	store := storage.NewMemoryStorage()

	// Initialiser les handlers
	todoHandler := handlers.NewTodoHandler(store)

	// Configuration du routeur
	r := gin.Default()

	// Routes API
	api := r.Group("/api/v1")
	{
		api.GET("/todos", todoHandler.GetTodos)
		api.POST("/todos", todoHandler.CreateTodo)
	}

	// Servir les fichiers statiques
	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Démarrer le serveur
	r.Run(":8080")
}
