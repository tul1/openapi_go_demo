package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	todoapi "github.com/tul1/openapi_go_demo/openapi"
)

// Todo represents a task item
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// Server holds the data and methods for the API
type Server struct {
	mu    sync.Mutex
	todos []Todo
}

// PtrString is a helper function to create a pointer to a string.
func PtrString(s string) *string {
	return &s
}

// GetTodos retrieves all TODOs
func (s *Server) GetTodos(c *gin.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []todoapi.PostTodosJSONBody
	for _, todo := range s.todos {
		result = append(result, todoapi.PostTodosJSONBody{Task: PtrString(todo.Task)})
	}
	c.JSON(http.StatusOK, result)
}

// PostTodos creates a new TODO item
func (s *Server) PostTodos(c *gin.Context) {
	var newTodo todoapi.PostTodosJSONBody
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := len(s.todos) + 1
	todo := Todo{ID: id, Task: *newTodo.Task}
	s.todos = append(s.todos, todo)

	c.JSON(http.StatusCreated, todo)
}

// GetTodosId retrieves a TODO by ID
func (s *Server) GetTodosId(c *gin.Context, id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, todo := range s.todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
}

// DeleteTodosId deletes a TODO by ID
func (s *Server) DeleteTodosId(c *gin.Context, id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
}

func main() {
	r := gin.Default()
	server := &Server{todos: []Todo{
		{ID: 1, Task: "Buy groceries"},
		{ID: 2, Task: "Write OpenAPI spec"},
	}}

	// Register routes based on the generated OpenAPI interface
	todoapi.RegisterHandlers(r, server)

	_ = r.Run(":8080")
}
