package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todoapi "github.com/tul1/openapi_go_demo/openapi"
)

type Server struct{}

// PtrString is a helper function to create a pointer to a string.
func PtrString(s string) *string {
	return &s
}

// Implement GetTodos
func (s *Server) GetTodos(c *gin.Context) {
	todos := []todoapi.PostTodosJSONBody{
		{Task: PtrString("Buy groceries")},
		{Task: PtrString("Write OpenAPI spec")},
	}
	c.JSON(http.StatusOK, todos)
}

// Implement PostTodos
func (s *Server) PostTodos(c *gin.Context) {
	var newTodo todoapi.PostTodosJSONBody
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.Task = PtrString("Newly added task")
	c.JSON(http.StatusCreated, newTodo)
}

// Implement GetTodosId
func (s *Server) GetTodosId(c *gin.Context, id int) {
	if id == 1 {
		todo := todoapi.PostTodosJSONBody{Task: PtrString("Buy groceries")}
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
	}
}

// Implement DeleteTodosId
func (s *Server) DeleteTodosId(c *gin.Context, id int) {
	// Simulate deletion logic (e.g., remove from database or memory)
	c.Status(http.StatusNoContent)
}

func main() {
	r := gin.Default()
	server := &Server{}

	// Register routes based on the generated OpenAPI interface
	todoapi.RegisterHandlers(r, server)

	_ = r.Run(":8080")
}
