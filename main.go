package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tul1/openapi_go_demo/db"
	"github.com/tul1/openapi_go_demo/model"
	todoapi "github.com/tul1/openapi_go_demo/openapi"
)

type DatabaseHandler interface {
	GetTodos() ([]model.Todo, error)
	AddTodo(todo model.Todo) error
	GetTodoByID(id int) (model.Todo, error)
	DeleteTodoByID(id int) error
}

type Server struct {
	dbHandler DatabaseHandler
}

func PtrString(s string) *string {
	return &s
}

func (s *Server) GetTodos(c *gin.Context) {
	todos, err := s.dbHandler.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (s *Server) PostTodos(c *gin.Context) {
	var newTodo model.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.dbHandler.AddTodo(newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func (s *Server) GetTodosId(c *gin.Context, id int) {
	todo, err := s.dbHandler.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (s *Server) DeleteTodosId(c *gin.Context, id int) {
	if err := s.dbHandler.DeleteTodoByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func main() {
	dsn := "todos.db"
	useGORM := os.Getenv("USE_GORM") == "true"

	var dbHandler DatabaseHandler
	var err error
	if useGORM {
		dbHandler, err = db.NewGORMHandler(dsn)
	} else {
		dbHandler, err = db.NewSQLHandler(dsn)
	}
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	server := &Server{dbHandler: dbHandler}

	todoapi.RegisterHandlers(r, server)

	_ = r.Run(":8080")
}
