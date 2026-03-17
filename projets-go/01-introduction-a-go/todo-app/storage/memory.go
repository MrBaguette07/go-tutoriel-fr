package storage

import (
	"sync"
	"todo-app/models"
)

type MemoryStorage struct {
	todos  []models.Todo
	nextID int
	mu     sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos:  make([]models.Todo, 0),
		nextID: 1,
	}
}

func (s *MemoryStorage) Create(todo *models.Todo) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	s.nextID++
	s.todos = append(s.todos, *todo)
}

func (s *MemoryStorage) GetAll() []models.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return append([]models.Todo(nil), s.todos...)
}
