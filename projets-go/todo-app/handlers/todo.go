package handlers

import (
	"net/http"
	"todo-app/models"
	"todo-app/storage"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	storage *storage.MemoryStorage
}

func NewTodoHandler(s *storage.MemoryStorage) *TodoHandler {
	return &TodoHandler{storage: s}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.NewTodo(req.Title)
	h.storage.Create(todo)

	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos := h.storage.GetAll()
	c.JSON(http.StatusOK, todos)
}
