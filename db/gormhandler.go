package db

import (
	"github.com/tul1/openapi_go_demo/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMHandler struct {
	db *gorm.DB
}

// NewGORMHandler initializes a new GORMHandler with a SQLite database.
func NewGORMHandler(dsn string) (*GORMHandler, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the Todo struct to create the table.
	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		return nil, err
	}

	return &GORMHandler{db: db}, nil
}

func (g *GORMHandler) GetTodos() ([]model.Todo, error) {
	var todos []model.Todo
	if err := g.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (g *GORMHandler) AddTodo(todo model.Todo) error {
	return g.db.Create(&todo).Error
}

func (g *GORMHandler) GetTodoByID(id int) (model.Todo, error) {
	var todo model.Todo
	if err := g.db.First(&todo, id).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func (g *GORMHandler) DeleteTodoByID(id int) error {
	return g.db.Delete(&model.Todo{}, id).Error
}
